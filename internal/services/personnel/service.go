package personnel

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	"go-gin-api/internal/repository/mysql/s_personnel"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	List(ctx core.Context, searchData *SearchData) (listData []*s_personnel.SPersonnel, err error)

	ListById(ctx core.Context, searchData *SearchData) (listData []*s_personnel.SPersonnel, err error)

	Create(ctx core.Context, personneldData *CreatePersonnelData) (id int64, err error)
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
