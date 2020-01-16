package main

import (
	"fmt"

	"github.com/gammazero/workerpool"
)

var Version string

func main() {
	fmt.Println("Start")
	wp := workerpool.New(2)
	wp.Submit(func() {
		fmt.Println("Pooling")
	})
	wp.StopWait()
}
