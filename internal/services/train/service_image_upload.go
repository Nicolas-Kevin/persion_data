package train

import (
	"io"
	"mime/multipart"
	"os"
)

func (s *service) ImageUpload(fheader *multipart.FileHeader) (string, error) {
	src, err := fheader.Open() //先把传过来的文件open，拿到文件的文件流
	if err != nil {
		return "", err
	}
	defer src.Close()
	dst := "upload/"
	out, err := os.Create(dst + fheader.Filename) //在本地创建一个文件
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src) //文件拷贝
	if err != nil {
		return "", err
	}
	return dst + fheader.Filename, err

}
