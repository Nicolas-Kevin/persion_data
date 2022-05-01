package train

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

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
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
