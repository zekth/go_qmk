package dependencies

import (
	"github.com/gammazero/workerpool"
	"github.com/patrickmn/go-cache"
	"github.com/zekth/go_qmk/api/internal/environment"
)

type Dependencies struct {
	Env     *environment.EnvVars
	Wp      *workerpool.WorkerPool
	Storage *cache.Cache
}
