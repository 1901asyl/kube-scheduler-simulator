bench-tiny-n6g7t               profile=tiny     -> node=node-597qs (108.7ms)
bench-tiny-jrzv7               profile=tiny     -> node=node-nktfw (217.9ms)
bench-tiny-sq2vd               profile=tiny     -> node=node-597qs (271.0ms)
bench-tiny-x6gjk               profile=tiny     -> node=node-8g8m5 (218.8ms)
bench-tiny-w67w6               profile=tiny     -> node=node-597qs (216.4ms)
bench-tiny-5lz6x               profile=tiny     -> node=node-597qs (216.6ms)
bench-tiny-8hg4s               profile=tiny     -> node=node-nktfw (216.0ms)
bench-tiny-h62xz               profile=tiny     -> node=node-nktfw (214.9ms)
bench-small-4s267              profile=small    -> node=node-597qs (217.5ms)
bench-small-tk96k              profile=small    -> node=node-nktfw (217.0ms)
bench-small-wgk8n              profile=small    -> node=node-597qs (217.2ms)
bench-small-hpzvq              profile=small    -> node=node-8g8m5 (216.0ms)
bench-small-pcdhx              profile=small    -> node=node-nktfw (216.1ms)
bench-small-lld4l              profile=small    -> node=node-597qs (215.7ms)
bench-small-bsvtm              profile=small    -> node=node-597qs (216.5ms)
bench-small-t5flp              profile=small    -> node=node-597qs (108.6ms)
bench-medium-l25ww             profile=medium   -> node=node-nktfw (215.7ms)
bench-medium-p9p4w             profile=medium   -> node=node-597qs (215.3ms)
bench-medium-fp7qv             profile=medium   -> node=node-8g8m5 (214.8ms)
bench-medium-nx5nf             profile=medium   -> node=node-597qs (217.0ms)
bench-medium-2z4vk             profile=medium   -> node=node-nktfw (215.9ms)
bench-medium-thpx8             profile=medium   -> node=node-597qs (215.7ms)
bench-medium-d9wh4             profile=medium   -> node=node-nktfw (215.6ms)
bench-medium-6ndss             profile=medium   -> node=node-597qs (214.9ms)
bench-large-fjbt7              profile=large    -> node=node-597qs (216.7ms)
bench-large-pdlrj              profile=large    -> node=node-nktfw (215.7ms)
bench-large-2rkmg              profile=large    -> node=node-597qs (215.8ms)
bench-large-t6k66              profile=large    -> node=node-8g8m5 (215.8ms)
bench-large-7td49              profile=large    -> node=node-nktfw (215.7ms)
bench-large-rw29g              profile=large    -> node=node-597qs (215.6ms)
bench-large-4fjmh              profile=large    -> node=node-597qs (216.2ms)
bench-large-2xfsq              profile=large    -> node=node-nktfw (218.0ms)
bench-xlarge-47zfs             profile=xlarge   -> node=node-597qs (216.9ms)
bench-xlarge-mf8mw             profile=xlarge   -> node=node-nktfw (215.1ms)
bench-xlarge-zc9hd             profile=xlarge   -> node=node-8g8m5 (214.2ms)
bench-xlarge-99s52             profile=xlarge   -> node=node-597qs (215.2ms)
bench-xlarge-jnv7l             profile=xlarge   -> node=node-597qs (214.8ms)
bench-xlarge-m2qjq             profile=xlarge   -> node=node-nktfw (214.0ms)
bench-xlarge-86qcd             profile=xlarge   -> node=node-597qs (215.2ms)
bench-xlarge-hdzjr             profile=xlarge   -> node=node-nktfw (213.6ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_0ms_results.csv
Per-pod scheduling latency: mean=211.96ms p95=218.77ms max=271.03ms
Метрики балансировки: results_timing_http/timing_http_noop0_0ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m8.752s
user	0m0.199s
sys	0m0.097s
=== [HTTP] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=50 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-w2gs6               profile=tiny     -> node=node-hl2vw (269.8ms)
bench-tiny-8pt89               profile=tiny     -> node=node-wgrpm (269.1ms)
bench-tiny-bg898               profile=tiny     -> node=node-hl2vw (320.9ms)
bench-tiny-8kq6g               profile=tiny     -> node=node-ksvn4 (268.9ms)
bench-tiny-m76jm               profile=tiny     -> node=node-wgrpm (268.7ms)
bench-tiny-db9c8               profile=tiny     -> node=node-hl2vw (269.2ms)
bench-tiny-9nmnq               profile=tiny     -> node=node-hl2vw (269.5ms)
bench-tiny-vrzpj               profile=tiny     -> node=node-wgrpm (270.4ms)
bench-small-w58pv              profile=small    -> node=node-hl2vw (268.8ms)
bench-small-mrqn5              profile=small    -> node=node-hl2vw (269.4ms)
bench-small-74gzz              profile=small    -> node=node-wgrpm (269.3ms)
bench-small-dzll6              profile=small    -> node=node-hl2vw (269.7ms)
bench-small-7vb52              profile=small    -> node=node-wgrpm (269.4ms)
bench-small-sjlhm              profile=small    -> node=node-ksvn4 (268.5ms)
bench-small-g5lrh              profile=small    -> node=node-hl2vw (268.8ms)
bench-small-sntzw              profile=small    -> node=node-hl2vw (268.5ms)
bench-medium-98hsm             profile=medium   -> node=node-wgrpm (267.3ms)
bench-medium-q282x             profile=medium   -> node=node-hl2vw (269.3ms)
bench-medium-2hmqr             profile=medium   -> node=node-ksvn4 (269.3ms)
bench-medium-g9fd6             profile=medium   -> node=node-wgrpm (269.1ms)
bench-medium-8jlbx             profile=medium   -> node=node-hl2vw (268.0ms)
bench-medium-d4kqj             profile=medium   -> node=node-hl2vw (268.3ms)
bench-medium-bv657             profile=medium   -> node=node-wgrpm (267.8ms)
bench-medium-w7rwb             profile=medium   -> node=node-hl2vw (269.9ms)
bench-large-24hc6              profile=large    -> node=node-hl2vw (269.9ms)
bench-large-b2c8w              profile=large    -> node=node-wgrpm (269.2ms)
bench-large-cf4xh              profile=large    -> node=node-hl2vw (268.9ms)
bench-large-kzmtc              profile=large    -> node=node-ksvn4 (270.7ms)
bench-large-cjc74              profile=large    -> node=node-wgrpm (270.5ms)
bench-large-7hdtb              profile=large    -> node=node-hl2vw (267.9ms)
bench-large-mdsf8              profile=large    -> node=node-hl2vw (269.4ms)
bench-large-8qrdj              profile=large    -> node=node-wgrpm (268.8ms)
bench-xlarge-fsvl6             profile=xlarge   -> node=node-hl2vw (266.7ms)
bench-xlarge-46btn             profile=xlarge   -> node=node-wgrpm (269.0ms)
bench-xlarge-t8bfl             profile=xlarge   -> node=node-hl2vw (161.5ms)
bench-xlarge-5nt96             profile=xlarge   -> node=node-ksvn4 (269.0ms)
bench-xlarge-m6scw             profile=xlarge   -> node=node-hl2vw (268.9ms)
bench-xlarge-bvgwd             profile=xlarge   -> node=node-wgrpm (269.6ms)
bench-xlarge-fdtmg             profile=xlarge   -> node=node-hl2vw (268.8ms)
bench-xlarge-njt89             profile=xlarge   -> node=node-wgrpm (270.0ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_50ms_results.csv
Per-pod scheduling latency: mean=267.67ms p95=270.66ms max=320.94ms
Метрики балансировки: results_timing_http/timing_http_noop0_50ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m10.979s
user	0m0.242s
sys	0m0.098s
=== [HTTP] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=100 ===
[+] down 10/10
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-fqc54               profile=tiny     -> node=node-jhqql (214.9ms)
bench-tiny-b7hkp               profile=tiny     -> node=node-mt57s (323.2ms)
bench-tiny-kphm9               profile=tiny     -> node=node-mt57s (377.1ms)
bench-tiny-sqj2v               profile=tiny     -> node=node-5jvct (326.5ms)
bench-tiny-ltb5d               profile=tiny     -> node=node-jhqql (323.2ms)
bench-tiny-nfnl6               profile=tiny     -> node=node-jhqql (323.3ms)
bench-tiny-qt25r               profile=tiny     -> node=node-mt57s (323.8ms)
bench-tiny-nqpj6               profile=tiny     -> node=node-mt57s (322.9ms)
bench-small-8vfrb              profile=small    -> node=node-mt57s (322.8ms)
bench-small-4qkgq              profile=small    -> node=node-mt57s (323.6ms)
bench-small-z4dwr              profile=small    -> node=node-jhqql (321.6ms)
bench-small-7mnsd              profile=small    -> node=node-5jvct (320.4ms)
bench-small-vdfpq              profile=small    -> node=node-jhqql (320.4ms)
bench-small-ctfkw              profile=small    -> node=node-mt57s (323.1ms)
bench-small-8kkcj              profile=small    -> node=node-mt57s (319.7ms)
bench-small-lzg2c              profile=small    -> node=node-mt57s (324.3ms)
bench-medium-mqc76             profile=medium   -> node=node-jhqql (324.5ms)
bench-medium-jk75k             profile=medium   -> node=node-mt57s (323.6ms)
bench-medium-6c4cw             profile=medium   -> node=node-5jvct (324.7ms)
bench-medium-n2zmn             profile=medium   -> node=node-jhqql (323.4ms)
bench-medium-wc97p             profile=medium   -> node=node-mt57s (323.6ms)
bench-medium-2hs9z             profile=medium   -> node=node-mt57s (322.4ms)
bench-medium-msvhv             profile=medium   -> node=node-jhqql (322.5ms)
bench-medium-x5b9m             profile=medium   -> node=node-mt57s (329.7ms)
bench-large-krf8s              profile=large    -> node=node-mt57s (323.6ms)
bench-large-bgz6r              profile=large    -> node=node-jhqql (323.2ms)
bench-large-qsl88              profile=large    -> node=node-5jvct (323.3ms)
bench-large-4cwb8              profile=large    -> node=node-mt57s (324.0ms)
bench-large-f9hjx              profile=large    -> node=node-jhqql (322.7ms)
bench-large-2ds82              profile=large    -> node=node-mt57s (323.3ms)
bench-large-qc6pw              profile=large    -> node=node-mt57s (324.6ms)
bench-large-fv2ml              profile=large    -> node=node-jhqql (214.4ms)
bench-xlarge-d26nr             profile=xlarge   -> node=node-mt57s (215.3ms)
bench-xlarge-l2jz4             profile=xlarge   -> node=node-jhqql (322.8ms)
bench-xlarge-thhdv             profile=xlarge   -> node=node-5jvct (324.1ms)
bench-xlarge-v8zg9             profile=xlarge   -> node=node-mt57s (323.2ms)
bench-xlarge-46q6r             profile=xlarge   -> node=node-mt57s (323.2ms)
bench-xlarge-68bl4             profile=xlarge   -> node=node-jhqql (322.9ms)
bench-xlarge-mntsp             profile=xlarge   -> node=node-mt57s (324.6ms)
bench-xlarge-7wlb7             profile=xlarge   -> node=node-jhqql (324.1ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_100ms_results.csv
Per-pod scheduling latency: mean=316.61ms p95=329.74ms max=377.10ms
Метрики балансировки: results_timing_http/timing_http_noop0_100ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m12.943s
user	0m0.264s
sys	0m0.114s
=== [HTTP] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=500 ===
[+] down 10/10
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-9zxjs               profile=tiny     -> node=node-9c294 (645.0ms)
bench-tiny-gtsvp               profile=tiny     -> node=node-9c294 (749.5ms)
bench-tiny-46kcv               profile=tiny     -> node=node-tjr7l (749.6ms)
bench-tiny-zpzgn               profile=tiny     -> node=node-tjr7l (697.9ms)
bench-tiny-bj85m               profile=tiny     -> node=node-84vwz (749.7ms)
bench-tiny-xrxl5               profile=tiny     -> node=node-tjr7l (749.3ms)
bench-tiny-lqdhl               profile=tiny     -> node=node-9c294 (750.0ms)
bench-tiny-dkrj9               profile=tiny     -> node=node-9c294 (749.3ms)
bench-small-v8qb7              profile=small    -> node=node-9c294 (752.6ms)
bench-small-gxls8              profile=small    -> node=node-9c294 (698.7ms)
bench-small-7bntx              profile=small    -> node=node-tjr7l (750.5ms)
bench-small-frqm2              profile=small    -> node=node-84vwz (704.2ms)
bench-small-p9brp              profile=small    -> node=node-tjr7l (697.1ms)
bench-small-jv7dx              profile=small    -> node=node-9c294 (703.8ms)
bench-small-hf5vk              profile=small    -> node=node-9c294 (695.8ms)
bench-small-pzdqh              profile=small    -> node=node-9c294 (696.5ms)
bench-medium-ljplr             profile=medium   -> node=node-9c294 (698.6ms)
bench-medium-w2bgf             profile=medium   -> node=node-tjr7l (695.7ms)
bench-medium-tp7ng             profile=medium   -> node=node-84vwz (696.7ms)
bench-medium-j7l2c             profile=medium   -> node=node-tjr7l (700.4ms)
bench-medium-ptttm             profile=medium   -> node=node-9c294 (694.9ms)
bench-medium-dc69r             profile=medium   -> node=node-9c294 (695.2ms)
bench-medium-v9r4z             profile=medium   -> node=node-tjr7l (699.2ms)
bench-medium-qvx25             profile=medium   -> node=node-9c294 (752.4ms)
bench-large-7pl57              profile=large    -> node=node-9c294 (700.8ms)
bench-large-8695k              profile=large    -> node=node-tjr7l (749.7ms)
bench-large-wtbhp              profile=large    -> node=node-9c294 (698.6ms)
bench-large-vrzbh              profile=large    -> node=node-84vwz (697.6ms)
bench-large-77bdz              profile=large    -> node=node-tjr7l (705.9ms)
bench-large-p46j8              profile=large    -> node=node-9c294 (696.8ms)
bench-large-hqlrx              profile=large    -> node=node-9c294 (696.3ms)
bench-large-8ztlm              profile=large    -> node=node-tjr7l (699.9ms)
bench-xlarge-g7xfs             profile=xlarge   -> node=node-9c294 (704.2ms)
bench-xlarge-5lghm             profile=xlarge   -> node=node-tjr7l (696.8ms)
bench-xlarge-fr7sf             profile=xlarge   -> node=node-84vwz (751.4ms)
bench-xlarge-fjt5p             profile=xlarge   -> node=node-9c294 (751.4ms)
bench-xlarge-jr78k             profile=xlarge   -> node=node-9c294 (643.3ms)
bench-xlarge-2cmwb             profile=xlarge   -> node=node-tjr7l (697.7ms)
bench-xlarge-g4k4r             profile=xlarge   -> node=node-9c294 (697.4ms)
bench-xlarge-nqk29             profile=xlarge   -> node=node-tjr7l (696.9ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_500ms_results.csv
Per-pod scheduling latency: mean=711.43ms p95=752.44ms max=752.60ms
Метрики балансировки: results_timing_http/timing_http_noop0_500ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m28.722s
user	0m0.456s
sys	0m0.211s
=== [HTTP] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=1500 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.6s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.2s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.1s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.3s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-sxvk6               profile=tiny     -> node=node-92pss (1613.0ms)
bench-tiny-4qzt6               profile=tiny     -> node=node-bl5v8 (1718.4ms)
bench-tiny-8zwmr               profile=tiny     -> node=node-bl5v8 (1767.3ms)
bench-tiny-nktfr               profile=tiny     -> node=node-92pss (1718.0ms)
bench-tiny-99zmd               profile=tiny     -> node=node-92pss (1714.5ms)
bench-tiny-vmzsk               profile=tiny     -> node=node-5vmtn (1719.4ms)
bench-tiny-kkwsk               profile=tiny     -> node=node-bl5v8 (1719.4ms)
bench-tiny-gx9wf               profile=tiny     -> node=node-bl5v8 (1713.9ms)
bench-small-lkkcq              profile=small    -> node=node-bl5v8 (1718.0ms)
bench-small-pqfq9              profile=small    -> node=node-92pss (1716.2ms)
bench-small-l7c69              profile=small    -> node=node-bl5v8 (1717.8ms)
bench-small-2m6zd              profile=small    -> node=node-5vmtn (1718.7ms)
bench-small-mgf2d              profile=small    -> node=node-bl5v8 (1604.8ms)
bench-small-8x8kb              profile=small    -> node=node-92pss (1719.1ms)
bench-small-rfr8g              profile=small    -> node=node-bl5v8 (1714.2ms)
bench-small-mxcxt              profile=small    -> node=node-bl5v8 (1609.6ms)
bench-medium-r7nkt             profile=medium   -> node=node-bl5v8 (1723.4ms)
bench-medium-s4lpn             profile=medium   -> node=node-92pss (1717.9ms)
bench-medium-qps5q             profile=medium   -> node=node-5vmtn (1715.9ms)
bench-medium-txgmc             profile=medium   -> node=node-bl5v8 (1714.7ms)
bench-medium-nfdwq             profile=medium   -> node=node-92pss (1712.5ms)
bench-medium-wl8gm             profile=medium   -> node=node-bl5v8 (1713.4ms)
bench-medium-vk2vx             profile=medium   -> node=node-92pss (1714.5ms)
bench-medium-cjfk8             profile=medium   -> node=node-bl5v8 (1717.0ms)
bench-large-wpc46              profile=large    -> node=node-bl5v8 (1717.0ms)
bench-large-z7cj9              profile=large    -> node=node-92pss (1721.8ms)
bench-large-xrlb6              profile=large    -> node=node-bl5v8 (1716.2ms)
bench-large-s7jn8              profile=large    -> node=node-5vmtn (1716.3ms)
bench-large-5c6mf              profile=large    -> node=node-92pss (1712.9ms)
bench-large-bznk7              profile=large    -> node=node-bl5v8 (1716.4ms)
bench-large-56nwv              profile=large    -> node=node-bl5v8 (1713.0ms)
bench-large-pg2pd              profile=large    -> node=node-92pss (1713.1ms)
bench-xlarge-zpjqw             profile=xlarge   -> node=node-bl5v8 (1717.8ms)
bench-xlarge-d6xtz             profile=xlarge   -> node=node-92pss (1718.4ms)
bench-xlarge-rqpvg             profile=xlarge   -> node=node-bl5v8 (1720.4ms)
bench-xlarge-867tz             profile=xlarge   -> node=node-5vmtn (1714.4ms)
bench-xlarge-t96s4             profile=xlarge   -> node=node-bl5v8 (1714.1ms)
bench-xlarge-hstsc             profile=xlarge   -> node=node-92pss (1716.1ms)
bench-xlarge-xwwbw             profile=xlarge   -> node=node-bl5v8 (1714.8ms)
bench-xlarge-25phx             profile=xlarge   -> node=node-92pss (1717.0ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop0_1500ms_results.csv
Per-pod scheduling latency: mean=1709.78ms p95=1723.43ms max=1767.29ms
Метрики балансировки: results_timing_http/timing_http_noop0_1500ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	1m8.654s
user	0m0.919s
sys	0m0.495s
=== [HTTP] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=0 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-vqnb8               profile=tiny     -> node=node-qghf5 (56.5ms)
bench-tiny-xtv4d               profile=tiny     -> node=node-dls4b (163.5ms)
bench-tiny-nlhvs               profile=tiny     -> node=node-dls4b (164.4ms)
bench-tiny-9rvxl               profile=tiny     -> node=node-jp9v9 (164.6ms)
bench-tiny-cbjw6               profile=tiny     -> node=node-dls4b (164.3ms)
bench-tiny-w9rrf               profile=tiny     -> node=node-qghf5 (164.7ms)
bench-tiny-cbjd4               profile=tiny     -> node=node-qghf5 (164.7ms)
bench-tiny-rvlbv               profile=tiny     -> node=node-dls4b (163.6ms)
bench-small-vj7ms              profile=small    -> node=node-dls4b (163.5ms)
bench-small-qqsrv              profile=small    -> node=node-dls4b (163.8ms)
bench-small-mv7rs              profile=small    -> node=node-qghf5 (56.3ms)
bench-small-bnkjb              profile=small    -> node=node-qghf5 (165.5ms)
bench-small-6n57m              profile=small    -> node=node-dls4b (163.5ms)
bench-small-wlrms              profile=small    -> node=node-jp9v9 (164.4ms)
bench-small-wmzwl              profile=small    -> node=node-dls4b (165.5ms)
bench-small-lgk7t              profile=small    -> node=node-dls4b (163.2ms)
bench-medium-66q6l             profile=medium   -> node=node-dls4b (163.9ms)
bench-medium-46tm4             profile=medium   -> node=node-qghf5 (163.4ms)
bench-medium-ksnp7             profile=medium   -> node=node-jp9v9 (165.6ms)
bench-medium-fqm96             profile=medium   -> node=node-qghf5 (163.4ms)
bench-medium-ml8cr             profile=medium   -> node=node-dls4b (162.8ms)
bench-medium-4xbqt             profile=medium   -> node=node-dls4b (163.5ms)
bench-medium-k9xzm             profile=medium   -> node=node-qghf5 (163.5ms)
bench-medium-4vxx9             profile=medium   -> node=node-dls4b (164.3ms)
bench-large-hblh9              profile=large    -> node=node-dls4b (164.7ms)
bench-large-flhsr              profile=large    -> node=node-qghf5 (164.2ms)
bench-large-czrr8              profile=large    -> node=node-jp9v9 (165.5ms)
bench-large-h2kw6              profile=large    -> node=node-dls4b (164.3ms)
bench-large-pf52f              profile=large    -> node=node-qghf5 (163.8ms)
bench-large-gvbst              profile=large    -> node=node-dls4b (163.2ms)
bench-large-vxzv9              profile=large    -> node=node-dls4b (163.3ms)
bench-large-j4zjp              profile=large    -> node=node-qghf5 (162.9ms)
bench-xlarge-pdqks             profile=xlarge   -> node=node-dls4b (162.5ms)
bench-xlarge-ft69b             profile=xlarge   -> node=node-qghf5 (163.5ms)
bench-xlarge-jfvdd             profile=xlarge   -> node=node-dls4b (163.9ms)
bench-xlarge-mwxd4             profile=xlarge   -> node=node-jp9v9 (163.6ms)
bench-xlarge-jzwv6             profile=xlarge   -> node=node-dls4b (164.3ms)
bench-xlarge-r2n2v             profile=xlarge   -> node=node-qghf5 (163.9ms)
bench-xlarge-mmptq             profile=xlarge   -> node=node-dls4b (163.9ms)
bench-xlarge-sbkc4             profile=xlarge   -> node=node-qghf5 (57.0ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop1_0ms_results.csv
Per-pod scheduling latency: mean=155.93ms p95=165.52ms max=165.56ms
Метрики балансировки: results_timing_http/timing_http_noop1_0ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m6.509s
user	0m0.192s
sys	0m0.082s
=== [HTTP] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=50 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.6s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.2s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.3s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.5s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-s8zmq               profile=tiny     -> node=node-d4d8f (110.5ms)
bench-tiny-d5z4p               profile=tiny     -> node=node-d4d8f (216.7ms)
bench-tiny-t6m2q               profile=tiny     -> node=node-krn5w (217.9ms)
bench-tiny-tmf5x               profile=tiny     -> node=node-d4d8f (218.1ms)
bench-tiny-kwp9r               profile=tiny     -> node=node-79b5l (218.5ms)
bench-tiny-bzx92               profile=tiny     -> node=node-krn5w (110.3ms)
bench-tiny-kmq96               profile=tiny     -> node=node-krn5w (218.7ms)
bench-tiny-zgzrd               profile=tiny     -> node=node-d4d8f (216.9ms)
bench-small-lm7kt              profile=small    -> node=node-d4d8f (217.5ms)
bench-small-xwjct              profile=small    -> node=node-krn5w (216.7ms)
bench-small-jlmf2              profile=small    -> node=node-d4d8f (217.8ms)
bench-small-pqx79              profile=small    -> node=node-79b5l (217.5ms)
bench-small-7dxzm              profile=small    -> node=node-krn5w (217.5ms)
bench-small-89bht              profile=small    -> node=node-d4d8f (217.1ms)
bench-small-7427v              profile=small    -> node=node-d4d8f (218.0ms)
bench-small-4cgvb              profile=small    -> node=node-krn5w (218.1ms)
bench-medium-8pc9v             profile=medium   -> node=node-d4d8f (217.5ms)
bench-medium-8zfqg             profile=medium   -> node=node-krn5w (219.7ms)
bench-medium-wzmtk             profile=medium   -> node=node-d4d8f (214.5ms)
bench-medium-cxx87             profile=medium   -> node=node-79b5l (215.6ms)
bench-medium-zxrx9             profile=medium   -> node=node-krn5w (214.4ms)
bench-medium-p7dhr             profile=medium   -> node=node-d4d8f (214.9ms)
bench-medium-5mgnk             profile=medium   -> node=node-d4d8f (215.6ms)
bench-medium-rrvjx             profile=medium   -> node=node-krn5w (215.5ms)
bench-large-zwt8t              profile=large    -> node=node-d4d8f (216.1ms)
bench-large-ngbsc              profile=large    -> node=node-d4d8f (220.5ms)
bench-large-22svm              profile=large    -> node=node-krn5w (218.7ms)
bench-large-zmmj6              profile=large    -> node=node-79b5l (217.5ms)
bench-large-n428k              profile=large    -> node=node-d4d8f (216.5ms)
bench-large-plx8z              profile=large    -> node=node-krn5w (218.3ms)
bench-large-mtpdx              profile=large    -> node=node-d4d8f (216.0ms)
bench-large-59zjw              profile=large    -> node=node-d4d8f (217.4ms)
bench-xlarge-mltwz             profile=xlarge   -> node=node-krn5w (218.7ms)
bench-xlarge-n2cv4             profile=xlarge   -> node=node-d4d8f (218.1ms)
bench-xlarge-ttgw2             profile=xlarge   -> node=node-79b5l (217.4ms)
bench-xlarge-698tf             profile=xlarge   -> node=node-d4d8f (217.2ms)
bench-xlarge-jgfnf             profile=xlarge   -> node=node-krn5w (217.3ms)
bench-xlarge-l8t55             profile=xlarge   -> node=node-d4d8f (216.4ms)
bench-xlarge-r2l58             profile=xlarge   -> node=node-krn5w (218.4ms)
bench-xlarge-r2tq5             profile=xlarge   -> node=node-d4d8f (218.4ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop1_50ms_results.csv
Per-pod scheduling latency: mean=211.97ms p95=219.69ms max=220.48ms
Метрики балансировки: results_timing_http/timing_http_noop1_50ms_balance.json
  CPU utilization std=0.0901 mean=0.6080
  Memory utilization std=0.0893 mean=0.6020

real	0m8.763s
user	0m0.208s
sys	0m0.104s
=== [HTTP] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=100 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-m95ml               profile=tiny     -> node=node-jtrnm (165.3ms)
bench-tiny-q4jbq               profile=tiny     -> node=node-xzjsn (271.4ms)
bench-tiny-sxpkj               profile=tiny     -> node=node-xzjsn (271.5ms)
bench-tiny-zl8rl               profile=tiny     -> node=node-jtrnm (271.2ms)
bench-tiny-nzkl9               profile=tiny     -> node=node-z5z9l (272.8ms)
bench-tiny-wrbvk               profile=tiny     -> node=node-jtrnm (269.2ms)
bench-tiny-77gdv               profile=tiny     -> node=node-xzjsn (270.6ms)
bench-tiny-2cwkv               profile=tiny     -> node=node-xzjsn (163.5ms)
bench-small-xkf9s              profile=small    -> node=node-xzjsn (270.8ms)
bench-small-95vgj              profile=small    -> node=node-xzjsn (272.8ms)
bench-small-qjc54              profile=small    -> node=node-jtrnm (269.8ms)
bench-small-dn8ws              profile=small    -> node=node-xzjsn (272.8ms)
bench-small-j9f4h              profile=small    -> node=node-jtrnm (270.9ms)
bench-small-p5g2s              profile=small    -> node=node-z5z9l (271.1ms)
bench-small-m76tt              profile=small    -> node=node-xzjsn (272.2ms)
bench-small-skxcf              profile=small    -> node=node-xzjsn (271.8ms)
bench-medium-qkn6h             profile=medium   -> node=node-xzjsn (276.7ms)
bench-medium-8hp7m             profile=medium   -> node=node-jtrnm (270.9ms)
bench-medium-9n9p4             profile=medium   -> node=node-z5z9l (272.4ms)
bench-medium-zqwmg             profile=medium   -> node=node-jtrnm (271.1ms)
bench-medium-5c4ln             profile=medium   -> node=node-xzjsn (270.7ms)
bench-medium-45krc             profile=medium   -> node=node-xzjsn (271.0ms)
bench-medium-vkdtn             profile=medium   -> node=node-jtrnm (270.4ms)
bench-medium-rprlh             profile=medium   -> node=node-xzjsn (271.4ms)
bench-large-dkf6f              profile=large    -> node=node-xzjsn (271.1ms)
bench-large-x6xq5              profile=large    -> node=node-jtrnm (269.0ms)
bench-large-hwjjd              profile=large    -> node=node-z5z9l (271.5ms)
bench-large-b5rz8              profile=large    -> node=node-xzjsn (271.4ms)
bench-large-bpq47              profile=large    -> node=node-jtrnm (270.6ms)
bench-large-24dsc              profile=large    -> node=node-xzjsn (270.1ms)
bench-large-zlsxh              profile=large    -> node=node-xzjsn (162.7ms)
bench-large-b6kst              profile=large    -> node=node-jtrnm (272.6ms)
bench-xlarge-86n89             profile=xlarge   -> node=node-xzjsn (272.1ms)
bench-xlarge-c27l6             profile=xlarge   -> node=node-jtrnm (270.2ms)
bench-xlarge-th4kl             profile=xlarge   -> node=node-z5z9l (270.0ms)
bench-xlarge-xhs2s             profile=xlarge   -> node=node-xzjsn (271.1ms)
bench-xlarge-46w2w             profile=xlarge   -> node=node-xzjsn (271.1ms)
bench-xlarge-hz555             profile=xlarge   -> node=node-jtrnm (271.0ms)
bench-xlarge-bjdf4             profile=xlarge   -> node=node-xzjsn (270.1ms)
bench-xlarge-vpqjj             profile=xlarge   -> node=node-jtrnm (271.4ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop1_100ms_results.csv
Per-pod scheduling latency: mean=263.21ms p95=272.84ms max=276.72ms
Метрики балансировки: results_timing_http/timing_http_noop1_100ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m10.808s
user	0m0.250s
sys	0m0.096s
=== [HTTP] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=500 ===
[+] down 10/10
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.3s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.3s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-htwj9               profile=tiny     -> node=node-hwswd (539.5ms)
bench-tiny-dn4j8               profile=tiny     -> node=node-z6q9r (646.7ms)
bench-tiny-zbdbf               profile=tiny     -> node=node-z6q9r (645.4ms)
bench-tiny-hznz8               profile=tiny     -> node=node-hwswd (645.2ms)
bench-tiny-9x9wc               profile=tiny     -> node=node-hwswd (535.5ms)
bench-tiny-bflf2               profile=tiny     -> node=node-z6q9r (644.7ms)
bench-tiny-8cqwq               profile=tiny     -> node=node-g22qc (653.5ms)
bench-tiny-qjl72               profile=tiny     -> node=node-z6q9r (644.6ms)
bench-small-sxkdh              profile=small    -> node=node-z6q9r (645.3ms)
bench-small-4k7cm              profile=small    -> node=node-z6q9r (644.8ms)
bench-small-fbrzf              profile=small    -> node=node-hwswd (644.5ms)
bench-small-p8vd8              profile=small    -> node=node-hwswd (646.8ms)
bench-small-x7mms              profile=small    -> node=node-g22qc (646.1ms)
bench-small-6k72j              profile=small    -> node=node-z6q9r (642.6ms)
bench-small-vvknc              profile=small    -> node=node-z6q9r (644.8ms)
bench-small-2vwnn              profile=small    -> node=node-hwswd (644.0ms)
bench-medium-6sd55             profile=medium   -> node=node-z6q9r (646.7ms)
bench-medium-9zs6q             profile=medium   -> node=node-z6q9r (648.0ms)
bench-medium-65xjx             profile=medium   -> node=node-hwswd (645.9ms)
bench-medium-vqcj2             profile=medium   -> node=node-g22qc (536.7ms)
bench-medium-k9vhr             profile=medium   -> node=node-z6q9r (645.5ms)
bench-medium-4pl4z             profile=medium   -> node=node-hwswd (647.9ms)
bench-medium-vx89g             profile=medium   -> node=node-z6q9r (645.5ms)
bench-medium-n8c7n             profile=medium   -> node=node-hwswd (648.1ms)
bench-large-vh6zh              profile=large    -> node=node-z6q9r (645.0ms)
bench-large-drs2s              profile=large    -> node=node-hwswd (647.5ms)
bench-large-8m58q              profile=large    -> node=node-z6q9r (645.9ms)
bench-large-z995f              profile=large    -> node=node-g22qc (644.9ms)
bench-large-xc77d              profile=large    -> node=node-z6q9r (647.7ms)
bench-large-qx2w8              profile=large    -> node=node-hwswd (645.5ms)
bench-large-6pmj8              profile=large    -> node=node-z6q9r (645.9ms)
bench-large-x6skq              profile=large    -> node=node-hwswd (646.8ms)
bench-xlarge-fphr8             profile=xlarge   -> node=node-z6q9r (647.2ms)
bench-xlarge-f96gp             profile=xlarge   -> node=node-hwswd (643.6ms)
bench-xlarge-7cbt8             profile=xlarge   -> node=node-z6q9r (645.0ms)
bench-xlarge-c82tm             profile=xlarge   -> node=node-g22qc (654.4ms)
bench-xlarge-flbhl             profile=xlarge   -> node=node-z6q9r (640.4ms)
bench-xlarge-rxvrj             profile=xlarge   -> node=node-hwswd (643.8ms)
bench-xlarge-ndrtb             profile=xlarge   -> node=node-z6q9r (644.1ms)
bench-xlarge-hnbrr             profile=xlarge   -> node=node-z6q9r (647.0ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop1_500ms_results.csv
Per-pod scheduling latency: mean=637.82ms p95=653.48ms max=654.41ms
Метрики балансировки: results_timing_http/timing_http_noop1_500ms_balance.json
  CPU utilization std=0.0996 mean=0.6010
  Memory utilization std=0.0989 mean=0.5951

real	0m25.789s
user	0m0.425s
sys	0m0.194s
=== [HTTP] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=1500 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container ml-server                                         Removed                                                                                                               0.5s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-z8d7t               profile=tiny     -> node=node-lwkcp (1565.0ms)
bench-tiny-qrg9k               profile=tiny     -> node=node-lwkcp (1676.1ms)
bench-tiny-lmbpf               profile=tiny     -> node=node-tzrpn (1665.0ms)
bench-tiny-bhjcx               profile=tiny     -> node=node-g9vxn (1667.0ms)
bench-tiny-gl5t6               profile=tiny     -> node=node-tzrpn (1662.8ms)
bench-tiny-5djbf               profile=tiny     -> node=node-tzrpn (1667.2ms)
bench-tiny-vfcf4               profile=tiny     -> node=node-lwkcp (1665.0ms)
bench-tiny-fh7m6               profile=tiny     -> node=node-lwkcp (1666.0ms)
bench-small-82q9c              profile=small    -> node=node-lwkcp (1665.5ms)
bench-small-h47db              profile=small    -> node=node-lwkcp (1677.2ms)
bench-small-wv2vs              profile=small    -> node=node-tzrpn (1666.3ms)
bench-small-46zv7              profile=small    -> node=node-tzrpn (1661.6ms)
bench-small-fhjdq              profile=small    -> node=node-lwkcp (1662.7ms)
bench-small-6shv5              profile=small    -> node=node-g9vxn (1663.8ms)
bench-small-2pl62              profile=small    -> node=node-lwkcp (1664.8ms)
bench-small-fdbmz              profile=small    -> node=node-lwkcp (1664.0ms)
bench-medium-42lx8             profile=medium   -> node=node-tzrpn (1663.4ms)
bench-medium-vwktz             profile=medium   -> node=node-lwkcp (1558.7ms)
bench-medium-nhwwm             profile=medium   -> node=node-g9vxn (1554.1ms)
bench-medium-9kxw4             profile=medium   -> node=node-tzrpn (1670.3ms)
bench-medium-lrm9w             profile=medium   -> node=node-lwkcp (1664.3ms)
bench-medium-l6l47             profile=medium   -> node=node-lwkcp (1665.5ms)
bench-medium-pswq4             profile=medium   -> node=node-tzrpn (1662.6ms)
bench-medium-hcdxr             profile=medium   -> node=node-lwkcp (1666.6ms)
bench-large-8hs78              profile=large    -> node=node-lwkcp (1663.1ms)
bench-large-r958s              profile=large    -> node=node-tzrpn (1665.7ms)
bench-large-tr47l              profile=large    -> node=node-lwkcp (1662.7ms)
bench-large-l8srv              profile=large    -> node=node-g9vxn (1671.8ms)
bench-large-nz2zc              profile=large    -> node=node-tzrpn (1661.0ms)
bench-large-g5fr8              profile=large    -> node=node-lwkcp (1664.6ms)
bench-large-q8mgc              profile=large    -> node=node-lwkcp (1663.9ms)
bench-large-7dfm7              profile=large    -> node=node-tzrpn (1668.4ms)
bench-xlarge-fsxp6             profile=xlarge   -> node=node-lwkcp (1665.2ms)
bench-xlarge-wdr9p             profile=xlarge   -> node=node-tzrpn (1668.9ms)
bench-xlarge-tqkrt             profile=xlarge   -> node=node-lwkcp (1665.6ms)
bench-xlarge-qkk78             profile=xlarge   -> node=node-g9vxn (1668.6ms)
bench-xlarge-bbzsn             profile=xlarge   -> node=node-lwkcp (1668.2ms)
bench-xlarge-24lfz             profile=xlarge   -> node=node-tzrpn (1668.1ms)
bench-xlarge-rc7sc             profile=xlarge   -> node=node-lwkcp (1667.7ms)
bench-xlarge-2thk5             profile=xlarge   -> node=node-tzrpn (1668.6ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_http/timing_http_noop1_1500ms_results.csv
Per-pod scheduling latency: mean=1658.19ms p95=1676.07ms max=1677.15ms
Метрики балансировки: results_timing_http/timing_http_noop1_1500ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	1m6.613s
user	0m0.946s
sys	0m0.474s

=== СВОДКА (HTTP) ===
NOOP=0 DELAY=0ms: time=0m8.752s real_latency=[mean=65.05ms p95=73.47ms]
NOOP=0 DELAY=50ms: time=0m10.979s real_latency=[mean=119.55ms p95=128.86ms]
NOOP=0 DELAY=100ms: time=0m12.943s real_latency=[mean=169.43ms p95=180.88ms]
NOOP=0 DELAY=500ms: time=0m28.722s real_latency=[mean=569.34ms p95=581.07ms]
NOOP=0 DELAY=1500ms: time=1m8.654s real_latency=[mean=1568.02ms p95=1577.22ms]
NOOP=1 DELAY=0ms: time=0m6.509s real_latency=[mean=3.53ms p95=4.53ms]
NOOP=1 DELAY=50ms: time=0m8.763s real_latency=[mean=54.15ms p95=55.55ms]
NOOP=1 DELAY=100ms: time=0m10.808s real_latency=[mean=104.19ms p95=105.53ms]
NOOP=1 DELAY=500ms: time=0m25.789s real_latency=[mean=503.96ms p95=504.95ms]
NOOP=1 DELAY=1500ms: time=1m6.613s real_latency=[mean=1503.99ms p95=1505.15ms]

