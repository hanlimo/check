package exporter

import (
	"fmt"
	"github.com/hanlimo/check/kubeconf"
	"github.com/prometheus/client_golang/prometheus"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type countCollector struct {
	countMetric *prometheus.Desc
}

func NewCheckCollector(count int32) *countCollector {
	return &countCollector{
		countMetric: prometheus.NewDesc("countMetric",
			"countMetric指标用来显示Pod重启",
			nil, nil,
		),
	}
}

//collector必要函数
func (collector *countCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.countMetric
}

func getPodCount(clientSet kubernetes.Clientset) (counts []int32) {
	pods, err := clientSet.CoreV1().Pods("").List(meta.ListOptions{})
	if err != nil{
		fmt.Printf("获取参数失败")
	}
	for _, pod := range pods.Items {
		counts = append(counts, pod.Status.ContainerStatuses[0].RestartCount)
	}
	return counts
}
//collector必要函数
func (collector *countCollector) Collect(ch chan<- prometheus.Metric) {

	clientSet, _ := kubeconf.KubeConfig_init()
	metricValues := getPodCount(clientSet)

	//此处应该有问题，没有对协程做判断，需要后续优化
	for count := range metricValues {
		ch <- prometheus.MustNewConstMetric(collector.countMetric, prometheus.CounterValue, float64(count))
	}


}


