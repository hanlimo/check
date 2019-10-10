package ops

import (
	"fmt"
	"github.com/hanlimo/check/kubeconf"
	"github.com/prometheus/client_golang/prometheus"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type countCollector struct {
	countMetric *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newCheckCollector(count int32) *countCollector {
	return &countCollector{
		countMetric: prometheus.NewDesc("countMetric",
			"countMetric指标用来显示Pod重启",
			nil, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *countCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
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

//Collect implements required collect function for all promehteus collectors
func (collector *countCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	clientSet, _ := kubeconf.Kubeconfig_init()
	metricValues := getPodCount(clientSet)

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	//此处应该有问题，没有对协程做判断，需要后续优化
	for count := range metricValues {
		ch <- prometheus.MustNewConstMetric(collector.countMetric, prometheus.CounterValue, float64(count))
	}


}


