#!/usr/bin/env python3
"""
Измерение CPU/сеть/диск overhead контейнера ml-server через `docker stats`.
Дополняет уже измеренную память в простое (~160.8 МБ) — единственный
незакрытый пункт: "CPU/сеть/диск ML-компонента изолированно не измерялись".

`docker stats --no-stream --format ...` отдаёт всё нужное одним вызовом на
контейнер: CPUPerc, MemUsage, NetIO, BlockIO — не нужны отдельные
инструменты (iostat/nethogs и т.п.), которых может не быть в среде.

Метод:
1. Простой — сэмплировать docker stats N раз с интервалом, пока ml-server
   поднят, ноды созданы, подов нет.
2. Нагрузка — сэмплировать непрерывно, пока в фоне идёт обычный
   run_benchmark.py (40 подов, статичная гетерогенная нагрузка).
3. Сравнить: CPU% (среднее/пик), дельту NetIO и BlockIO за время прогона
   (не мгновенное значение — оно кумулятивное с момента старта контейнера).

Запуск:
  # простой (кластер уже поднят, ноды созданы, подов нет)
  python3 measure_ml_server_overhead.py --mode idle --duration 30 --out idle_stats.json

  # нагрузка (запустить ПАРАЛЛЕЛЬНО с run_benchmark.py в другом терминале)
  python3 measure_ml_server_overhead.py --mode load --duration 60 --out load_stats.json

  # сравнение
  python3 measure_ml_server_overhead.py --compare idle_stats.json load_stats.json
"""
import argparse
import json
import re
import subprocess
import time


def parse_size_to_bytes(s: str) -> float:
    """'12.3MB' / '1.2GB' / '512B' / '0B' -> байты (float)."""
    s = s.strip()
    m = re.match(r"([\d.]+)\s*([A-Za-z]+)", s)
    if not m:
        return 0.0
    value, unit = float(m.group(1)), m.group(2).lower()
    mult = {
        "b": 1, "kb": 1e3, "mb": 1e6, "gb": 1e9,
        "kib": 1024, "mib": 1024**2, "gib": 1024**3,
    }.get(unit, 1)
    return value * mult


def sample_docker_stats(container: str = "ml-server") -> dict:
    """Один снимок docker stats для контейнера. Формат вывода:
    CPUPerc\tMemUsage\tNetIO\tBlockIO
    например: '2.34%\t45.2MiB / 1GiB\t1.2kB / 3.4kB\t0B / 12kB'
    """
    out = subprocess.run(
        ["docker", "stats", "--no-stream", "--format",
         "{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}\t{{.BlockIO}}", container],
        capture_output=True, text=True, timeout=10,
    )
    if out.returncode != 0:
        raise RuntimeError(f"docker stats failed: {out.stderr}")

    cpu_str, mem_str, net_str, block_str = out.stdout.strip().split("\t")
    cpu_pct = float(cpu_str.replace("%", ""))

    mem_used_str = mem_str.split("/")[0].strip()
    mem_used_bytes = parse_size_to_bytes(mem_used_str)

    net_rx_str, net_tx_str = [x.strip() for x in net_str.split("/")]
    net_rx_bytes = parse_size_to_bytes(net_rx_str)
    net_tx_bytes = parse_size_to_bytes(net_tx_str)

    block_read_str, block_write_str = [x.strip() for x in block_str.split("/")]
    block_read_bytes = parse_size_to_bytes(block_read_str)
    block_write_bytes = parse_size_to_bytes(block_write_str)

    return {
        "timestamp": time.time(),
        "cpu_pct": cpu_pct,
        "mem_used_bytes": mem_used_bytes,
        "net_rx_bytes": net_rx_bytes,   # кумулятивно с момента старта контейнера
        "net_tx_bytes": net_tx_bytes,
        "block_read_bytes": block_read_bytes,  # кумулятивно
        "block_write_bytes": block_write_bytes,
    }


def run_sampling(container: str, duration: float, interval: float) -> list:
    samples = []
    t0 = time.time()
    while time.time() - t0 < duration:
        try:
            samples.append(sample_docker_stats(container))
        except Exception as e:
            print(f"[WARN] сэмпл не удался: {e}")
        time.sleep(interval)
    return samples


def summarize(samples: list) -> dict:
    if not samples:
        raise SystemExit("Нет ни одного успешного сэмпла.")
    cpu_vals = [s["cpu_pct"] for s in samples]
    mem_vals = [s["mem_used_bytes"] for s in samples]

    return {
        "n_samples": len(samples),
        "duration_s": samples[-1]["timestamp"] - samples[0]["timestamp"],
        "cpu_pct_mean": sum(cpu_vals) / len(cpu_vals),
        "cpu_pct_max": max(cpu_vals),
        "mem_used_bytes_mean": sum(mem_vals) / len(mem_vals),
        "mem_used_bytes_max": max(mem_vals),
        # дельты кумулятивных счётчиков за время сэмплирования — это и есть
        # реальный трафик/I-O за период, а не мгновенное значение
        "net_rx_delta_bytes": samples[-1]["net_rx_bytes"] - samples[0]["net_rx_bytes"],
        "net_tx_delta_bytes": samples[-1]["net_tx_bytes"] - samples[0]["net_tx_bytes"],
        "block_read_delta_bytes": samples[-1]["block_read_bytes"] - samples[0]["block_read_bytes"],
        "block_write_delta_bytes": samples[-1]["block_write_bytes"] - samples[0]["block_write_bytes"],
        "raw_samples": samples,
    }


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--mode", choices=["idle", "load"])
    ap.add_argument("--container", default="ml-server")
    ap.add_argument("--duration", type=float, default=30.0, help="секунд сэмплирования")
    ap.add_argument("--interval", type=float, default=1.0, help="секунд между сэмплами")
    ap.add_argument("--out", default="stats.json")
    ap.add_argument("--compare", nargs=2, metavar=("IDLE_JSON", "LOAD_JSON"))
    args = ap.parse_args()

    if args.compare:
        idle = json.load(open(args.compare[0]))
        load = json.load(open(args.compare[1]))

        print(f"{'Метрика':<28}{'Простой':>15}{'Нагрузка':>15}{'Дельта':>15}")
        print(f"{'CPU% (среднее)':<28}{idle['cpu_pct_mean']:>14.2f}%{load['cpu_pct_mean']:>14.2f}%"
              f"{load['cpu_pct_mean']-idle['cpu_pct_mean']:>14.2f}%")
        print(f"{'CPU% (пик)':<28}{idle['cpu_pct_max']:>14.2f}%{load['cpu_pct_max']:>14.2f}%"
              f"{load['cpu_pct_max']-idle['cpu_pct_max']:>14.2f}%")
        print(f"{'Память (среднее, МБ)':<28}{idle['mem_used_bytes_mean']/1e6:>14.1f}"
              f"{load['mem_used_bytes_mean']/1e6:>14.1f}"
              f"{(load['mem_used_bytes_mean']-idle['mem_used_bytes_mean'])/1e6:>14.1f}")
        print(f"{'Сеть RX за период (КБ)':<28}{idle['net_rx_delta_bytes']/1e3:>14.1f}"
              f"{load['net_rx_delta_bytes']/1e3:>14.1f}")
        print(f"{'Сеть TX за период (КБ)':<28}{idle['net_tx_delta_bytes']/1e3:>14.1f}"
              f"{load['net_tx_delta_bytes']/1e3:>14.1f}")
        print(f"{'Диск чтение за период (КБ)':<28}{idle['block_read_delta_bytes']/1e3:>14.1f}"
              f"{load['block_read_delta_bytes']/1e3:>14.1f}")
        print(f"{'Диск запись за период (КБ)':<28}{idle['block_write_delta_bytes']/1e3:>14.1f}"
              f"{load['block_write_delta_bytes']/1e3:>14.1f}")

        # Нормировка сетевого трафика на под — содержательное число для отчёта
        if load.get("n_pods_hint"):
            n = load["n_pods_hint"]
            net_per_pod = (load['net_rx_delta_bytes'] + load['net_tx_delta_bytes']) / n
            print(f"\nСетевой трафик на под (RX+TX): {net_per_pod/1e3:.2f} КБ "
                  f"(при n_pods_hint={n})")
        return

    if not args.mode:
        raise SystemExit("Укажите --mode idle|load, либо --compare IDLE_JSON LOAD_JSON")

    print(f"Сэмплирую {args.container} в режиме '{args.mode}' "
          f"({args.duration}с, интервал {args.interval}с)...")
    samples = run_sampling(args.container, args.duration, args.interval)
    summary = summarize(samples)

    print(f"Сэмплов: {summary['n_samples']}, фактическая длительность: {summary['duration_s']:.1f}с")
    print(f"CPU% среднее/пик: {summary['cpu_pct_mean']:.2f}% / {summary['cpu_pct_max']:.2f}%")
    print(f"Память среднее/пик: {summary['mem_used_bytes_mean']/1e6:.1f} / "
          f"{summary['mem_used_bytes_max']/1e6:.1f} МБ")
    print(f"Сеть RX/TX за период: {summary['net_rx_delta_bytes']/1e3:.1f} / "
          f"{summary['net_tx_delta_bytes']/1e3:.1f} КБ")
    print(f"Диск read/write за период: {summary['block_read_delta_bytes']/1e3:.1f} / "
          f"{summary['block_write_delta_bytes']/1e3:.1f} КБ")

    with open(args.out, "w") as f:
        json.dump(summary, f, indent=2)
    print(f"\nСохранено: {args.out}")


if __name__ == "__main__":
    main()
