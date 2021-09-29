package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zngue/go_file/app/api/upload"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/sign_chan"
)

func main() {
	if cerrs := pkg.NewConfig(); cerrs != nil {
		logrus.Fatal(cerrs)
		return
	}
	mysql, err2 := pkg.NewMysql()
	fmt.Println(err2)
	if mysql != nil {

	}
	port := viper.GetString("AppPort")
	run, errs := pkg.GinRun(port, func(engine *gin.Engine) {
		group := engine.Group("file")
		engine.Static("images", "images")
		group.POST("upload", upload.Upload)
		group.POST("uploadMax", upload.MaxUpload)
	})
	if errs != nil {
		sign_chan.SignLog(errs)
	}
	var err error
	go func() {
		err = run.ListenAndServe()
		if err != nil {
			sign_chan.SignLog(err)
		}
	}()
	sign_chan.SignChalNotify()
	sign_chan.ListClose(func(ctx context.Context) error {
		return run.Shutdown(ctx)
	})
}
