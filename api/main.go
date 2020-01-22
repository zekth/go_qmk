package main

import (
	"fmt"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/patrickmn/go-cache"
	"github.com/zekth/go_qmk/api/internal/dependencies"
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
	c := cache.New(5*time.Minute, 10*time.Minute)
	r := gin.New()

	d := dependencies.Dependencies{
		Env:     &env,
		Wp:      wp,
		Storage: c,
	}

	// Middlewares usage
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.WorkerpoolInjector(wp))
	r.Use(middlewares.EnvInjector(env))
	r.Use(middlewares.StorageInjector(c))

	// routing
	routes.MakeRoutes(r, d)

	if err := r.Run(); err != nil {
		fmt.Println("Fatal error")
	}
}
