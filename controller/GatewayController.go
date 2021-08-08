package controller

import (
	"log"
	"net/http"
	"net/url"

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
	// first solution
	c.SetCookie("wage", "123", 10, "/", c.Request.URL.Hostname(), false, true)
	c.SetCookie("amount", "13123", 10, "/", c.Request.URL.Hostname(), false, true)
	location1 := url.URL{Path: "http://127.0.0.1:22520/"}
	c.Redirect(http.StatusFound, location1.RequestURI())

	// second solution
	q := url.Values{}
	q.Set("wage", "123")
	q.Set("amount", "13123")
	location2 := url.URL{Path: "http://127.0.0.1:22520/api/callback/query_params", RawQuery: q.Encode()}
	c.Redirect(http.StatusFound, location2.RequestURI())
}
