package train

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type expRequest struct{}

type expResponse struct{}

// Exp 导出
// @Summary 导出
// @Description 导出
//@Tags API.personnel
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body expRequest true "请求信息"
// @Success 200 {object} expResponse
// @Failure 400 {object} code.Failure
//@Router /api/personnel/exp [post]
func (h *handler) Exp() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
