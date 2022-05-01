package personnel

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/s_personnel"
)

type CreatePersonnelData struct {
	UserName         string // 姓名
	UserAge          int32  // 年龄
	UserSex          int32  // 性别（0：女，1：男）
	UserDate         string // 出生年月日
	PoliticalOutLook int32  // 社会面貌：0 :普通公民，1：团员，2：党员
	Company          string // 工作单位
	Position         string // 职称/等级
	Major            string // 专业
	UserLevel        string // 级别
	UserHeight       int32  // 身高cm
	UserWeight       int32  // 体重cm
	ShoeSize         int32  // 鞋号cm
	HeadNum          int32  // 头围cm
	BloodType        string // 血型cm
	IdNum            string // 身份证号
	BankNum          string // 银行卡号
	PhoneNum         string // 电话号
	UserTeam         int64  // 部队编号
	TrainStatus      string // 演训情况
	RewardStatus     string // 奖励情况
}

func (s *service) Create(ctx core.Context, personneldData *CreatePersonnelData) (id int64, err error) {

	model := s_personnel.NewModel()
	model.UserName = personneldData.UserName
	model.UserAge = personneldData.UserAge
	model.UserSex = personneldData.UserSex
	model.UserDate = personneldData.UserDate
	model.PoliticalOutLook = personneldData.PoliticalOutLook
	model.Company = personneldData.Company
	model.Position = personneldData.Position
	model.Major = personneldData.Major
	model.UserLevel = personneldData.UserLevel
	model.UserHeight = personneldData.UserHeight
	model.UserWeight = personneldData.UserWeight
	model.ShoeSize = personneldData.ShoeSize
	model.HeadNum = personneldData.HeadNum
	model.BloodType = personneldData.BloodType
	model.IdNum = personneldData.IdNum
	model.BankNum = personneldData.BankNum
	model.PhoneNum = personneldData.PhoneNum
	model.UserTeam = personneldData.UserTeam
	model.TrainStatus = personneldData.TrainStatus
	model.RewardStatus = personneldData.RewardStatus

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
