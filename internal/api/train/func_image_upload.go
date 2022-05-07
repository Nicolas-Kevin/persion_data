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
	errCode int `json:"err_code"`
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
		//file, he, err := ctx.Request().FormFile("file")
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//filename := he.Filename
		//open, err := he.Open()
		//var c []byte
		//at, err := open.ReadAt(c, he.Size)
		//if err != nil && err.Error() != "EOF" {
		//	fmt.Println(err.Error(), at)
		//	return
		//}
		//
		//f2, err := os.Open(filename + ".png")
		//f2.Write(c)
		//defer f2.Close()
		//f := []byte{}
		//fmt.Println(file.Read(f))
		// 解析出文件名，获取最后一个元素信息：A20_Apex.apk

		//body, _ := ioutil.ReadAll(ctx.Request().Body)
		//
		//fmt.Println("---body/--- \r\n " + string(body))
		//mm := make(map[string]interface{})
		//json.Unmarshal(body, &mm)
		//fmt.Println(mm)
		headers := ctx.Request().MultipartForm.File["file"]
		fmt.Println(headers)
		for _, fheader := range headers {
			err := h.trainService.ImageUpload(fheader)
			if err != nil {
				fmt.Println(err)
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AuthorizedCreateError,
					code.Text(code.AuthorizedCreateError)).WithError(err),
				)
				return
			}
		}
		req.errCode = 200
		ctx.Payload(nil)
	}
}
