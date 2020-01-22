package middlewares

import (
	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/zekth/go_qmk/api/internal/environment"
)

// WorkerpoolInjector Inject the Workerpool pointer in the gin context to be available
func WorkerpoolInjector(w *workerpool.WorkerPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("worker", w)
		c.Next()
	}
}

// EnvInjector Inject the value of the environment in the gin context
func EnvInjector(e environment.EnvVars) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("env", e)
		c.Next()
	}
}

// StorageInjector Inject the pointer of the storage provider
func StorageInjector(s *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("storage", s)
		c.Next()
	}
}
