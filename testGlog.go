package main

import (
	"flag"
	"fmt"

	"my/tests/testGlog/glog"
)

func main() {
	fmt.Println("main start...")
	defer glog.Flush()
	flag.Parse()
	fmt.Println(*(glog.V(5)))
	glog.V(1).Info("glog info level 1")
	glog.V(1).Error("glog error level 1")
	glog.V(3).Error("glog error level 3")
	glog.V(0).Warning("glog warning level 1")
	glog.V(9).Info("glog info level 9")
}
