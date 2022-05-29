package personnel

import (
	"github.com/go-playground/validator/v10"
	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/services/personnel"
	"net/http"
	"regexp"
)

type createRequest struct {
	UserName         string `form:"UserName" validator:"gt=0"`                // 姓名
	UserAge          int32  `form:"UserAge" validator:"gt=0"`                 // 年龄
	UserSex          int32  `form:"UserSex" validator:"oneof=0 1"`            // 性别（0：女，1：男）
	UserDate         string `form:"UserDate" validator:"gt=0"`                // 出生年月日
	PoliticalOutLook int32  `form:"PoliticalOutLook" validator:"oneof=0 1 2"` // 社会面貌：0 :普通公民，1：团员，2：党员
	Company          string `form:"Company" validator:"gt=0"`                 // 工作单位
	Position         string `form:"Position" `                                // 职称/等级
	Major            string `form:"Major" `                                   // 专业
	UserLevel        string `form:"UserLevel" validator:"gt=0"`               // 级别
	UserHeight       int32  `form:"UserHeight" validator:"gt=0"`              // 身高cm
	UserWeight       int32  `form:"UserWeight" validator:"gt=0"`              // 体重cm
	ShoeSize         int32  `form:"ShoeSize" validator:"gt=0"`                // 鞋号cm
	HeadNum          int32  `form:"HeadNum" validator:"gt=0"`                 // 头围cm
	BloodType        string `form:"BloodType" validator:"gt=0"`               // 血型cm
	IdNum            string `form:"IdNum" validator:"gt=0,len=18"`            // 身份证号
	BankNum          string `form:"BankNum" validator:"gt=0"`                 // 银行卡号
	PhoneNum         string `form:"PhoneNum" validator:"gt=0,len=11"`         // 电话号
	UserTeam         int64  `form:"UserTeam"`                                 // 部队编号
	TrainStatus      string `form:"TrainStatus" validator:"gt=0"`             // 演训情况
	RewardStatus     string `form:"RewardStatus" validator:"gt=0"`            // 奖励情况

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
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		validate := validator.New()
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		createData := new(personnel.CreatePersonnelData)
		validate.RegisterStructValidation(idcardValidatng, createData)
		validate.RegisterStructValidation(phoneNumValidatng, createData)
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

		//参数校验
		err2 := validate.Struct(createData)
		if err2 != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithError(err2),
			)
			return
		}
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

func phoneNumValidatng(sl validator.StructLevel) {
	personnelData := sl.Current().Interface().(*personnel.CreatePersonnelData)
	if len(personnelData.PhoneNum) == 11 {
		matchString, err := regexp.MatchString("^1[3456789]\\d{9}$", personnelData.IdNum)
		if err != nil || !matchString {
			sl.ReportError("PhoneNum", "PhoneNum", "PhoneNum", "手机号格式错误", "")
		}
	} else {
		sl.ReportError("PhoneNum", "PhoneNum", "PhoneNum", "手机号格式错误", "")
	}
}

//身份正校验
func idcardValidatng(sl validator.StructLevel) {
	personnelData := sl.Current().Interface().(*personnel.CreatePersonnelData)
	if len(personnelData.IdNum) == 18 {
		matchString, err := regexp.MatchString("^([1-6][1-9]|50)\\d{4}(18|19|20)\\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$", personnelData.IdNum)
		if err != nil || !matchString {
			sl.ReportError("IdNum", "IdNum", "IdNum", "身份证格式错误", "")
		}
	} else {
		sl.ReportError("IdNum", "IdNum", "IdNum", "身份证格式错误", "")
	}

}
