package server

import (
	"github.com/gin-gonic/gin"
	"pixel-haven/interval/api"
)

func registerRoutes(router *gin.Engine) {
	group := router.RouterGroup.Group("/api/v1")

	// 在 "/api/v1" 下注册文件上传路由
	api.Upload(group)
}
