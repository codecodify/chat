package service

import (
	"github.com/codecodify/chat/dao"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"github.com/codecodify/chat/vars"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	defer conn.Close()
	user := ctx.MustGet("Auth").(*model.User)
	wc[user.Identity] = conn
	for {
		ms := &vars.MessageStruct{}
		err := conn.ReadJSON(ms)
		if err != nil {
			helper.Error(ctx, err)
			return
		}

		// 判断用户是否在指定房间里
		if _, err := dao.GetUserRoomByUserIdentityRoomIdentity(user.Identity, ms.RoomIdentity); err != nil {
			helper.Error(ctx, err)
			return
		}

		go func() {
			// 保存消息
			message := &model.Message{
				UserIdentity: user.Identity,
				RoomIdentity: ms.RoomIdentity,
				Data:         ms.Message,
			}
			if err := dao.InsertOneMessage(message); err != nil {
				log.Printf("保存消息失败%s\n", err)
				helper.Error(ctx, err)
				return
			}
		}()

		// 获取聊天室在线用户
		userRooms, err := dao.GetUserRoomByRoomIdentity(ms.RoomIdentity)
		if err != nil {
			helper.Error(ctx, err)
			return
		}

		for _, room := range userRooms {
			if r, ok := wc[room.UserIdentity]; ok {
				if err := r.WriteJSON(ms); err != nil {
					helper.Error(ctx, err)
					return
				}
			}
		}

	}
}
