package model

type Room struct {
	Identity     string `bson:"identity"`
	UserIdentity string `bson:"user_identity"` // 用户唯一标识
	Number       string `bson:"number"`        // 房间号
	Name         string `bson:"name"`          // 房间名称
	Info         string `bson:"info"`          // 房间简介
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 操作时间
}

func (r Room) CollectionName() string {
	return "rooms"
}
