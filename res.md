bench-tiny-fjt6j               profile=tiny     -> node=node-pwvw4 (56.4ms)
bench-tiny-r9494               profile=tiny     -> node=node-hq4vt (163.3ms)
bench-tiny-8xcht               profile=tiny     -> node=node-b7gqv (164.1ms)
bench-tiny-vdtdc               profile=tiny     -> node=node-hq4vt (163.9ms)
bench-tiny-cr59q               profile=tiny     -> node=node-pwvw4 (164.4ms)
bench-tiny-95kvz               profile=tiny     -> node=node-hq4vt (163.3ms)
bench-tiny-xkxjj               profile=tiny     -> node=node-hq4vt (162.3ms)
bench-tiny-4sq9w               profile=tiny     -> node=node-pwvw4 (163.5ms)
bench-small-tqrsk              profile=small    -> node=node-hq4vt (163.6ms)
bench-small-fpnz8              profile=small    -> node=node-pwvw4 (163.2ms)
bench-small-7c4vl              profile=small    -> node=node-b7gqv (163.1ms)
bench-small-9dcwr              profile=small    -> node=node-hq4vt (163.3ms)
bench-small-rvqxk              profile=small    -> node=node-pwvw4 (164.4ms)
bench-small-bn8vb              profile=small    -> node=node-hq4vt (163.4ms)
bench-small-t2jqh              profile=small    -> node=node-hq4vt (164.7ms)
bench-small-z9227              profile=small    -> node=node-pwvw4 (164.9ms)
bench-medium-4xpgd             profile=medium   -> node=node-b7gqv (164.8ms)
bench-medium-8t2j6             profile=medium   -> node=node-hq4vt (164.1ms)
bench-medium-lp5ff             profile=medium   -> node=node-pwvw4 (163.2ms)
bench-medium-dq7m9             profile=medium   -> node=node-hq4vt (165.1ms)
bench-medium-ngcrk             profile=medium   -> node=node-pwvw4 (164.4ms)
bench-medium-ghhrd             profile=medium   -> node=node-hq4vt (163.6ms)
bench-medium-dk8gp             profile=medium   -> node=node-b7gqv (163.7ms)
bench-medium-vnxl6             profile=medium   -> node=node-hq4vt (163.1ms)
bench-large-5n8pj              profile=large    -> node=node-pwvw4 (162.3ms)
bench-large-ccxkk              profile=large    -> node=node-hq4vt (163.1ms)
bench-large-r5q8b              profile=large    -> node=node-b7gqv (56.8ms)
bench-large-h9gfk              profile=large    -> node=node-hq4vt (162.4ms)
bench-large-8v7tz              profile=large    -> node=node-pwvw4 (163.8ms)
bench-large-42nzd              profile=large    -> node=node-hq4vt (164.5ms)
bench-large-rnnlk              profile=large    -> node=node-pwvw4 (163.9ms)
bench-large-vsk9x              profile=large    -> node=node-hq4vt (164.0ms)
bench-xlarge-x9d96             profile=xlarge   -> node=node-hq4vt (163.8ms)
bench-xlarge-hkt9h             profile=xlarge   -> node=node-b7gqv (163.6ms)
bench-xlarge-4pgdm             profile=xlarge   -> node=node-pwvw4 (164.2ms)
bench-xlarge-8xrpb             profile=xlarge   -> node=node-hq4vt (163.9ms)
bench-xlarge-49gpb             profile=xlarge   -> node=node-pwvw4 (163.5ms)
bench-xlarge-c9pmp             profile=xlarge   -> node=node-hq4vt (164.9ms)
bench-xlarge-ghv4z             profile=xlarge   -> node=node-hq4vt (164.6ms)
bench-xlarge-xzlpt             profile=xlarge   -> node=node-b7gqv (56.2ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_0ms_results.csv
Per-pod scheduling latency: mean=155.73ms p95=164.92ms max=165.08ms
Метрики балансировки: results_timing_http/timing_http_noop0_0ms_balance.json
  CPU utilization std=0.0948 mean=0.6670
  Memory utilization std=0.0958 mean=0.6610

real	0m6.516s
user	0m0.190s
sys	0m0.081s
Переключено на NNG. Порт ml-server должен быть 5555 — проверьте compose.local.yml вручную:
Проверка: grep transport= в рабочем plugin.go должен показать transport=nng
164:	klog.Infof("AIScore inference latency: node=%s pod=%s elapsed_ms=%.2f transport=nng",
171:		klog.Warningf("AIScore fallback: node=%s pod=%s err=%v elapsed_ms=%.2f transport=nng"


 => [simulator-scheduler build-env 7/8] COPY . .                                                                                                                                    0.2s
 => CANCELED [ml-server 4/6] RUN pip install --no-cache-dir -r requirements.txt                                                                                                     2.8s
 => ERROR [simulator-scheduler build-env 8/8] RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache     go build -v -o ./bin/scheduler ./cmd/scheduler/sc  2.5s
------                                                                                                                                                                                   
 > [simulator-scheduler build-env 8/8] RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache     go build -v -o ./bin/scheduler ./cmd/scheduler/scheduler.go:   
2.448 docs/sample/ai_score/plugin.go:15:2: no required module provides package go.nanomsg.org/mangos/v3; to add it:                                                                      
2.448   go get go.nanomsg.org/mangos/v3                                                                                                                                                  
2.448 docs/sample/ai_score/plugin.go:16:2: no required module provides package go.nanomsg.org/mangos/v3/protocol/req; to add it:                                                         
2.448   go get go.nanomsg.org/mangos/v3/protocol/req                                                                                                                                     
2.448 docs/sample/ai_score/plugin.go:18:2: no required module provides package go.nanomsg.org/mangos/v3/transport/tcp; to add it:                                                        
2.448 	go get go.nanomsg.org/mangos/v3/transport/tcp
------
[+] build 0/2
 ⠙ Image kube-scheduler-simulator-ml-server Building                                                                                                                                 3.8s
 ⠙ Image simulator-scheduler                Building                                                                                                                                 3.8s
Dockerfile:19

--------------------

  18 |     COPY . .

  19 | >>> RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \

  20 | >>>     go build -v -o ./bin/scheduler ./cmd/scheduler/scheduler.go

  21 |     

--------------------

target simulator-scheduler: failed to solve: process "/bin/sh -c go build -v -o ./bin/scheduler ./cmd/scheduler/scheduler.go" did not complete successfully: exit code: 1

=== [NNG] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=0 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                           
 
 bench-tiny-bn4g2               profile=tiny     -> node=node-m25ml (56.1ms)
bench-tiny-mjkrq               profile=tiny     -> node=node-bpx6v (163.2ms)
bench-tiny-pzlkq               profile=tiny     -> node=node-wkxhh (163.8ms)
bench-tiny-x4t6p               profile=tiny     -> node=node-bpx6v (164.3ms)
bench-tiny-k4kk5               profile=tiny     -> node=node-m25ml (165.8ms)
bench-tiny-w4lwf               profile=tiny     -> node=node-bpx6v (163.7ms)
bench-tiny-wqqd2               profile=tiny     -> node=node-bpx6v (57.1ms)
bench-tiny-wmksz               profile=tiny     -> node=node-m25ml (166.1ms)
bench-small-bstm8              profile=small    -> node=node-bpx6v (163.8ms)
bench-small-jqdlf              profile=small    -> node=node-m25ml (163.0ms)
bench-small-d8ffj              profile=small    -> node=node-wkxhh (163.9ms)
bench-small-cqf6k              profile=small    -> node=node-bpx6v (163.2ms)
bench-small-28fjx              profile=small    -> node=node-bpx6v (164.1ms)
bench-small-ztsrg              profile=small    -> node=node-m25ml (164.8ms)
bench-small-mm2g7              profile=small    -> node=node-bpx6v (165.1ms)
bench-small-s7k8p              profile=small    -> node=node-m25ml (163.7ms)
bench-medium-z24wq             profile=medium   -> node=node-wkxhh (163.3ms)
bench-medium-n75z4             profile=medium   -> node=node-bpx6v (163.7ms)
bench-medium-lwtgz             profile=medium   -> node=node-m25ml (163.9ms)
bench-medium-zld2h             profile=medium   -> node=node-bpx6v (162.8ms)
bench-medium-v8lfn             profile=medium   -> node=node-m25ml (163.9ms)
bench-medium-pqsv5             profile=medium   -> node=node-bpx6v (166.0ms)
bench-medium-9jhkj             profile=medium   -> node=node-wkxhh (164.7ms)
bench-medium-wrb7d             profile=medium   -> node=node-bpx6v (164.3ms)
bench-large-fjzr6              profile=large    -> node=node-m25ml (164.7ms)
bench-large-4q6m6              profile=large    -> node=node-bpx6v (163.8ms)
bench-large-k82zm              profile=large    -> node=node-wkxhh (165.8ms)
bench-large-f8jx9              profile=large    -> node=node-bpx6v (163.6ms)
bench-large-r9v9t              profile=large    -> node=node-m25ml (164.0ms)
bench-large-lqprf              profile=large    -> node=node-bpx6v (163.7ms)
bench-large-f95zx              profile=large    -> node=node-m25ml (163.8ms)
bench-large-9sb42              profile=large    -> node=node-bpx6v (164.1ms)
bench-xlarge-4hwt7             profile=xlarge   -> node=node-bpx6v (163.6ms)
bench-xlarge-2jtjw             profile=xlarge   -> node=node-wkxhh (164.5ms)
bench-xlarge-kfh8x             profile=xlarge   -> node=node-m25ml (164.3ms)
bench-xlarge-xd4fd             profile=xlarge   -> node=node-bpx6v (163.7ms)
bench-xlarge-vh5jm             profile=xlarge   -> node=node-m25ml (162.8ms)
bench-xlarge-mbvcd             profile=xlarge   -> node=node-bpx6v (164.0ms)
bench-xlarge-mwwvd             profile=xlarge   -> node=node-bpx6v (164.3ms)
bench-xlarge-8kgg5             profile=xlarge   -> node=node-wkxhh (164.3ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_0ms_results.csv
Per-pod scheduling latency: mean=158.74ms p95=165.97ms max=166.11ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_0ms_balance.json
  CPU utilization std=0.0948 mean=0.6670
  Memory utilization std=0.0958 mean=0.6610

real	0m6.636s
user	0m0.213s
sys	0m0.063s
transport | NOOP |  DELAY |  mean elapsed_ms
----------------------------------------------
/kube-scheduler-simulator/compare_transport_overhead.py:29: UserWarning: loadtxt: input contained no data: "results_timing_http/real_latency_timing_http_noop0_0ms.txt"
  return np.loadtxt(path)
    HTTP |    0 |      0 | n/a
    HTTP |    0 |     50 | n/a
    HTTP |    0 |    100 | n/a
    HTTP |    0 |    500 | n/a
    HTTP |    0 |   1500 | n/a
    HTTP |    1 |      0 | n/a
    HTTP |    1 |     50 | n/a
    HTTP |    1 |    100 | n/a
    HTTP |    1 |    500 | n/a
    HTTP |    1 |   1500 | n/a
/kube-scheduler-simulator/compare_transport_overhead.py:29: UserWarning: loadtxt: input contained no data: "results_timing_nng/real_latency_timing_nng_noop0_0ms.txt"
  return np.loadtxt(path)
     NNG |    0 |      0 | n/a
     NNG |    0 |     50 | n/a
     NNG |    0 |    100 | n/a
     NNG |    0 |    500 | n/a
     NNG |    0 |   1500 | n/a
     NNG |    1 |      0 | n/a
     NNG |    1 |     50 | n/a
     NNG |    1 |    100 | n/a
     NNG |    1 |    500 | n/a
     NNG |    1 |   1500 | n/a

=== Разложение overhead при ARTIFICIAL_DELAY_MS=0 ===

Если два значения 'вклад инференса' выше совпадают в пределах шума —
это подтверждает, что transport_diff действительно изолирует транспорт,
а не какой-то побочный эффект смены библиотеки.

