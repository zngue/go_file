package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_file/app/api/upload"
	"github.com/zngue/go_helper/pkg/common_run"
)

//group := engine.Group("file")
//		engine.Static("images", "images")
//		group.POST("upload", upload.Upload)
//		group.POST("uploadMax", upload.MaxUpload)
func main() {
	common_run.CommonGinRun(
		common_run.FnRouter(func(engine *gin.Engine) {
			group := engine.Group("file")
			engine.Static("images", "images")
			group.POST("upload", upload.Upload)
			group.POST("uploadMax", upload.MaxUpload)
		}),
		common_run.IsRegisterCenter(true),
	)
}
