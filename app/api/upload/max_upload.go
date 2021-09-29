package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/shamsher31/goimgtype"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/code"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func FileContentType(out multipart.File) string {

	buf := make([]byte, 512)
	read, err := out.Read(buf)

	fmt.Println(read, err)

	contentType := http.DetectContentType(buf)

	return contentType
}

func MaxUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	ext := path.Ext(file.Filename)
	var block float64 = 0

	f, err := file.Open()

	contentType := FileContentType(f)
	fmt.Println(err)
	fmt.Println(contentType)
	fmt.Println(block)
	fileId := RangeFileName()
	errsr := code.CreateMutiDir("./images/" + fileId)
	if errsr != nil {
		api.Error(ctx, api.Msg(errsr.Error()))
		return
	}
	fmt.Println(fileId)
	for {
		saveFile, openErr := os.OpenFile("./images/"+fileId+"/"+cast.ToString(block+1)+ext, os.O_CREATE|os.O_RDWR, 777)
		fmt.Println(openErr)
		buf := make([]byte, 1024*500)
		n, errs := f.Read(buf)
		if errs != nil && err != io.EOF {
			fmt.Println(errs)
		}
		if n == 0 {
			break
		}
		_, errsFr := saveFile.Write(buf)
		if errsFr != nil {
			fmt.Println(errsFr)
		}
		block++
	}

}
