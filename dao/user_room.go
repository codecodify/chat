package dao

import (
	"context"
	"fmt"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*model.UserRoom, error) {
	var userRoom model.UserRoom
	err := model.Mongo.Collection(userRoom.CollectionName()).FindOne(
		context.Background(),
		bson.M{"user_identity": userIdentity, "room_identity": roomIdentity},
	).Decode(&userRoom)
	return &userRoom, err
}

func GetUserRoomByRoomIdentity(roomIdentity string) ([]*model.UserRoom, error) {
	var userRooms []*model.UserRoom
	cursor, err := model.Mongo.Collection(userRooms[0].CollectionName()).Find(
		context.Background(),
		bson.M{"room_identity": roomIdentity},
	)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &userRooms)
	return userRooms, err
}

func JudgeUserIsFriend(userIdentity1, userIdentity2 string) bool {
	var roomIdentities = make([]string, 0)
	cursor, err := model.Mongo.Collection(model.UserRoom{}.CollectionName()).Find(
		context.Background(),
		bson.M{"user_identity": userIdentity1, "room_type": 1},
	)
	if err != nil {
		log.Printf("查询userIdentity1:%s错误:%s\n", userIdentity1, err)
		return false
	}

	for cursor.Next(context.Background()) {
		room := new(model.UserRoom)
		err = cursor.Decode(room)
		if err != nil {
			fmt.Printf("处理userIdentity1:%s游标错误:%s\n", userIdentity1, err)
			return false
		}
		roomIdentities = append(roomIdentities, room.RoomIdentity)
	}

	documents, err := model.Mongo.Collection(model.UserRoom{}.CollectionName()).CountDocuments(
		context.Background(),
		bson.M{
			"user_identity": userIdentity2,
			"room_type":     1,
			"room_identity": bson.M{
				"$in": roomIdentities,
			},
		},
	)

	if err != nil {
		log.Printf("查询userIdentity2:%s错误:%s\n", userIdentity2, err)
		return false
	}

	if documents > 0 {
		return true
	}
	return false
}

func GetUserRoomIdentity(userIdentity1, userIdentity2 string) string {
	roomIdentities := make([]string, 0)
	cursor, err := model.Mongo.Collection(model.UserRoom{}.CollectionName()).Find(
		context.Background(),
		bson.M{"user_identity": userIdentity1, "room_type": 1},
	)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return ""
	}
	var userRoom model.UserRoom
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&userRoom)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			return ""
		}
		roomIdentities = append(roomIdentities, userRoom.RoomIdentity)
	}

	err = model.Mongo.Collection(model.UserRoom{}.CollectionName()).FindOne(
		context.Background(),
		bson.M{"user_identity": userIdentity2, "room_type": 1, "room_identity": bson.M{"$in": roomIdentities}},
	).Decode(&userRoom)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return ""
	}
	return userRoom.Identity
}

func CreateUserRoom(userRoom *model.UserRoom) error {
	userRoom.CreatedAt = helper.GetCurrentTime()
	userRoom.UpdatedAt = helper.GetCurrentTime()
	userRoom.Identity = helper.GetUUID()
	_, err := model.Mongo.Collection(userRoom.CollectionName()).InsertOne(context.Background(), userRoom)
	return err
}

func DeleteUserRoom(roomIdentity string) error {
	_, err := model.Mongo.Collection(model.UserRoom{}.CollectionName()).DeleteOne(context.Background(), bson.M{"room_identity": roomIdentity})
	return err
}
