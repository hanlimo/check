package kubeconf

//设计思路，通过初始化函数将该结构体初始化
type CheckBox struct {
	podName []string
	podStatus string
	containerName []string
	containerRestartCount int32
	podRestartCount int32
}
func (*CheckBox) CheckInit(config string)  {}

type Checker interface {
	CheckInit()
	CountAlert()
	ConfGet()
}
