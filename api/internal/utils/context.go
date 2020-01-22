package utils

import (
	"context"

	"github.com/gammazero/workerpool"
	"github.com/patrickmn/go-cache"
	"github.com/zekth/go_qmk/api/internal/dependencies"
	"github.com/zekth/go_qmk/api/internal/environment"
)

type contextKey string

func (c contextKey) String() string {
	return "ctx_" + string(c)
}

var (
	contextKeyEnv        = contextKey("env")
	contextKeyWorkerPool = contextKey("wp")
	contextKeyWStorage   = contextKey("storage")
)

func InjectEnvInContext(c context.Context, d dependencies.Dependencies) context.Context {
	c = context.WithValue(c, contextKeyEnv, d.Env)
	c = context.WithValue(c, contextKeyWorkerPool, d.Wp)
	c = context.WithValue(c, contextKeyWorkerPool, d.Storage)
	return c
}

func EnvFromContext(c context.Context) *environment.EnvVars {
	e := c.Value(contextKeyEnv).(*environment.EnvVars)
	return e
}

func WorkerFromContext(c context.Context) *workerpool.WorkerPool {
	wp := c.Value(contextKeyWorkerPool).(*workerpool.WorkerPool)
	return wp
}

func StorageFromContext(c context.Context) *cache.Cache {
	s := c.Value(contextKeyWStorage).(*cache.Cache)
	return s
}
