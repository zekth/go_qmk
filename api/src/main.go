package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gammazero/workerpool"
)

var Version string

func foo() {
  fmt.Println("Bar")
}

func main() {
	fmt.Println("Start")
	wp := workerpool.New(2)
	wp.Submit(func() {
		fmt.Println("Pooling")
	})
	wp.StopWait()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": Version,
		})
	})

	r.Static("/ui", "./ui")
	r.Run()
}
