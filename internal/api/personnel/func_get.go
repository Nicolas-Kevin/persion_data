package personnel

import (
	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql/s_personnel"
	"go-gin-api/internal/services/personnel"
	"net/http"
)

type getRequest struct {
	Id int64 `form:"id"`
}

type getResponse struct {
	Code int                     `json:"code"`
	List *s_personnel.SPersonnel `json:"data"`
}

// GetById 人员列表
// @Summary 人员列表
// @Description 人员列表
// @Tags API.personnel
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/personnel/getById [get]
// @Security LoginToken
func (h *handler) GetById() core.HandlerFunc {
	return func(c core.Context) {
		req := new(getRequest)
		rep := new(getResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		searchData := new(personnel.SearchData)
		searchData.Id = req.Id
		list, err := h.personnelServcie.ListById(c, new(personnel.SearchData))
		if err != nil && len(list) > 0 {
			rep.Code = 0
			rep.List = list[0]
		}
		//c.HTML("personnel_create", list[0])
		c.Payload(rep)
	}
}
