package main

import (
	"fmt"
	"github.com/hanlimo/check/kubeconf"
	"github.com/hanlimo/check/ops"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Pod struct {
	clientSet kubernetes.Clientset
	containerStatus v1.ContainerStatus
	v1.PodList
}

func main() {
	clientSet, err := kubeconf.Kubeconfig_init()
	if err != nil {
		fmt.Printf("error")
	}
	fmt.Printf("集群pod状态总览：\n")
	ops.PodPrint(clientSet)
	fmt.Printf("当前集群不正常pod列表：\n")
	ops.RestartCount(clientSet)
	fmt.Printf("磁盘统计状况：\n")
	ops.DirDetect()
}

