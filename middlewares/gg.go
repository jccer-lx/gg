package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/params"
	"github.com/lvxin0315/gg/validation"
	"github.com/sirupsen/logrus"
	"net/http"
)

//把接受参数放到中间件，这样应该能减少api代码量
func GGMiddleware(c *gin.Context) {
	//路由第二位必须是api
	if !isApi(c) {
		return
	}
	err := ggBeforeMiddleware(c)
	if err != nil {
		logrus.Error("ggBeforeMiddleware error", err)
		SetGGError(c, err)
	}

	if err == nil {
		//执行业务api
		c.Next()
	}
	//执行后通过上下文做出api的返回值
	ggAfterMiddleware(c)
}

func ggBeforeMiddleware(c *gin.Context) error {
	logrus.Info("ggBeforeMiddleware")
	c.Keys["output"] = new(helper.Output)
	c.Keys["error"] = nil
	c.Keys["params"] = nil
	//参数处理
	p := params.NewRouterParamInterface(c.HandlerName())
	if p != nil {
		err := GGBindParams(c, p)
		if err != nil {
			return err
		}
		c.Keys["params"] = p
	}

	//session
	//TODO
	return nil
}

func ggAfterMiddleware(c *gin.Context) {
	logrus.Info("ggAfterMiddleware")
	op := c.Keys["output"].(*helper.Output)
	if c.Keys["error"] != nil {
		op.Err = c.Keys["error"].(error)
	}
	op.ReturnOutput()
	c.JSON(http.StatusOK, op)
	c.Abort()
}

func isApi(c *gin.Context) bool {
	uriList := helper.GetUriStringList(c)
	if len(uriList) < 3 {
		return false
	}
	return uriList[2] == "api"
}

//通用接受参数方法
func GGBindParams(c *gin.Context, data params.GGParams) error {
	if data == nil {
		return nil
	}
	err := c.ShouldBind(data)
	if err != nil {
		logrus.Error("middle ShouldBind:", err)
		return err
	}
	//validate
	err = validation.Check(data)
	if err != nil {
		logrus.Error("middle Validation:", err)
		return err
	}
	return nil
}

//通用的设置返回error
func SetGGError(c *gin.Context, err error) {
	c.Keys["error"] = err
}
