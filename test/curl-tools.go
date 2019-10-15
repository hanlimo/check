package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func init()  {
	if len(os.Args) != 2{
		fmt.Printf("缺少参数")
		os.Exit(-1)
	}
}

func main(){
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, " World!")
	io.Copy(os.Stdout, &b)
}
