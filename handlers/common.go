package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/middlewares"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/services"
	"net/http"
	"time"
)

//通用404页面
func errorView(c *gin.Context) {
	c.HTML(http.StatusNotFound, "layout/404.html", nil)
}

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "layout/common.html", nil)
}

//api通用返回值
func apiReturn(c *gin.Context, op *helper.Output) {
	op.ReturnOutput()
	c.JSON(http.StatusOK, op)
}

//通用接受参数方法
func ggBindParams(c *gin.Context, data params.GGParams) error {
	return middlewares.GGBindParams(c, data)
}

//通用的view
func GGView(c *gin.Context) {
	module := c.Param("module")
	action := c.Param("action")
	if module == "" || action == "" {
		errorView(c)
		return
	}
	//所有参数
	data := map[string]interface{}{
		"ID": c.Param("id"),
	}
	c.HTML(http.StatusOK, fmt.Sprintf("%s/%s.tpl", module, action), data)
}

//通用的获取output
func ggOutput(c *gin.Context) *helper.Output {
	return c.Keys["output"].(*helper.Output)
}

//通用的获取params
func ggParams(c *gin.Context) params.GGParams {
	if c.Keys["params"] == nil {
		return nil
	}
	return c.Keys["params"].(params.GGParams)
}

//通用的设置返回error
func setGGError(c *gin.Context, err error) {
	middlewares.SetGGError(c, err)
}

//通过的获取token
func getGGToken(c *gin.Context) string {
	return c.Keys["token"].(string)
}

//通用的列表
func ggList(c *gin.Context, model interface{}, pagination *helper.Pagination, orderBy ...string) {
	output := ggOutput(c)
	//分页参数
	err := c.ShouldBind(pagination)
	if err != nil {
		setGGError(c, err)
		return
	}
	err = services.GetList(model, pagination, orderBy...)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = pagination.Data
	output.Count = pagination.Count
}

//通用图片上传
func UploadPicApi(c *gin.Context) {
	output := ggOutput(c)
	fileList, err := c.MultipartForm()
	if err != nil {
		setGGError(c, err)
		return
	}
	if fileList.File["pic_file"] == nil || len(fileList.File["pic_file"]) == 0 {
		setGGError(c, fmt.Errorf("pic不能为空"))
		return
	}
	saveDir := fmt.Sprintf("assets/uploads/%d", time.Now().UnixNano())
	err = helper.CreateDir(saveDir)
	if err != nil {
		setGGError(c, fmt.Errorf("pic不能为空"))
		return
	}
	//单图上传
	if len(fileList.File["pic_file"]) == 1 {
		savePath := fmt.Sprintf("%s/%s", saveDir, fileList.File["pic_file"][0].Filename)
		err = c.SaveUploadedFile(fileList.File["pic_file"][0], savePath)
		if err != nil {
			setGGError(c, err)
			return
		}
		output.Data = "/" + savePath
	} else {
		//多图
		var savePathList []string
		for _, picFile := range fileList.File["pic_file"] {
			savePath := fmt.Sprintf("%s/%s", saveDir, picFile.Filename)
			err = c.SaveUploadedFile(picFile, savePath)
			if err != nil {
				setGGError(c, err)
				return
			}
			savePathList = append(savePathList, "/"+savePath)
		}
		output.Data = savePathList
	}
}

//通用富文本上传图片
func LayEditUploadPicApi(c *gin.Context) {
	output := ggOutput(c)
	fileList, err := c.MultipartForm()
	if err != nil {
		setGGError(c, err)
		return
	}
	if fileList.File["file"] == nil || len(fileList.File["file"]) == 0 {
		setGGError(c, fmt.Errorf("pic不能为空"))
		return
	}
	saveDir := fmt.Sprintf("assets/uploads/%d", time.Now().UnixNano())
	err = helper.CreateDir(saveDir)
	if err != nil {
		setGGError(c, fmt.Errorf("pic不能为空"))
		return
	}
	savePath := fmt.Sprintf("%s/%s", saveDir, fileList.File["file"][0].Filename)
	err = c.SaveUploadedFile(fileList.File["file"][0], savePath)
	if err != nil {
		setGGError(c, err)
		return
	}
	output.Data = map[string]string{
		"src":   "/" + savePath,
		"title": fileList.File["file"][0].Filename,
	}
}
