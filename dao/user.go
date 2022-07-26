package dao

import (
	"context"
	"github.com/codecodify/chat/helper"
	"github.com/codecodify/chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CreateUser(user *model.User) error {
	user.Identity = helper.GetUUID()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.Password = helper.GetMd5(user.Password)
	_, err := model.Mongo.Collection(user.CollectionName()).InsertOne(context.Background(), user)
	return err
}

func GetUserByAccount(account string) (user *model.User, err error) {
	user = &model.User{}
	err = model.Mongo.Collection(user.CollectionName()).FindOne(context.Background(), bson.M{"account": account}).Decode(&user)
	return
}

func GetUserByIdentity(identity string) (user *model.User, err error) {
	user = &model.User{}
	err = model.Mongo.Collection(user.CollectionName()).FindOne(context.Background(), bson.M{"identity": identity}).Decode(&user)
	return
}
