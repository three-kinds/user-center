package gin_utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

var lMIME2Extension = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
}

func ReceiveUploadedImage(ctx *gin.Context, fileHeader *multipart.FileHeader, relativePathHeader string) (relativePath string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", se.ValidationError("can not open uploaded file")

	}
	defer func() {
		_ = file.Close()
	}()

	headerData := make([]byte, 512)
	_, err = file.Read(headerData)
	if err != nil {
		return "", se.ValidationError("can not read uploaded file")
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", se.ValidationError("can not seek uploaded file")
	}

	fileMIME := http.DetectContentType(headerData)
	extension, ok := lMIME2Extension[fileMIME]
	if !ok {
		return "", se.ValidationError("invalid image file")
	}

	relativePath = fmt.Sprintf("%s.%s", relativePathHeader, extension)
	absPath := filepath.Join(initializers.Config.MediaRoot, relativePath)
	err = ctx.SaveUploadedFile(fileHeader, absPath)
	if err != nil {
		return "", se.ValidationError("save uploaded file failed")
	}

	return relativePath, nil
}
