package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"scaffold/pkg/snowflake"
)

func UploadImgHandler(c *gin.Context) {
	f, err := c.FormFile("img") //name值与input 中的name 一致
	if err != nil {
		zap.L().Error("c.FormFile(\"img\") failed err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//将获取的文件保存到本地
	ext := filepath.Ext(f.Filename)
	uid := snowflake.GenID()
	fName := fmt.Sprintf("%d%s", uid, ext)
	dst := fmt.Sprintf("/app/statics/img/%s", fName)
	//dst := filepath.Join("app", "statics", "img", fName)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		zap.L().Error("c.SaveUploadedFile(f, dst) failed,err:", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"imgUrl": fName,
	})
}
