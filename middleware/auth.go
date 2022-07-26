package middleware

import (
	"errors"
	"github.com/codecodify/chat/dao"
	"github.com/codecodify/chat/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从header获取token
		token := strings.TrimSpace(strings.ReplaceAll(ctx.GetHeader("Authorization"), "Bearer ", ""))
		if len(token) == 0 {
			helper.Error(ctx, errors.New("请登陆"))
			return
		}
		// 验证token
		userClaims, err := helper.ParseToken(token)
		if err != nil {
			helper.Error(ctx, err)
			return
		}
		if user, err := dao.GetUserByIdentity(userClaims.Identity); err == nil {
			ctx.Set("Auth", user)
		} else {
			helper.Error(ctx, err)
		}
	}
}
