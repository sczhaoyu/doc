package main

import (
	ctr "github.com/sczhaoyu/doc/controllers"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ctr.StartHttp()

}
