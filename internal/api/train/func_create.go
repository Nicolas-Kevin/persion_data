package train

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/train"
	"net/http"
)

type createRequest struct {
	TrainName    string `form:"trainName"`    // 演训名称
	TrainData    string `form:"trainData"`    // 时间年月日
	TrainOrg     string `form:"trainOrg"`     // 地点（市）
	TrainNum     int32  `form:"trainNum"`     // 参赛人员数量
	TrainProject string `form:"trainProject"` // 科目
	TrainPlace   string `form:"trainPlace"`   // 组织方
	TrainTime    string `form:"trainTime"`    // 宣传图片，视频
}

type createResponse struct {
	Id  int32  `json:"id"`
	Msg string `json:"msg"`
}

// Create 创建
// @Summary 创建
// @Description 创建
//@Tags API.train
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
//@Router /api/train/create [get]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		rep := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		createdata := new(train.CreateTrainData)
		createdata.TrainNum = req.TrainNum
		createdata.TrainData = req.TrainData
		createdata.TrainName = req.TrainName

		createdata.TrainOrg = req.TrainOrg
		createdata.TrainTime = req.TrainTime
		createdata.TrainPlace = req.TrainPlace

		create, err := h.trainService.Create(ctx, createdata)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedCreateError,
				code.Text(code.AuthorizedCreateError)).WithError(err),
			)
			return
		}

		rep.Id = create
		rep.Msg = "添加成功"
		ctx.Payload(rep)

	}
}
