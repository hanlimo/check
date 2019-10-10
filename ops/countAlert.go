package ops

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)


func RestartCount(clientSet kubernetes.Clientset) {

	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err !=nil {
		fmt.Printf("Restart_Count program fsailed.")
	}
	for _, pod := range pods.Items {
		m := len(pod.Status.ContainerStatuses)
		for i := 0 ;i<m; i++ {
			count := pod.Status.ContainerStatuses[i].RestartCount
			podName := pod.Name
			containerName := pod.Status.ContainerStatuses[i].Name
			CountAlert(count, podName, containerName)
			//exporterOut(count)
		}
	}

}

func CountAlert(count int32, podName string, containerName string) {
	if count > 5 {
		fmt.Printf("Pod:%s中%s容器重启%d次，请相关人员排查\n", podName, containerName, count)
	} else {
		fmt.Printf("容器%25s正常\n",containerName)
	}
}
func exporterOut(count int32) {

	//Create a new instance of the countCollector and
	//register it with the prometheus client.
	foo := newCheckCollector(count)
	prometheus.MustRegister(foo)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("即将运行到 8848 端口")
	log.Fatal(http.ListenAndServe(":8848", nil))
}
