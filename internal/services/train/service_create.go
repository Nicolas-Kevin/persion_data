package train

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/s_train"
)

type CreateTrainData struct {
	TrainName    string // 演训名称
	TrainData    string // 时间年月日
	TrainOrg     string // 地点（市）
	TrainNum     int32  // 参赛人员数量
	TrainProject string // 科目
	TrainPlace   string // 组织方
	TrainTime    string // 宣传图片，视频
}

func (s *service) Create(ctx core.Context, trainData *CreateTrainData) (id int32, err error) {

	model := s_train.NewModel()
	model.TrainName = trainData.TrainName
	model.TrainData = trainData.TrainData
	model.TrainOrg = trainData.TrainOrg
	model.TrainNum = trainData.TrainNum
	model.TrainProject = trainData.TrainProject
	model.TrainPlace = trainData.TrainPlace
	model.TrainTime = trainData.TrainTime

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
