package personnel

import (
	"fmt"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/personnel"
)

type createRequest struct {
	UserName         string `form:"UserName"`         // 姓名
	UserAge          int32  `form:"UserAge"`          // 年龄
	UserSex          int32  `form:"UserSex"`          // 性别（0：女，1：男）
	UserDate         string `form:"UserDate"`         // 出生年月日
	PoliticalOutLook int32  `form:"PoliticalOutLook"` // 社会面貌：0 :普通公民，1：团员，2：党员
	Company          string `form:"Company"`          // 工作单位
	Position         string `form:"Position"`         // 职称/等级
	Major            string `form:"Major"`            // 专业
	UserLevel        string `form:"UserLevel"`        // 级别
	UserHeight       int32  `form:"UserHeight"`       // 身高cm
	UserWeight       int32  `form:"UserWeight"`       // 体重cm
	ShoeSize         int32  `form:"ShoeSize"`         // 鞋号cm
	HeadNum          int32  `form:"HeadNum"`          // 头围cm
	BloodType        string `form:"BloodType"`        // 血型cm
	IdNum            string `form:"IdNum"`            // 身份证号
	BankNum          string `form:"BankNum"`          // 银行卡号
	PhoneNum         string `form:"PhoneNum"`         // 电话号
	UserTeam         int64  `form:"UserTeam"`         // 部队编号
	TrainStatus      string `form:"TrainStatus"`      // 演训情况
	RewardStatus     string `form:"RewardStatus"`     // 奖励情况

}

type createResponse struct {
	Id  int64  `json:"id"`
	Msg string `json:"msg"`
}

// Create 创建
// @Summary 创建
// @Description 创建
//@Tags API.personnel
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
//@Router /api/personnel/create [get]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	fmt.Println("add")
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		createData := new(personnel.CreatePersonnelData)
		createData.UserName = req.UserName
		createData.UserAge = req.UserAge
		createData.UserSex = req.UserSex
		createData.UserDate = req.UserDate
		createData.PoliticalOutLook = req.PoliticalOutLook
		createData.Company = req.Company
		createData.Position = req.Position
		createData.Major = req.Major
		createData.UserLevel = req.UserLevel
		createData.UserHeight = req.UserHeight
		createData.UserWeight = req.UserWeight
		createData.ShoeSize = req.ShoeSize
		createData.HeadNum = req.HeadNum
		createData.BloodType = req.BloodType
		createData.IdNum = req.IdNum
		createData.BankNum = req.BankNum
		createData.PhoneNum = req.PhoneNum
		createData.UserTeam = req.UserTeam
		createData.TrainStatus = req.TrainStatus
		createData.RewardStatus = req.RewardStatus
		id, err := h.personnelServcie.Create(ctx, createData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		res.Msg = "添加成功"
		ctx.Payload(res)
	}
}
