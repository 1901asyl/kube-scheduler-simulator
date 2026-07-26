package main


import (
	"fmt"
	"os"

	"k8s.io/component-base/cli"
	_ "k8s.io/component-base/logs/json/register"
	_ "k8s.io/component-base/metrics/prometheus/clientgo"
	_ "k8s.io/component-base/metrics/prometheus/version"
	"k8s.io/klog"

	"sigs.k8s.io/kube-scheduler-simulator/simulator/docs/sample/ai_score"
	"sigs.k8s.io/kube-scheduler-simulator/simulator/docs/sample/ai_score_heuristic"
	"sigs.k8s.io/kube-scheduler-simulator/simulator/pkg/debuggablescheduler"
)

func main() {
	command, cancelFn, err := debuggablescheduler.NewSchedulerCommand(
		debuggablescheduler.WithPlugin(ai_score.Name, ai_score.New),
		debuggablescheduler.WithPlugin(ai_score_heuristic.Name, ai_score_heuristic.New),
	)
	if err != nil {
		klog.Info(fmt.Sprintf("failed to build the debuggablescheduler command: %+v", err))
		os.Exit(1)
	}
	code := cli.Run(command)

	cancelFn()
	os.Exit(code)
}



