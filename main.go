package main

import (
	"check/kubeconf"
	"check/ops"
	"fmt"
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
	fmt.Printf("集群pod状态总览：")
	ops.PodPrint(clientSet)
	fmt.Printf("当前集群不正常pod列表：\n")
	ops.RestartCount(clientSet)
}

