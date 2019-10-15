package ops

import (
	"fmt"
	"github.com/hanlimo/check/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	meta1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type Pods struct { //下一步重构需要的结构体 podsRaw->pods
	PodsRaw
	restartCount	int32
	nameSpace 		string
}

type PodsRaw struct {
	podList			*v1.PodList
	restartSum 		[]int32
	err 			error
}

func Run(clientSet kubernetes.Clientset)  {
	//下一步重构需重命名的变量名 podsRaw->pods
	podsRaw := PodsRaw{}
	podsRaw.podList, podsRaw.err = clientSet.CoreV1().Pods("").List(meta1.ListOptions{})
	if podsRaw.err != nil {
		print("Pod解包失败")
	}
	//restartCountSum为获取到的所有集群Pod重启次数，默认是第一个容器
	restartCountSum := podsRaw.RestartCountSum(podsRaw)

	podItem := podsRaw.podList.Items
	for i:=0; i<len(podItem); i++ {
		if restartCountSum[i] > 5 {
			fmt.Printf("PodsRaw:%s中%s容器重启%d次，请相关人员排查\n",
				podItem[i].Name, podItem[i].Status.ContainerStatuses, restartCountSum[i])
		} else {
			fmt.Printf("容器%25s正常\n",podItem[i].Name)
		}
		exporterOut(restartCountSum[i])
	}
	return
}

func (*PodsRaw) RestartCountSum(pods PodsRaw) (restartSum []int32) {

	if pods.err !=nil {
		fmt.Printf("Restart_Count program fsailed.")
	}

	for _, pod := range pods.podList.Items {
		count := pod.Status.ContainerStatuses[0].RestartCount
		pods.restartSum = append(restartSum, count)
		//一个pod有多个容器时使用该循环
		//m := len(pod.Status.ContainerStatuses)
		//for i := 0 ;i<m; i++ {count := pod.Status.ContainerStatuses[i].RestartCount}
	}
	return pods.restartSum
}

func exporterOut(count int32) {

	//创建controller实例并在prometheus客户端注册
	foo := exporter.NewCheckCollector(count)
	prometheus.MustRegister(foo)

	//在 /metrics 下暴露端口
	http.Handle("/metrics", promhttp.Handler())
	log.Info("即将运行到 8848 端口")
	log.Fatal(http.ListenAndServe(":8848", nil))
}
