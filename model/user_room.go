package model

type UserRoom struct {
	Identity     string `bson:"identity"`
	UserIdentity string `bson:"user_identity"` // 用户唯一标识
	RoomIdentity string `bson:"room_identity"` // 房间唯一标识
	RoomType     int    `bson:"room_type"`     // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
	UpdatedAt    int64  `bson:"updated_at"`    // 操作时间
}

func (r UserRoom) CollectionName() string {
	return "user_rooms"
}
