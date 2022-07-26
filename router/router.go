package router

import (
	"github.com/codecodify/chat/middleware"
	"github.com/codecodify/chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户注册
	r.POST("/register", service.Register)
	// 用户登陆
	r.POST("/login", service.Login)

	// 需要登陆
	auth := r.Group("/").Use(middleware.AuthCheck())
	{
		// 用户信息
		auth.GET("user/info", service.UserInfo)
		// 添加用户
		auth.POST("user/add", service.AddUser)
		// 删除用户
		auth.DELETE("user/delete", service.DeleteUser)
		// 聊天室
		auth.GET("websocket/message", service.WebsocketMessage)
		// 聊天记录
		auth.GET("chat/list", service.ChatList)
	}

	return r
}
