package ops

import (
	"fmt"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)


//通过实现 clientSet 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
//PodInterface 接口拥有操作 Pods 资源的方法，例如 Create、Update、Get、List 等方法
func PodPrint(clientSet kubernetes.Clientset) (error string) {
	// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pods 列表
	pods, err := clientSet.CoreV1().Pods("").List(meta.ListOptions{})
	if pods.Items == nil {
		fmt.Printf("集群状态获取失败\n")
		return "集群状态获取失败"
	} else {
		fmt.Printf("当前k8s集群总共有%d个pod。\n", len(pods.Items))
	}

	namespaces, _ := clientSet.CoreV1().Namespaces().List(meta.ListOptions{})
	// 获取指定 namespace 中的 Pods 列表信息
	//namespace := "kube-system"
	for _, namespace := range namespaces.Items{ //遍历namespaces
		//fmt.Printf("%s\n", namespace.Name)
		pods, err = clientSet.CoreV1().Pods(namespace.Name).List(meta.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, pod := range pods.Items {
			count := pod.Status.ContainerStatuses[0].RestartCount
			fmt.Printf(" %10s | %30s | 状态: %s | 重启次数: %d\n", namespace.Name, pod.ObjectMeta.Name, pod.Status.Phase,  count)

		}
	}
	return
	//time.Sleep(10 * time.Second)
}
