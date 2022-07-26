package dao

import (
	"context"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRoom(room *model.Room) error {
	room.CreatedAt = helper.GetCurrentTime()
	room.UpdatedAt = helper.GetCurrentTime()
	room.Identity = helper.GetUUID()
	_, err := model.Mongo.Collection(room.CollectionName()).InsertOne(context.Background(), room)
	return err
}

func DeleteRoom(roomIdentity string) error {
	_, err := model.Mongo.Collection(model.Room{}.CollectionName()).DeleteOne(context.Background(), bson.M{"identity": roomIdentity})
	return err
}
