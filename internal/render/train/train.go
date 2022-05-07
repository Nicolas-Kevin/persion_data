package train

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	"go-gin-api/internal/repository/redis"
	"go.uber.org/zap"
)

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger: logger,
		//cache:  cache,
		db: db,
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("train_list", nil)
	}
}

func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("train_create", nil)
	}
}
