package train

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/services/train"
	"github.com/xinliangnote/go-gin-api/pkg/hash"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// List 获取演训信息
	// @Tags API.train
	// @Router /api/train/list [get]
	List() core.HandlerFunc
	// Create 创建演训信息
	//@Tags API.train
	//@Router /api/train/create [get]
	Create() core.HandlerFunc

	// Exp 导出演训信息
	//@Tags API.train
	//@Router /api/train/exp [post]
	Exp() core.HandlerFunc

	// ImageUpload 图片上传
	//@Tags API.train
	//@Router /api/train/imageUpload [get]
	ImageUpload() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	hashids      hash.Hash
	db           mysql.Repo
	trainService train.Service
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger:       logger,
		db:           db,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		trainService: train.New(db),
	}
}

func (h *handler) i() {}
