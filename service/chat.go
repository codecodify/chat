package service

import (
	"errors"
	"github.com/codecodify/chat/dao"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ChatList(ctx *gin.Context) {
	roomIdentity := ctx.Query("room_identity")
	if len(roomIdentity) == 0 {
		helper.Error(ctx, errors.New("房间号不能为空"))
		return
	}
	user := ctx.MustGet("Auth").(*model.User)
	// 判断用户是否属于该房间
	if _, err := dao.GetUserRoomByUserIdentityRoomIdentity(roomIdentity, user.Identity); err != nil {
		helper.Error(ctx, errors.New("非法访问"))
		return
	}
	pageIndex, _ := strconv.ParseInt(ctx.Query("page_index"), 10, 32)
	pageSize, _ := strconv.ParseInt(ctx.Query("page_size"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	chatList, err := dao.GetMessageListByRoomIdentity(roomIdentity, &pageSize, &skip)
	if err != nil {
		helper.Error(ctx, errors.New("系统异常"))
		return
	}
	helper.Success(ctx, chatList)
}
