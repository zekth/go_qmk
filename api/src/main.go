package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "github.com/gammazero/workerpool"
  "github.com/zekth/go_qmk/src/controllers"
)

var Version string

func foo() {
  fmt.Println("Bar")
}

func main() {
	fmt.Println("Start:"+Version)
	wp := workerpool.New(2)
	wp.Submit(func() {
		fmt.Println("Pooling")
	})
  wp.StopWait()
  
	r := gin.New()

	r.Use(gin.Logger())
  r.Use(gin.Recovery())
  
	r.GET("/ping", controllers.Ping)

	r.Static("/ui", "./ui")
	r.Run()
}
