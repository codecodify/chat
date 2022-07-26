package model

type User struct {
	Identity  string `bson:"identity" json:"identity"`
	Account   string `bson:"account" json:"account" binding:"required"`   // 账号
	Password  string `bson:"password" json:"password" binding:"required"` // 密码
	Nickname  string `bson:"nickname" json:"nickname" binding:"required"` // 昵称
	Sex       int    `bson:"sex" json:"sex" binding:"oneof=0 1 2"`        // 0-未知 1-男 2-女
	Email     string `bson:"email" json:"email" binding:"required"`       // 邮箱
	Avatar    string `bson:"avatar" json:"avatar" binding:"required"`     // 头像
	CreatedAt int64  `bson:"created_at" json:"created_at"`                // 创建时间
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`                // 操作时间
}

func (u User) CollectionName() string {
	return "users"
}
