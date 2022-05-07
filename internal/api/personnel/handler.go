package personnel

import (
	"go-gin-api/configs"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	"go-gin-api/internal/services/personnel"
	"go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// List 人员列表
	// @Tags API.personnel
	// @Router /api/personnel/list [get]
	List() core.HandlerFunc
	// Create 创建
	//@Tags API.personnel
	//@Router /api/personnel/create [get]
	Create() core.HandlerFunc

	// Exp 导出
	//@Tags API.personnel
	//@Router /api/personnel/exp [post]
	Exp() core.HandlerFunc

	// GetById 获取指定id用户信息
	//@Tags API.personnel
	//@Router /api/personnel/getbyid [post]
	GetById() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	hashids hash.Hash
	db      mysql.Repo

	personnelServcie personnel.Service
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger:           logger,
		db:               db,
		hashids:          hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		personnelServcie: personnel.New(db),
	}
}

func (h *handler) i() {}
