package main

import "fmt"

type Printer interface {
	podprint()
}

type Pod1 struct {
	user string
	class string
}

type Pod2 struct {
	user string
	class string
	level string
}
func (a *Pod1) podprint()  {
	fmt.Printf("user：%s - class: %s\n", a.user, a.class)
}

func (a *Pod2) podprint() {
	fmt.Printf("user：%s - class: %s,his level is %s\n", a.user, a.class, a.level)
}
func PrinterExec(n Printer){
	n.podprint()
}
func main()  {
	lilis := Pod1{
		"lii",
		"Pod1",
	}
	PrinterExec(&lilis)
	solo := Pod2{
		"Dell Walker",
		"Engineer",
		"Captain",
	}
	PrinterExec(&solo)
}
