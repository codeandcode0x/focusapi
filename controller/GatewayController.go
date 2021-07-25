package controller

import (
	"focusapi/util"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// controller struct
type GatewayController struct {
	apiVersion string
}

/**
1. 使用 split 分割 route, 进行 model 初始化
2. 使用 gateway proxy 统一处理
*/

// time out mid
func (uc *GatewayController) Gateway(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var gatewayController *GatewayController
		log.Println("gateway handler ", c.Request.Header)
		router.GET("/movie/add", gatewayController.Proxy)
		router.GET("/movie/update", gatewayController.Proxy)
	}
}

func (uc *GatewayController) Proxy(c *gin.Context) {
	urls := strings.Split(c.Request.URL.Path, "/")
	log.Println("url .....", urls)
	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
	})
}
