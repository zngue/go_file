package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/xinliangnote/go-util/md5"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/code"
	"math/rand"
	path2 "path"
	"strings"
	"time"
)

// RangeFileName /*
func RangeFileName() string {
	one := RandString(100) + cast.ToString(RandomInt(10, 999999999)) + cast.ToString(time.Now().Unix())
	s := md5.MD5(one)
	return s
}

// RandomInt /*
func RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}

// RandString /*
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// Upload /*
func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		api.Error(ctx, api.Msg(err.Error()))
		return
	}
	ext := path2.Ext(file.Filename)
	all := strings.ReplaceAll(ext, ".", "")
	format := time.Now().Format("200601")
	prex := "images/"
	path := format + "/" + all + "/"

	errsr := code.CreateMutiDir(prex + path)
	if errsr != nil {
		api.Error(ctx, api.Msg(errsr.Error()))
		return
	}
	uppth := path + RangeFileName() + ext
	err = ctx.SaveUploadedFile(file, prex+uppth)
	if err != nil {
		api.Error(ctx, api.Msg(err.Error()))
		return
	}
	//imgUrlHost := viper.GetString("img_url")
	api.Success(ctx, api.Data("/"+prex+uppth))
	return
}
