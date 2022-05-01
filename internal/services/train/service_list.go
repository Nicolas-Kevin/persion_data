package train

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/s_train"
)

type SearchData struct {
	Id           int32  //
	TrainName    string // 演训名称
	TrainData    string // 时间年月日
	TrainOrg     string // 地点（市）
	TrainNum     int32  // 参赛人员数量
	TrainProject string // 科目
	TrainPlace   string // 组织方
	TrainTime    string // 宣传图片，视频
	Deleted      int32  //
}

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*s_train.STrain, err error) {

	qb := s_train.NewQueryBuilder()

	if searchData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchData.Id)
	}

	listData, err = qb.WhereId(mysql.EqualPredicate,searchData.Id).
		OrderById(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
