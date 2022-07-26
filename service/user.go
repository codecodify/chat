package service

import (
	"errors"
	"fmt"
	"github.com/codecodify/chat/dao"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"github.com/gin-gonic/gin"
	"log"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		log.Printf("用户注册绑定参数失败: %s\n", err)
		helper.Error(ctx, err)
		return
	}
	if err := dao.CreateUser(&user); err != nil {
		log.Printf("添加用户失败: %s\n", err)
		helper.Error(ctx, err)
		return
	}
	helper.Success(ctx, nil)
}

// Login 用户登陆
func Login(ctx *gin.Context) {
	var req = struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Error(ctx, err)
		return
	}
	if len(req.Account) == 0 || len(req.Password) == 0 {
		helper.Error(ctx, errors.New("帐号或密码不能为空"))
		return
	}
	user, err := dao.GetUserByAccount(req.Account)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	if user.Password != helper.GetMd5(req.Password) {
		helper.Error(ctx, errors.New("用户名或密码错误"))
		return
	}

	if token, err := helper.GenerateToken(user.Identity, user.Email); err == nil {
		helper.Success(ctx, gin.H{"token": token})
	} else {
		helper.Error(ctx, err)
	}

}

// UserInfo 用户信息
func UserInfo(ctx *gin.Context) {
	user := ctx.MustGet("Auth").(*model.User)
	helper.Success(ctx, user)
}

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	var req = struct {
		Account string `json:"account"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Error(ctx, err)
		return
	}

	// 查找用户是否存在
	friend, err := dao.GetUserByAccount(req.Account)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	user := ctx.MustGet("Auth").(*model.User)

	// 检查好友是否已存在
	if dao.JudgeUserIsFriend(friend.Identity, user.Identity) {
		helper.Error(ctx, errors.New("互为好友，不可重复添加"))
		return
	}

	// 保存房间
	room := &model.Room{
		UserIdentity: user.Identity,
	}
	if err := dao.CreateRoom(room); err != nil {
		helper.Error(ctx, errors.New(fmt.Sprintf("数据库异常:%s\n", err)))
		return
	}

	// 保存用户与房间的关联
	userRoom := &model.UserRoom{
		UserIdentity: user.Identity,
		RoomIdentity: room.Identity,
		RoomType:     1,
	}
	if err := dao.CreateUserRoom(userRoom); err != nil {
		helper.Error(ctx, errors.New(fmt.Sprintf("数据库异常:%s\n", err)))
		return
	}

	userRoom = &model.UserRoom{
		UserIdentity: friend.Identity,
		RoomIdentity: room.Identity,
		RoomType:     1,
	}
	if err := dao.CreateUserRoom(userRoom); err != nil {
		helper.Error(ctx, errors.New(fmt.Sprintf("数据库异常:%s\n", err)))
		return
	}
	helper.Success(ctx, "添加成功")
}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	var req = struct {
		Identity string `json:"identity"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.Error(ctx, err)
		return
	}

	// 查找用户是否存在
	friend, err := dao.GetUserByIdentity(req.Identity)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	user := ctx.MustGet("Auth").(*model.User)

	// 检查好友是否已存在
	if !dao.JudgeUserIsFriend(friend.Identity, user.Identity) {
		helper.Error(ctx, errors.New("不是好友，不可删除"))
		return
	}

	// 获取房间Identity
	roomIdentity := dao.GetUserRoomIdentity(req.Identity, user.Identity)
	if len(roomIdentity) == 0 {
		helper.Error(ctx, errors.New("不是好友，不可删除"))
		return
	}

	// 删除用户与房间的关联
	if err := dao.DeleteUserRoom(roomIdentity); err != nil {
		helper.Error(ctx, errors.New(fmt.Sprintf("数据库异常:%s\n", err)))
		return
	}

	// 删除房间
	if err := dao.DeleteRoom(roomIdentity); err != nil {
		helper.Error(ctx, errors.New(fmt.Sprintf("数据库异常:%s\n", err)))
		return
	}
	helper.Success(ctx, "删除成功")
}
