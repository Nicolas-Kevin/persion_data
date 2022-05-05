package train

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/s_train"
	"mime/multipart"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	List(ctx core.Context, searchData *SearchData) (listData []*s_train.STrain, err error)

	Create(ctx core.Context, trainData *CreateTrainData) (id int32, err error)

	ImageUpload(fheader *multipart.FileHeader) (err error)
}

type service struct {
	db mysql.Repo
}

func New(db mysql.Repo) Service {
	return &service{
		db: db,
	}
}

func (s *service) i() {

}
