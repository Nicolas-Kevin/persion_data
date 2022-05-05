package train

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/s_train"
	"github.com/xinliangnote/go-gin-api/internal/services/train"
	"net/http"
)

type trainInfo struct {
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
type listRequest struct {
}

type listResponse struct {
	Code  int              `json:"code"`
	Count int              `json:"count"`
	List  []s_train.STrain `json:"data"`
}

// List 人员列表
// @Summary 人员列表
// @Description 人员列表
// @Tags API.train
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/tran/list [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		res := new(listResponse)
		resListData, err := h.trainService.List(ctx, new(train.SearchData))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res.List = make([]s_train.STrain, len(resListData))
		for k, v := range resListData {
			if err != nil {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := s_train.STrain{
				Id:           v.Id,
				TrainName:    v.TrainName,
				TrainData:    v.TrainData,
				TrainOrg:     v.TrainOrg,
				TrainNum:     v.TrainNum,
				TrainProject: v.TrainProject,
				TrainPlace:   v.TrainPlace,
				TrainTime:    v.TrainTime,
				Deleted:      v.Deleted,
			}
			res.Code = 0
			res.Count = len(resListData)
			res.List[k] = data
		}
		ctx.Payload(res)
	}
}
