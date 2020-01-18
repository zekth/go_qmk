package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/zekth/go_qmk/api/src/controllers"
)

var Version string

func foo() {
	fmt.Println("Bar")
}

func workerpoolInjector(w *workerpool.WorkerPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("worker", w)
		c.Next()
	}
}

type EnvVars struct {
	WorkerNumber int `default:2`
}

func main() {
	fmt.Println("Start:" + Version)
	var env EnvVars
	if err := envconfig.Process("go_qmk", &env); err != nil {
		fmt.Println("Unable to get env vars")
	}

	wp := workerpool.New(env.WorkerNumber)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(workerpoolInjector(wp))
	r.GET("/ping", controllers.Ping)

	r.Static("/ui", "./ui")
	if err := r.Run(); err != nil {
		fmt.Println("Fatal error")
	}
}
