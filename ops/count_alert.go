package ops

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
		}
	}

}

func CountAlert(count int32, podName string, containerName string) {
	if count > 5 {
		fmt.Printf("Pod:%s中%s容器重启%d次，请相关人员排查\n",count, podName, containerName)
	} else {
		fmt.Printf("容器%25s正常\n",containerName)
	}
}
