package main

import (
	"fmt"

	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/zekth/go_qmk/api/internal/environment"
	"github.com/zekth/go_qmk/api/internal/middlewares"
	"github.com/zekth/go_qmk/api/internal/routes"
)

// Version of the application. Githash
var Version string

func foo() {
	fmt.Println("Bar")
}

func main() {
	fmt.Println("Start:" + Version)
	var env environment.EnvVars
	env.Version = Version
	if err := envconfig.Process("go_qmk", &env); err != nil {
		fmt.Println("Unable to get env vars")
	}

	wp := workerpool.New(env.WorkerNumber)

	r := gin.New()

	// Middlewares usage
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.WorkerpoolInjector(wp))
	r.Use(middlewares.EnvInjector(env))

	// routing
	routes.MakeRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Println("Fatal error")
	}
}
