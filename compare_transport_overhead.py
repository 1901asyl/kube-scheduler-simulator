"""
Сравнивает results_timing_http/ (транспорт=HTTP) и results_timing_nng/ (транспорт=NNG),
каждый со свипом NOOP_MODE={0,1} x ARTIFICIAL_DELAY_MS={0,50,100,500,1500}.

Разложение overhead (при ARTIFICIAL_DELAY_MS=0):
  elapsed_ms(transport, NOOP=1) ~= transport_overhead(transport)
  elapsed_ms(transport, NOOP=0) - elapsed_ms(transport, NOOP=1) ~= inference_time

  transport_overhead(HTTP) - transport_overhead(NNG)
      = elapsed_ms(HTTP, NOOP=1, DELAY=0) - elapsed_ms(NNG, NOOP=1, DELAY=0)
      -> чистый вклад транспорта, без какой-либо примеси инференса

  inference_time не должен зависеть от транспорта (та же модель, те же фичи) —
  совпадение inference_time(HTTP) и inference_time(NNG) по обеим формулам ниже
  само по себе служит проверкой корректности изоляции.
"""
import numpy as np

DELAYS = [0, 50, 100, 500, 1500]
NOOPS = [0, 1]
TRANSPORTS = {
    "HTTP": "results_timing_http",
    "NNG": "results_timing_nng",
}


def load(path):
    try:
        return np.loadtxt(path)
    except OSError:
        return np.array([])


def mean_or_none(data):
    return float(data.mean()) if data.size else None


def std_or_none(data):
    return float(data.std()) if data.size else None


def label(transport, noop, delay):
    prefix = "http" if transport == "HTTP" else "nng"
    return f"timing_{prefix}_noop{noop}_{delay}ms"


def main():
    # 1. Полная таблица
    print(f"{'transport':>8} | {'NOOP':>4} | {'DELAY':>6} | {'mean elapsed_ms':>16} | {'std':>8} | {'n':>4}")
    print("-" * 62)
    means = {}
    stds = {}
    for transport, out_dir in TRANSPORTS.items():
        for noop in NOOPS:
            for delay in DELAYS:
                lbl = label(transport, noop, delay)
                data = load(f"{out_dir}/real_latency_{lbl}.txt")
                m = mean_or_none(data)
                s = std_or_none(data)
                means[(transport, noop, delay)] = m
                stds[(transport, noop, delay)] = s
                print(f"{transport:>8} | {noop:>4} | {delay:>6} | "
                      f"{'n/a' if m is None else f'{m:.2f}':>16} | "
                      f"{'n/a' if s is None else f'{s:.2f}':>8} | "
                      f"{data.size:>4}")

    # 2. Разложение при DELAY=0
    print()
    print("=== Разложение overhead при ARTIFICIAL_DELAY_MS=0 ===")
    http_noop1 = means.get(("HTTP", 1, 0))
    nng_noop1 = means.get(("NNG", 1, 0))
    http_noop0 = means.get(("HTTP", 0, 0))
    nng_noop0 = means.get(("NNG", 0, 0))
    http_noop1_std = stds.get(("HTTP", 1, 0)) or 0.0
    nng_noop1_std = stds.get(("NNG", 1, 0)) or 0.0

    if http_noop1 is not None and nng_noop1 is not None:
        transport_diff = http_noop1 - nng_noop1
        print(f"Вклад транспорта (HTTP - NNG, NOOP=1): {transport_diff:.2f} мс "
              f"(std HTTP={http_noop1_std:.2f}, std NNG={nng_noop1_std:.2f})")
        if transport_diff < max(http_noop1_std, nng_noop1_std):
            print("  -> ПРЕДУПРЕЖДЕНИЕ: разница меньше std одного из транспортов —")
            print("     нужно больше повторов, прежде чем утверждать значимость.")

    if http_noop0 is not None and http_noop1 is not None:
        print(f"Вклад инференса модели через HTTP: {http_noop0 - http_noop1:.2f} мс")
    if nng_noop0 is not None and nng_noop1 is not None:
        print(f"Вклад инференса модели через NNG:  {nng_noop0 - nng_noop1:.2f} мс")

    print()
    print("Если два значения 'вклад инференса' выше расходятся сильнее, чем сам")
    print("'вклад транспорта' — одного прогона на комбинацию недостаточно для")
    print("уверенного вывода; нужно несколько независимых повторов и")
    print("доверительный интервал, а не одно среднее по одному прогону.")


if __name__ == "__main__":
    main()
