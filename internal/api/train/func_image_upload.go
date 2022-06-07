package train

import (
	"fmt"
	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	"net/http"
)

type imageUploadRequest struct {
}

type imageUploadResponse struct {
	ErrCode int    `json:"err_code"`
	Msg     string `json:"msg"`
}

// ImageUpload 文件上传
// @Summary 图片上传
// @Description 图片上传
//@Tags API.train
// @Accept application/multipart
// @Produce json
// @Param Request body imageUploadRequest true "请求信息"
// @Success 200 {object} imageUploadResponse
// @Failure 400 {object} code.Failure
//@Router /api/train/imageUpload [get]
func (h *handler) ImageUpload() core.HandlerFunc {

	return func(ctx core.Context) {
		req := new(imageUploadResponse)
		ctx.Request().ParseMultipartForm(ctx.Request().ContentLength)
		headers := ctx.Request().MultipartForm.File["file"]
		filePaht := ""
		for _, fheader := range headers {
			path, err := h.trainService.ImageUpload(fheader)
			if err != nil {
				fmt.Println(err)
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AuthorizedCreateError,
					code.Text(code.AuthorizedCreateError)).WithError(err),
				)

				return
			}
			filePaht += path + ","
		}
		req.ErrCode = http.StatusOK
		req.Msg = filePaht
		ctx.Payload(req)

	}
}
