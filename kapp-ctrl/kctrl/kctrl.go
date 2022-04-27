package kctrl

import (
	"fmt"
	"k8s.io/client-go/rest"
)

func Stuff() {
	fmt.Println("yo")
	_ = rest.Config{}
}
