bench-tiny-zh985               profile=tiny     -> node=node-zqbvf (108.7ms)
bench-tiny-989nx               profile=tiny     -> node=node-dkpdc (215.1ms)
bench-tiny-bzwwb               profile=tiny     -> node=node-zqbvf (217.5ms)
bench-tiny-snk6w               profile=tiny     -> node=node-hkvwb (215.9ms)
bench-tiny-x2vwt               profile=tiny     -> node=node-dkpdc (214.4ms)
bench-tiny-7gkbw               profile=tiny     -> node=node-zqbvf (218.4ms)
bench-tiny-jwg6w               profile=tiny     -> node=node-dkpdc (216.6ms)
bench-tiny-qxkmz               profile=tiny     -> node=node-zqbvf (216.2ms)
bench-small-8mf6q              profile=small    -> node=node-zqbvf (216.4ms)
bench-small-xkwh4              profile=small    -> node=node-zqbvf (215.5ms)
bench-small-rcscs              profile=small    -> node=node-dkpdc (216.2ms)
bench-small-hnrw4              profile=small    -> node=node-zqbvf (215.5ms)
bench-small-xrpx7              profile=small    -> node=node-dkpdc (216.5ms)
bench-small-45z67              profile=small    -> node=node-hkvwb (215.0ms)
bench-small-26wln              profile=small    -> node=node-zqbvf (215.4ms)
bench-small-sjb92              profile=small    -> node=node-zqbvf (215.5ms)
bench-medium-lsvtf             profile=medium   -> node=node-zqbvf (215.2ms)
bench-medium-f7fv7             profile=medium   -> node=node-dkpdc (215.1ms)
bench-medium-7n2xl             profile=medium   -> node=node-hkvwb (214.4ms)
bench-medium-rblmh             profile=medium   -> node=node-zqbvf (216.2ms)
bench-medium-s5khp             profile=medium   -> node=node-dkpdc (215.5ms)
bench-medium-pvgrm             profile=medium   -> node=node-zqbvf (215.2ms)
bench-medium-lf9zv             profile=medium   -> node=node-dkpdc (214.2ms)
bench-medium-gmrrd             profile=medium   -> node=node-zqbvf (216.5ms)
bench-large-fgr5m              profile=large    -> node=node-zqbvf (216.2ms)
bench-large-b4mwx              profile=large    -> node=node-dkpdc (215.7ms)
bench-large-n8t92              profile=large    -> node=node-zqbvf (216.3ms)
bench-large-qdpc4              profile=large    -> node=node-hkvwb (214.8ms)
bench-large-snkjv              profile=large    -> node=node-dkpdc (109.0ms)
bench-large-dgrdw              profile=large    -> node=node-zqbvf (215.7ms)
bench-large-zxlcd              profile=large    -> node=node-zqbvf (216.0ms)
bench-large-6prjb              profile=large    -> node=node-dkpdc (216.9ms)
bench-xlarge-p9qqz             profile=xlarge   -> node=node-zqbvf (269.6ms)
bench-xlarge-7cd6g             profile=xlarge   -> node=node-dkpdc (214.9ms)
bench-xlarge-rclrz             profile=xlarge   -> node=node-hkvwb (215.2ms)
bench-xlarge-pttl4             profile=xlarge   -> node=node-zqbvf (215.8ms)
bench-xlarge-rnxzr             profile=xlarge   -> node=node-zqbvf (217.4ms)
bench-xlarge-vtjrq             profile=xlarge   -> node=node-dkpdc (215.3ms)
bench-xlarge-2xsbd             profile=xlarge   -> node=node-zqbvf (216.7ms)
bench-xlarge-qfdrw             profile=xlarge   -> node=node-dkpdc (214.6ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_0ms_results.csv
Per-pod scheduling latency: mean=211.78ms p95=218.40ms max=269.65ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_0ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m8.738s
user	0m0.206s
sys	0m0.091s
=== [NNG] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=50 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-n726b               profile=tiny     -> node=node-zc8xp (269.4ms)
bench-tiny-b7dgs               profile=tiny     -> node=node-9l72w (377.2ms)
bench-tiny-7k4cv               profile=tiny     -> node=node-9l72w (376.2ms)
bench-tiny-s4gqq               profile=tiny     -> node=node-9l72w (378.1ms)
bench-tiny-j8h7v               profile=tiny     -> node=node-zc8xp (375.1ms)
bench-tiny-qrsp8               profile=tiny     -> node=node-9l72w (375.7ms)
bench-tiny-5h5cb               profile=tiny     -> node=node-7mf92 (376.4ms)
bench-tiny-4dk4c               profile=tiny     -> node=node-zc8xp (377.4ms)
bench-small-fcrcs              profile=small    -> node=node-9l72w (375.9ms)
bench-small-rqcnj              profile=small    -> node=node-zc8xp (375.8ms)
bench-small-nmzc6              profile=small    -> node=node-9l72w (377.7ms)
bench-small-mhklw              profile=small    -> node=node-zc8xp (376.2ms)
bench-small-8w9dl              profile=small    -> node=node-9l72w (375.5ms)
bench-small-ctq7z              profile=small    -> node=node-7mf92 (376.2ms)
bench-small-prnb9              profile=small    -> node=node-9l72w (376.4ms)
bench-small-r4hb2              profile=small    -> node=node-9l72w (378.0ms)
bench-medium-4rgtp             profile=medium   -> node=node-9l72w (376.1ms)
bench-medium-c9qhs             profile=medium   -> node=node-zc8xp (374.1ms)
bench-medium-ps8jb             profile=medium   -> node=node-7mf92 (375.9ms)
bench-medium-qrznj             profile=medium   -> node=node-9l72w (377.8ms)
bench-medium-5kg6j             profile=medium   -> node=node-zc8xp (377.0ms)
bench-medium-25qww             profile=medium   -> node=node-9l72w (375.8ms)
bench-medium-z22d7             profile=medium   -> node=node-zc8xp (377.9ms)
bench-medium-lwwhc             profile=medium   -> node=node-9l72w (375.7ms)
bench-large-cfwng              profile=large    -> node=node-9l72w (375.3ms)
bench-large-kltrf              profile=large    -> node=node-zc8xp (376.4ms)
bench-large-zh494              profile=large    -> node=node-9l72w (377.3ms)
bench-large-dnbjz              profile=large    -> node=node-7mf92 (379.0ms)
bench-large-sz9zj              profile=large    -> node=node-zc8xp (376.0ms)
bench-large-2g856              profile=large    -> node=node-9l72w (376.6ms)
bench-large-4mlnn              profile=large    -> node=node-9l72w (375.1ms)
bench-large-5nfpx              profile=large    -> node=node-zc8xp (378.1ms)
bench-xlarge-mmdc4             profile=xlarge   -> node=node-9l72w (430.9ms)
bench-xlarge-649gg             profile=xlarge   -> node=node-zc8xp (375.3ms)
bench-xlarge-58b6k             profile=xlarge   -> node=node-7mf92 (378.0ms)
bench-xlarge-rdpqc             profile=xlarge   -> node=node-9l72w (376.9ms)
bench-xlarge-ccq96             profile=xlarge   -> node=node-9l72w (375.3ms)
bench-xlarge-4dckn             profile=xlarge   -> node=node-zc8xp (373.5ms)
bench-xlarge-hs5p8             profile=xlarge   -> node=node-9l72w (375.3ms)
bench-xlarge-vtfkk             profile=xlarge   -> node=node-zc8xp (373.4ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_50ms_results.csv
Per-pod scheduling latency: mean=375.00ms p95=379.03ms max=430.91ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_50ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m15.263s
user	0m0.291s
sys	0m0.124s
=== [NNG] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=100 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.1s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.3s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-bm8zf               profile=tiny     -> node=node-jl7dk (537.7ms)
bench-tiny-h5jgc               profile=tiny     -> node=node-7s86n (535.5ms)
bench-tiny-47mcv               profile=tiny     -> node=node-7s86n (544.5ms)
bench-tiny-5zhgd               profile=tiny     -> node=node-cvxh2 (536.0ms)
bench-tiny-njscq               profile=tiny     -> node=node-7s86n (537.2ms)
bench-tiny-9m2g5               profile=tiny     -> node=node-7s86n (536.9ms)
bench-tiny-8rrgg               profile=tiny     -> node=node-jl7dk (539.5ms)
bench-tiny-g2lm7               profile=tiny     -> node=node-jl7dk (538.2ms)
bench-small-7pp9z              profile=small    -> node=node-7s86n (537.1ms)
bench-small-hrz9b              profile=small    -> node=node-jl7dk (537.1ms)
bench-small-qqpcf              profile=small    -> node=node-7s86n (533.7ms)
bench-small-f2rpl              profile=small    -> node=node-7s86n (538.3ms)
bench-small-s4kjg              profile=small    -> node=node-jl7dk (431.1ms)
bench-small-6vh95              profile=small    -> node=node-cvxh2 (537.7ms)
bench-small-grllj              profile=small    -> node=node-7s86n (537.8ms)
bench-small-7dq2q              profile=small    -> node=node-jl7dk (536.5ms)
bench-medium-t5nl7             profile=medium   -> node=node-7s86n (536.0ms)
bench-medium-dwkxl             profile=medium   -> node=node-jl7dk (536.9ms)
bench-medium-hmd8z             profile=medium   -> node=node-7s86n (537.6ms)
bench-medium-k84m8             profile=medium   -> node=node-cvxh2 (537.9ms)
bench-medium-tcwd5             profile=medium   -> node=node-7s86n (537.1ms)
bench-medium-zmp8t             profile=medium   -> node=node-jl7dk (537.8ms)
bench-medium-r8w9h             profile=medium   -> node=node-7s86n (538.3ms)
bench-medium-xlhzs             profile=medium   -> node=node-jl7dk (538.6ms)
bench-large-wggv9              profile=large    -> node=node-7s86n (535.6ms)
bench-large-vwmbr              profile=large    -> node=node-7s86n (536.9ms)
bench-large-g67pp              profile=large    -> node=node-jl7dk (429.0ms)
bench-large-6rppv              profile=large    -> node=node-cvxh2 (538.3ms)
bench-large-xdv77              profile=large    -> node=node-7s86n (536.6ms)
bench-large-vx4z6              profile=large    -> node=node-jl7dk (536.5ms)
bench-large-b7xxn              profile=large    -> node=node-7s86n (536.8ms)
bench-large-9d6ng              profile=large    -> node=node-7s86n (536.6ms)
bench-xlarge-8fwt7             profile=xlarge   -> node=node-jl7dk (596.0ms)
bench-xlarge-jffkl             profile=xlarge   -> node=node-7s86n (429.8ms)
bench-xlarge-82wgx             profile=xlarge   -> node=node-cvxh2 (538.3ms)
bench-xlarge-v4tbk             profile=xlarge   -> node=node-7s86n (537.5ms)
bench-xlarge-xbcrz             profile=xlarge   -> node=node-jl7dk (538.7ms)
bench-xlarge-mx62k             profile=xlarge   -> node=node-7s86n (537.1ms)
bench-xlarge-sfbdw             profile=xlarge   -> node=node-jl7dk (535.1ms)
bench-xlarge-dzl85             profile=xlarge   -> node=node-7s86n (537.3ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_100ms_results.csv
Per-pod scheduling latency: mean=530.78ms p95=544.50ms max=596.02ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_100ms_balance.json
  CPU utilization std=0.0901 mean=0.6080
  Memory utilization std=0.0893 mean=0.6020

real	0m21.503s
user	0m0.363s
sys	0m0.171s
=== [NNG] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=500 ===
[+] down 10/10
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.2s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-lsphk               profile=tiny     -> node=node-m6bwq (1771.9ms)
bench-tiny-wqzhq               profile=tiny     -> node=node-hj5pd (1719.7ms)
bench-tiny-r49ng               profile=tiny     -> node=node-hj5pd (1722.2ms)
bench-tiny-zpmnz               profile=tiny     -> node=node-hj5pd (1713.1ms)
bench-tiny-rpn2r               profile=tiny     -> node=node-m6bwq (1715.5ms)
bench-tiny-7wrhf               profile=tiny     -> node=node-z4t6c (1717.7ms)
bench-tiny-wqdjr               profile=tiny     -> node=node-hj5pd (1603.2ms)
bench-tiny-6v9wb               profile=tiny     -> node=node-m6bwq (1716.6ms)
bench-small-ftd6z              profile=small    -> node=node-hj5pd (1726.8ms)
bench-small-d4nl7              profile=small    -> node=node-hj5pd (1716.8ms)
bench-small-rpqcl              profile=small    -> node=node-m6bwq (1716.5ms)
bench-small-rhw6t              profile=small    -> node=node-z4t6c (1712.2ms)
bench-small-kh56q              profile=small    -> node=node-m6bwq (1713.5ms)
bench-small-zbn24              profile=small    -> node=node-hj5pd (1720.0ms)
bench-small-5r9h8              profile=small    -> node=node-hj5pd (1765.8ms)
bench-small-wnmdh              profile=small    -> node=node-m6bwq (1716.3ms)
bench-medium-fqsrr             profile=medium   -> node=node-hj5pd (1764.6ms)
bench-medium-vvzrr             profile=medium   -> node=node-hj5pd (1771.5ms)
bench-medium-mmxv2             profile=medium   -> node=node-m6bwq (1719.9ms)
bench-medium-mtr7m             profile=medium   -> node=node-z4t6c (1769.3ms)
bench-medium-gcwt4             profile=medium   -> node=node-hj5pd (1714.5ms)
bench-medium-thlwt             profile=medium   -> node=node-m6bwq (1715.8ms)
bench-medium-vllbv             profile=medium   -> node=node-hj5pd (1716.2ms)
bench-medium-gj48w             profile=medium   -> node=node-m6bwq (1765.0ms)
bench-large-f5m64              profile=large    -> node=node-hj5pd (1712.9ms)
bench-large-pwwmc              profile=large    -> node=node-m6bwq (1718.0ms)
bench-large-5lq4v              profile=large    -> node=node-hj5pd (1717.9ms)
bench-large-m8mn4              profile=large    -> node=node-z4t6c (1722.3ms)
bench-large-84jqx              profile=large    -> node=node-hj5pd (1758.0ms)
bench-large-d5f2g              profile=large    -> node=node-m6bwq (1708.2ms)
bench-large-lgcsc              profile=large    -> node=node-hj5pd (1715.5ms)
bench-large-gvt5m              profile=large    -> node=node-m6bwq (1717.0ms)
bench-xlarge-d9lwc             profile=xlarge   -> node=node-hj5pd (1766.7ms)
bench-xlarge-x4t6q             profile=xlarge   -> node=node-m6bwq (1714.9ms)
bench-xlarge-fls5b             profile=xlarge   -> node=node-hj5pd (1715.2ms)
bench-xlarge-lvf8h             profile=xlarge   -> node=node-z4t6c (1712.2ms)
bench-xlarge-jxkwp             profile=xlarge   -> node=node-hj5pd (1712.8ms)
bench-xlarge-qtp8q             profile=xlarge   -> node=node-m6bwq (1717.5ms)
bench-xlarge-gdxtf             profile=xlarge   -> node=node-hj5pd (1717.3ms)
bench-xlarge-qjq4z             profile=xlarge   -> node=node-hj5pd (1712.4ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_500ms_results.csv
Per-pod scheduling latency: mean=1723.58ms p95=1771.52ms max=1771.88ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_500ms_balance.json
  CPU utilization std=0.0996 mean=0.6010
  Memory utilization std=0.0989 mean=0.5951

real	1m9.203s
user	0m0.952s
sys	0m0.465s
=== [NNG] NOOP_MODE=0 ARTIFICIAL_DELAY_MS=1500 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.6s
bench-tiny-djsnv               profile=tiny     -> node=node-8fxcs (2153.0ms)
bench-tiny-dp9gj               profile=tiny     -> node=node-bqzxq (2145.4ms)
bench-tiny-64l4r               profile=tiny     -> node=node-bqzxq (2151.1ms)
bench-tiny-m69g4               profile=tiny     -> node=node-8fxcs (2144.0ms)
bench-tiny-b9694               profile=tiny     -> node=node-bqzxq (2151.2ms)
bench-tiny-x8mwq               profile=tiny     -> node=node-bqzxq (2157.0ms)
bench-tiny-bgv6g               profile=tiny     -> node=node-8fxcs (2148.1ms)
bench-tiny-n7lk2               profile=tiny     -> node=node-2bmxc (2143.2ms)
bench-small-8pjfz              profile=small    -> node=node-bqzxq (2147.8ms)
bench-small-26rsr              profile=small    -> node=node-8fxcs (2143.2ms)
bench-small-p2pcb              profile=small    -> node=node-bqzxq (2150.5ms)
bench-small-p7zkv              profile=small    -> node=node-2bmxc (2150.5ms)
bench-small-8phxd              profile=small    -> node=node-bqzxq (2146.0ms)
bench-small-hqt5v              profile=small    -> node=node-8fxcs (2149.5ms)
bench-small-cbzj6              profile=small    -> node=node-bqzxq (2144.0ms)
bench-small-b6pj8              profile=small    -> node=node-bqzxq (2150.2ms)
bench-medium-7lx8n             profile=medium   -> node=node-8fxcs (2156.5ms)
bench-medium-zbhkt             profile=medium   -> node=node-bqzxq (2142.0ms)
bench-medium-xpk2v             profile=medium   -> node=node-2bmxc (2151.9ms)
bench-medium-dswpr             profile=medium   -> node=node-8fxcs (2140.0ms)
bench-medium-hw8zg             profile=medium   -> node=node-bqzxq (2143.0ms)
bench-medium-nlxzw             profile=medium   -> node=node-bqzxq (2141.1ms)
bench-medium-kthsr             profile=medium   -> node=node-8fxcs (2146.0ms)
bench-medium-5w9jb             profile=medium   -> node=node-bqzxq (2146.7ms)
bench-large-fg964              profile=large    -> node=node-bqzxq (2147.8ms)
bench-large-ktgsl              profile=large    -> node=node-8fxcs (2143.9ms)
bench-large-wwp2h              profile=large    -> node=node-bqzxq (2142.2ms)
bench-large-px89h              profile=large    -> node=node-2bmxc (2148.2ms)
bench-large-8cg62              profile=large    -> node=node-8fxcs (2147.8ms)
bench-large-nrthb              profile=large    -> node=node-bqzxq (2143.8ms)
bench-large-x65t2              profile=large    -> node=node-bqzxq (2143.6ms)
bench-large-wsq6s              profile=large    -> node=node-8fxcs (2146.1ms)
bench-xlarge-w69nn             profile=xlarge   -> node=node-bqzxq (2143.0ms)
bench-xlarge-6n42z             profile=xlarge   -> node=node-8fxcs (2146.0ms)
bench-xlarge-crmcq             profile=xlarge   -> node=node-bqzxq (2146.1ms)
bench-xlarge-99s2j             profile=xlarge   -> node=node-2bmxc (2142.6ms)
bench-xlarge-jcnkw             profile=xlarge   -> node=node-bqzxq (2139.8ms)
bench-xlarge-scvqg             profile=xlarge   -> node=node-8fxcs (2138.9ms)
bench-xlarge-92j96             profile=xlarge   -> node=node-bqzxq (2141.5ms)
bench-xlarge-fl79x             profile=xlarge   -> node=node-8fxcs (2146.3ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop0_1500ms_results.csv
Per-pod scheduling latency: mean=2146.23ms p95=2156.54ms max=2156.96ms
Метрики балансировки: results_timing_nng/timing_nng_noop0_1500ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	1m26.121s
user	0m1.130s
sys	0m0.623s
=== [NNG] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=0 ===
[+] down 10/10
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.3s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.5s
 ✔ Container simulator-server                                  Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.6s
bench-tiny-777gm               profile=tiny     -> node=node-d77v8 (164.3ms)
bench-tiny-2gkj4               profile=tiny     -> node=node-d77v8 (164.1ms)
bench-tiny-b7psr               profile=tiny     -> node=node-wtfrj (164.3ms)
bench-tiny-7l8sn               profile=tiny     -> node=node-d77v8 (167.0ms)
bench-tiny-nh8mt               profile=tiny     -> node=node-d77v8 (164.5ms)
bench-tiny-ldpr5               profile=tiny     -> node=node-d5f67 (164.1ms)
bench-tiny-grk4r               profile=tiny     -> node=node-wtfrj (164.1ms)
bench-tiny-lktm8               profile=tiny     -> node=node-wtfrj (56.4ms)
bench-small-6nfkw              profile=small    -> node=node-d77v8 (164.0ms)
bench-small-9tvcq              profile=small    -> node=node-wtfrj (164.7ms)
bench-small-t75gk              profile=small    -> node=node-d77v8 (164.1ms)
bench-small-k6sgp              profile=small    -> node=node-wtfrj (164.6ms)
bench-small-fpcmm              profile=small    -> node=node-d77v8 (164.4ms)
bench-small-hfq6j              profile=small    -> node=node-d5f67 (164.7ms)
bench-small-twwh6              profile=small    -> node=node-d77v8 (163.8ms)
bench-small-8xv27              profile=small    -> node=node-d77v8 (163.5ms)
bench-medium-742wj             profile=medium   -> node=node-d77v8 (164.1ms)
bench-medium-pswcs             profile=medium   -> node=node-wtfrj (163.3ms)
bench-medium-dqbvp             profile=medium   -> node=node-d5f67 (162.9ms)
bench-medium-r59x5             profile=medium   -> node=node-d77v8 (162.7ms)
bench-medium-qj8p6             profile=medium   -> node=node-wtfrj (165.3ms)
bench-medium-6f6pr             profile=medium   -> node=node-d77v8 (164.8ms)
bench-medium-6vwbw             profile=medium   -> node=node-wtfrj (162.7ms)
bench-medium-fbsgh             profile=medium   -> node=node-d77v8 (172.5ms)
bench-large-xfhq7              profile=large    -> node=node-d77v8 (163.1ms)
bench-large-6l5z6              profile=large    -> node=node-wtfrj (163.5ms)
bench-large-sg98g              profile=large    -> node=node-d5f67 (163.2ms)
bench-large-57f75              profile=large    -> node=node-d77v8 (163.1ms)
bench-large-2h22s              profile=large    -> node=node-wtfrj (164.1ms)
bench-large-4jv2d              profile=large    -> node=node-d77v8 (164.2ms)
bench-large-rn82t              profile=large    -> node=node-d77v8 (163.7ms)
bench-large-bnnp7              profile=large    -> node=node-wtfrj (163.9ms)
bench-xlarge-r2www             profile=xlarge   -> node=node-d77v8 (163.9ms)
bench-xlarge-sxv6k             profile=xlarge   -> node=node-wtfrj (54.8ms)
bench-xlarge-4k57v             profile=xlarge   -> node=node-d5f67 (162.4ms)
bench-xlarge-w9plj             profile=xlarge   -> node=node-d77v8 (164.7ms)
bench-xlarge-7cktr             profile=xlarge   -> node=node-d77v8 (164.2ms)
bench-xlarge-hd7c8             profile=xlarge   -> node=node-wtfrj (164.3ms)
bench-xlarge-gwt4j             profile=xlarge   -> node=node-d77v8 (163.7ms)
bench-xlarge-v85mg             profile=xlarge   -> node=node-wtfrj (163.3ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop1_0ms_results.csv
Per-pod scheduling latency: mean=158.77ms p95=167.04ms max=172.51ms
Метрики балансировки: results_timing_nng/timing_nng_noop1_0ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m6.633s
user	0m0.193s
sys	0m0.080s
=== [NNG] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=50 ===
[+] down 10/10
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.1s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.6s
bench-tiny-tcl4b               profile=tiny     -> node=node-c5vmd (217.5ms)
bench-tiny-frptl               profile=tiny     -> node=node-qb4qf (324.9ms)
bench-tiny-mpsqp               profile=tiny     -> node=node-c5vmd (325.0ms)
bench-tiny-ssgff               profile=tiny     -> node=node-c5vmd (325.5ms)
bench-tiny-5h2pc               profile=tiny     -> node=node-c5vmd (218.0ms)
bench-tiny-sb89s               profile=tiny     -> node=node-9vhxl (323.5ms)
bench-tiny-5p4q2               profile=tiny     -> node=node-qb4qf (323.7ms)
bench-tiny-mvvd6               profile=tiny     -> node=node-qb4qf (322.6ms)
bench-small-db7b8              profile=small    -> node=node-c5vmd (321.2ms)
bench-small-klqsp              profile=small    -> node=node-qb4qf (325.0ms)
bench-small-hjf6f              profile=small    -> node=node-c5vmd (323.4ms)
bench-small-mkrx4              profile=small    -> node=node-qb4qf (270.1ms)
bench-small-mmjj9              profile=small    -> node=node-c5vmd (323.8ms)
bench-small-d67kn              profile=small    -> node=node-9vhxl (323.8ms)
bench-small-9s28q              profile=small    -> node=node-c5vmd (324.6ms)
bench-small-6zrxv              profile=small    -> node=node-qb4qf (323.7ms)
bench-medium-khmn7             profile=medium   -> node=node-c5vmd (323.5ms)
bench-medium-942b5             profile=medium   -> node=node-c5vmd (322.9ms)
bench-medium-g655d             profile=medium   -> node=node-qb4qf (322.4ms)
bench-medium-qvckb             profile=medium   -> node=node-9vhxl (322.5ms)
bench-medium-gmkvr             profile=medium   -> node=node-qb4qf (323.7ms)
bench-medium-7j9sz             profile=medium   -> node=node-c5vmd (323.9ms)
bench-medium-9szpd             profile=medium   -> node=node-c5vmd (324.0ms)
bench-medium-27nrl             profile=medium   -> node=node-qb4qf (322.4ms)
bench-large-zhwp2              profile=large    -> node=node-c5vmd (323.7ms)
bench-large-8vw56              profile=large    -> node=node-c5vmd (323.6ms)
bench-large-vmfzd              profile=large    -> node=node-qb4qf (324.2ms)
bench-large-6sfdq              profile=large    -> node=node-9vhxl (323.6ms)
bench-large-kqb44              profile=large    -> node=node-c5vmd (322.7ms)
bench-large-64764              profile=large    -> node=node-qb4qf (269.9ms)
bench-large-rgzvp              profile=large    -> node=node-c5vmd (322.9ms)
bench-large-dkl9l              profile=large    -> node=node-qb4qf (324.2ms)
bench-xlarge-blv69             profile=xlarge   -> node=node-c5vmd (323.8ms)
bench-xlarge-k9dqq             profile=xlarge   -> node=node-c5vmd (325.0ms)
bench-xlarge-crzd8             profile=xlarge   -> node=node-qb4qf (270.2ms)
bench-xlarge-rjnpk             profile=xlarge   -> node=node-9vhxl (322.8ms)
bench-xlarge-x6bs7             profile=xlarge   -> node=node-c5vmd (322.6ms)
bench-xlarge-xpnht             profile=xlarge   -> node=node-qb4qf (325.8ms)
bench-xlarge-wsd5n             profile=xlarge   -> node=node-c5vmd (271.9ms)
bench-xlarge-pwtxr             profile=xlarge   -> node=node-c5vmd (326.1ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop1_50ms_results.csv
Per-pod scheduling latency: mean=313.11ms p95=325.81ms max=326.09ms
Метрики балансировки: results_timing_nng/timing_nng_noop1_50ms_balance.json
  CPU utilization std=0.0996 mean=0.6010
  Memory utilization std=0.0989 mean=0.5951

real	0m12.811s
user	0m0.262s
sys	0m0.116s
=== [NNG] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=100 ===
[+] down 10/10
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.5s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-frontend                                Started                                                                                                               0.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.6s
bench-tiny-nslgv               profile=tiny     -> node=node-wvjcw (321.8ms)
bench-tiny-fnwhz               profile=tiny     -> node=node-rlrb2 (430.3ms)
bench-tiny-297mj               profile=tiny     -> node=node-rlrb2 (431.3ms)
bench-tiny-sksj5               profile=tiny     -> node=node-rlrb2 (430.9ms)
bench-tiny-whz2d               profile=tiny     -> node=node-fhxrb (429.0ms)
bench-tiny-svd5q               profile=tiny     -> node=node-wvjcw (429.5ms)
bench-tiny-9qdlz               profile=tiny     -> node=node-rlrb2 (326.2ms)
bench-tiny-qvgfl               profile=tiny     -> node=node-wvjcw (430.5ms)
bench-small-8dfpp              profile=small    -> node=node-rlrb2 (429.9ms)
bench-small-xtwws              profile=small    -> node=node-rlrb2 (430.9ms)
bench-small-8xfz4              profile=small    -> node=node-wvjcw (433.7ms)
bench-small-7wfz4              profile=small    -> node=node-wvjcw (432.7ms)
bench-small-7cmnx              profile=small    -> node=node-fhxrb (433.7ms)
bench-small-zjdnd              profile=small    -> node=node-rlrb2 (432.2ms)
bench-small-bxg4d              profile=small    -> node=node-rlrb2 (429.9ms)
bench-small-dzkmr              profile=small    -> node=node-rlrb2 (430.3ms)
bench-medium-xb82n             profile=medium   -> node=node-rlrb2 (429.0ms)
bench-medium-4486d             profile=medium   -> node=node-wvjcw (437.9ms)
bench-medium-4hkvm             profile=medium   -> node=node-fhxrb (430.2ms)
bench-medium-mlm57             profile=medium   -> node=node-wvjcw (429.8ms)
bench-medium-f8nv5             profile=medium   -> node=node-rlrb2 (430.7ms)
bench-medium-gsn9n             profile=medium   -> node=node-rlrb2 (430.3ms)
bench-medium-gm4cs             profile=medium   -> node=node-wvjcw (431.3ms)
bench-medium-dr6rk             profile=medium   -> node=node-rlrb2 (482.0ms)
bench-large-q6w4h              profile=large    -> node=node-rlrb2 (430.6ms)
bench-large-98z65              profile=large    -> node=node-wvjcw (436.0ms)
bench-large-prk7j              profile=large    -> node=node-rlrb2 (431.0ms)
bench-large-656gt              profile=large    -> node=node-fhxrb (430.2ms)
bench-large-44jrz              profile=large    -> node=node-wvjcw (483.3ms)
bench-large-hzxx2              profile=large    -> node=node-rlrb2 (430.4ms)
bench-large-mcs6x              profile=large    -> node=node-rlrb2 (430.3ms)
bench-large-5bzzm              profile=large    -> node=node-wvjcw (430.5ms)
bench-xlarge-hgrlj             profile=xlarge   -> node=node-rlrb2 (489.7ms)
bench-xlarge-sbnf9             profile=xlarge   -> node=node-wvjcw (430.2ms)
bench-xlarge-qc7kd             profile=xlarge   -> node=node-rlrb2 (431.1ms)
bench-xlarge-j6rmz             profile=xlarge   -> node=node-fhxrb (432.3ms)
bench-xlarge-n6nbs             profile=xlarge   -> node=node-rlrb2 (430.9ms)
bench-xlarge-2ktsz             profile=xlarge   -> node=node-wvjcw (430.0ms)
bench-xlarge-nd4z9             profile=xlarge   -> node=node-rlrb2 (431.9ms)
bench-xlarge-s6g24             profile=xlarge   -> node=node-wvjcw (430.2ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop1_100ms_results.csv
Per-pod scheduling latency: mean=429.81ms p95=483.28ms max=489.70ms
Метрики балансировки: results_timing_nng/timing_nng_noop1_100ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	0m17.463s
user	0m0.317s
sys	0m0.144s
=== [NNG] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=500 ===
[+] down 10/10
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.2s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.1s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.3s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.5s
bench-tiny-qnhpg               profile=tiny     -> node=node-pl2rk (1669.1ms)
bench-tiny-tmttk               profile=tiny     -> node=node-tctp8 (1557.6ms)
bench-tiny-52kcs               profile=tiny     -> node=node-pl2rk (1661.4ms)
bench-tiny-srkgb               profile=tiny     -> node=node-nnt5p (1663.1ms)
bench-tiny-dfjst               profile=tiny     -> node=node-tctp8 (1665.4ms)
bench-tiny-dw7cr               profile=tiny     -> node=node-tctp8 (1673.9ms)
bench-tiny-6wk8g               profile=tiny     -> node=node-pl2rk (1661.7ms)
bench-tiny-kh7pv               profile=tiny     -> node=node-pl2rk (1674.7ms)
bench-small-ncr86              profile=small    -> node=node-pl2rk (1555.7ms)
bench-small-srsgx              profile=small    -> node=node-pl2rk (1664.3ms)
bench-small-f2qr5              profile=small    -> node=node-tctp8 (1666.9ms)
bench-small-xr99l              profile=small    -> node=node-nnt5p (1663.4ms)
bench-small-kdqdt              profile=small    -> node=node-tctp8 (1555.0ms)
bench-small-dcwpz              profile=small    -> node=node-pl2rk (1673.2ms)
bench-small-f4qbb              profile=small    -> node=node-pl2rk (1660.1ms)
bench-small-f57ff              profile=small    -> node=node-pl2rk (1661.8ms)
bench-medium-g9m2s             profile=medium   -> node=node-tctp8 (1664.2ms)
bench-medium-xhztf             profile=medium   -> node=node-pl2rk (1661.8ms)
bench-medium-7l798             profile=medium   -> node=node-nnt5p (1660.0ms)
bench-medium-6j9x6             profile=medium   -> node=node-pl2rk (1665.0ms)
bench-medium-rvpg5             profile=medium   -> node=node-tctp8 (1661.7ms)
bench-medium-rtnbb             profile=medium   -> node=node-pl2rk (1662.9ms)
bench-medium-kh228             profile=medium   -> node=node-tctp8 (1662.1ms)
bench-medium-xvbbs             profile=medium   -> node=node-pl2rk (1661.8ms)
bench-large-rvdzk              profile=large    -> node=node-pl2rk (1663.0ms)
bench-large-mjsph              profile=large    -> node=node-tctp8 (1661.3ms)
bench-large-7hvbf              profile=large    -> node=node-pl2rk (1663.9ms)
bench-large-hpjpc              profile=large    -> node=node-nnt5p (1661.5ms)
bench-large-lpwrd              profile=large    -> node=node-tctp8 (1665.1ms)
bench-large-r8dpf              profile=large    -> node=node-pl2rk (1662.9ms)
bench-large-nd8v7              profile=large    -> node=node-pl2rk (1664.1ms)
bench-large-khqkc              profile=large    -> node=node-tctp8 (1661.0ms)
bench-xlarge-h7xct             profile=xlarge   -> node=node-pl2rk (1668.4ms)
bench-xlarge-84mcv             profile=xlarge   -> node=node-tctp8 (1662.9ms)
bench-xlarge-qx7jn             profile=xlarge   -> node=node-pl2rk (1662.0ms)
bench-xlarge-t5bcq             profile=xlarge   -> node=node-nnt5p (1664.3ms)
bench-xlarge-rvkrw             profile=xlarge   -> node=node-pl2rk (1666.2ms)
bench-xlarge-lh4z8             profile=xlarge   -> node=node-tctp8 (1661.7ms)
bench-xlarge-hcf4x             profile=xlarge   -> node=node-pl2rk (1657.7ms)
bench-xlarge-nstnc             profile=xlarge   -> node=node-tctp8 (1674.1ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop1_500ms_results.csv
Per-pod scheduling latency: mean=1656.17ms p95=1674.08ms max=1674.71ms
Метрики балансировки: results_timing_nng/timing_nng_noop1_500ms_balance.json
  CPU utilization std=0.0959 mean=0.6132
  Memory utilization std=0.0949 mean=0.6072

real	1m6.537s
user	0m0.921s
sys	0m0.467s
=== [NNG] NOOP_MODE=1 ARTIFICIAL_DELAY_MS=1500 ===
[+] down 10/10
 ✔ Container simulator-scheduler                               Removed                                                                                                               0.2s
 ✔ Container simulator-server                                  Removed                                                                                                               0.2s
 ✔ Container simulator-frontend                                Removed                                                                                                               0.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Removed                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Removed                                                                                                              10.3s
 ✔ Container ml-server                                         Removed                                                                                                              10.2s
 ✔ Network kube-scheduler-simulator_default                    Removed                                                                                                               0.1s
 ✔ Volume kube-scheduler-simulator_conf                        Removed                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Removed                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Removed                                                                                                               0.2s
[+] up 10/10
 ✔ Network kube-scheduler-simulator_default                    Created                                                                                                               0.0s
 ✔ Network kube-scheduler-simulator_simulator-internal-network Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_simulator-etcd-data         Created                                                                                                               0.0s
 ✔ Volume kube-scheduler-simulator_conf                        Created                                                                                                               0.0s
 ✔ Container simulator-cluster                                 Started                                                                                                               0.3s
 ✔ Container kube-scheduler-simulator-init-container-1         Started                                                                                                               0.4s
 ✔ Container simulator-server                                  Started                                                                                                               0.4s
 ✔ Container ml-server                                         Started                                                                                                               0.3s
 ✔ Container simulator-frontend                                Started                                                                                                               0.4s
 ✔ Container simulator-scheduler                               Started                                                                                                               0.6s
bench-tiny-6vv8h               profile=tiny     -> node=node-wql7j (2065.9ms)
bench-tiny-gffhl               profile=tiny     -> node=node-phdfn (2156.5ms)
bench-tiny-jbljj               profile=tiny     -> node=node-phdfn (2146.9ms)
bench-tiny-8zqbb               profile=tiny     -> node=node-rk287 (2148.3ms)
bench-tiny-bj6fz               profile=tiny     -> node=node-wql7j (2132.9ms)
bench-tiny-r8fsc               profile=tiny     -> node=node-wql7j (2147.1ms)
bench-tiny-m2bxk               profile=tiny     -> node=node-phdfn (2151.0ms)
bench-tiny-lz5n6               profile=tiny     -> node=node-phdfn (2145.5ms)
bench-small-n8j5w              profile=small    -> node=node-phdfn (2147.5ms)
bench-small-lx452              profile=small    -> node=node-phdfn (2143.7ms)
bench-small-ntdrl              profile=small    -> node=node-wql7j (2157.1ms)
bench-small-25h7g              profile=small    -> node=node-phdfn (2145.7ms)
bench-small-k4wdh              profile=small    -> node=node-wql7j (2144.5ms)
bench-small-7qjkx              profile=small    -> node=node-rk287 (2144.5ms)
bench-small-hcfnd              profile=small    -> node=node-phdfn (2151.0ms)
bench-small-6x5fn              profile=small    -> node=node-wql7j (2147.2ms)
bench-medium-7c2pw             profile=medium   -> node=node-phdfn (2145.9ms)
bench-medium-8zv6s             profile=medium   -> node=node-phdfn (2040.5ms)
bench-medium-6wzxd             profile=medium   -> node=node-wql7j (2138.0ms)
bench-medium-pfrzm             profile=medium   -> node=node-rk287 (2144.3ms)
bench-medium-875sf             profile=medium   -> node=node-wql7j (2144.6ms)
bench-medium-rlwmq             profile=medium   -> node=node-phdfn (2144.1ms)
bench-medium-xcm2c             profile=medium   -> node=node-phdfn (2144.3ms)
bench-medium-fr47d             profile=medium   -> node=node-wql7j (2144.4ms)
bench-large-fdcvm              profile=large    -> node=node-phdfn (2142.5ms)
bench-large-2l4rq              profile=large    -> node=node-phdfn (2144.8ms)
bench-large-2hgpl              profile=large    -> node=node-wql7j (2145.9ms)
bench-large-k9fl5              profile=large    -> node=node-rk287 (2142.0ms)
bench-large-vm9tj              profile=large    -> node=node-phdfn (2147.5ms)
bench-large-hfbnz              profile=large    -> node=node-wql7j (2144.2ms)
bench-large-rwsvx              profile=large    -> node=node-phdfn (2144.2ms)
bench-large-c52s5              profile=large    -> node=node-phdfn (2153.6ms)
bench-xlarge-5b5nr             profile=xlarge   -> node=node-wql7j (2147.0ms)
bench-xlarge-x5xhn             profile=xlarge   -> node=node-phdfn (2151.3ms)
bench-xlarge-5pzs4             profile=xlarge   -> node=node-rk287 (2142.1ms)
bench-xlarge-x2h9p             profile=xlarge   -> node=node-phdfn (2146.3ms)
bench-xlarge-b7m9w             profile=xlarge   -> node=node-wql7j (2150.9ms)
bench-xlarge-shzrh             profile=xlarge   -> node=node-phdfn (2041.8ms)
bench-xlarge-hjpft             profile=xlarge   -> node=node-wql7j (2035.8ms)
bench-xlarge-bsxxf             profile=xlarge   -> node=node-phdfn (2143.9ms)

Готово: 40/40 подов запланировано.
Результаты: results_timing_nng/timing_nng_noop1_1500ms_results.csv
Per-pod scheduling latency: mean=2136.13ms p95=2156.48ms max=2157.07ms
Метрики балансировки: results_timing_nng/timing_nng_noop1_1500ms_balance.json
  CPU utilization std=0.0901 mean=0.6080
  Memory utilization std=0.0893 mean=0.6020

real	1m25.716s
user	0m1.151s
sys	0m0.595s

=== СВОДКА (NNG) ===
NOOP=0 DELAY=0ms: time=0m8.738s real_latency=[mean=52.96ms p95=80.33ms]
NOOP=0 DELAY=50ms: time=0m15.263s real_latency=[mean=156.75ms p95=237.27ms]
NOOP=0 DELAY=100ms: time=0m21.503s real_latency=[mean=256.77ms p95=387.67ms]
NOOP=0 DELAY=500ms: time=1m9.203s real_latency=[mean=1057.11ms p95=1586.93ms]
NOOP=0 DELAY=1500ms: time=1m26.121s real_latency=[mean=1996.81ms p95=2001.13ms]
NOOP=1 DELAY=0ms: time=0m6.633s real_latency=[mean=0.56ms p95=0.79ms]
NOOP=1 DELAY=50ms: time=0m12.811s real_latency=[mean=101.26ms p95=151.94ms]
NOOP=1 DELAY=100ms: time=0m17.463s real_latency=[mean=201.20ms p95=301.88ms]
NOOP=1 DELAY=500ms: time=1m6.537s real_latency=[mean=1001.19ms p95=1501.81ms]
NOOP=1 DELAY=1500ms: time=1m25.716s real_latency=[mean=1996.43ms p95=2001.07ms]

