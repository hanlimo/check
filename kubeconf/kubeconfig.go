package kubeconf

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//获取集群的kubeconfig文件
func KubeConfig_init() (kubernetes.Clientset, error ) {
	var kubeConfig *string
	kubeConfig = flag.String("kubeConfig", "/Users/hanlimo/Desktop/admin.conf", "absolute path to the kubeConfig file")
	flag.Parse()

	//在 kubeConfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientSet
	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	return *clientSet, nil
}

