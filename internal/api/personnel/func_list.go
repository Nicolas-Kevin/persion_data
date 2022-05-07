package personnel

import (
	"net/http"

	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql/s_personnel"
	"go-gin-api/internal/services/personnel"
)

type listData struct {
	Id               int64  //
	UserName         string // 姓名
	UserAge          int32  // 年龄
	UserSex          int32  // 性别（0：女，1：男）
	UserDate         string // 出生年月日
	PoliticalOutLook string // 社会面貌：0 :普通公民，1：团员，2：党员
	Company          string // 工作单位
	Position         string // 职称/等级
	Major            string // 专业
	UserLevel        string // 级别
	UserHeight       int32  // 身高cm
	UserWeight       int32  // 体重cm
	ShoeSize         int32  // 鞋号cm
	HeadNum          int32  // 头围cm
	BloodType        int32  // 血型cm
	IdNum            string // 身份证号
	BankNum          string // 银行卡号
	PhoneNum         string // 电话号
	UserTeam         int64  // 部队编号
	TrainStatus      string // 演训情况
	RewardStatus     string // 奖励情况
	Deleted          int32  //
}

type listRequest struct{}

type listResponse struct {
	Code  int                      `json:"code"`
	Count int                      `json:"count"`
	List  []s_personnel.SPersonnel `json:"data"`
}

// List 人员列表
// @Summary 人员列表
// @Description 人员列表
// @Tags API.personnel
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/personnel/list [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		//req := new(listRequest)
		res := new(listResponse)

		resListData, err := h.personnelServcie.List(ctx, new(personnel.SearchData))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res.List = make([]s_personnel.SPersonnel, len(resListData))
		for k, v := range resListData {
			if err != nil {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := s_personnel.SPersonnel{
				Id:               v.Id,
				UserName:         v.UserName,
				UserAge:          v.UserAge,
				UserSex:          v.UserSex,
				UserDate:         v.UserDate,
				PoliticalOutLook: v.PoliticalOutLook,
				Company:          v.Company,
				Position:         v.Position,
				Major:            v.Major,
				UserLevel:        v.UserLevel,
				UserHeight:       v.UserHeight,
				UserWeight:       v.UserWeight,
				ShoeSize:         v.ShoeSize,
				HeadNum:          v.HeadNum,
				BloodType:        v.BloodType,
				IdNum:            v.IdNum,
				BankNum:          v.BankNum,
				PhoneNum:         v.PhoneNum,
				UserTeam:         v.UserTeam,
				TrainStatus:      v.TrainStatus,
				RewardStatus:     v.RewardStatus,
				Deleted:          v.Deleted,
			}
			res.Code = 0
			res.Count = len(resListData)
			res.List[k] = data
		}
		ctx.Payload(res)
	}
}
