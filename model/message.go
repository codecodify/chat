package model

type Message struct {
	Identity     string `bson:"identity"`
	UserIdentity string `bson:"user_identity"` // 用户唯一标识
	RoomIdentity string `bson:"room_identity"` // 房间唯一表示
	Data         string `bson:"data"`          // 发送的数据
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 操作时间
}

func (m Message) CollectionName() string {
	return "messages"
}
