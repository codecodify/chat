package dao

import (
	"context"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOneMessage 保存消息
func InsertOneMessage(message *model.Message) error {
	message.CreatedAt = helper.GetCurrentTime()
	message.UpdatedAt = helper.GetCurrentTime()
	message.Identity = helper.GetUUID()
	_, err := model.Mongo.Collection(message.CollectionName()).InsertOne(context.Background(), message)
	return err
}

func GetMessageListByRoomIdentity(roomIdentity string, limit, skip *int64) ([]*model.Message, error) {
	cursor, err := model.Mongo.Collection(model.Message{}.CollectionName()).Find(
		context.Background(),
		bson.M{"room_identity": roomIdentity},
		&options.FindOptions{
			Limit: limit,
			Skip:  skip,
			Sort:  bson.M{"created_at": -1},
		},
	)
	if err != nil {
		return nil, err
	}
	var messageList []*model.Message
	err = cursor.All(context.Background(), &messageList)
	return messageList, err
}
