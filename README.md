# 项目介绍
实现简易即时通讯，与好友聊天本质也是建立聊天室，在添加好友时就创建聊天室。

# 技术栈
* `websocket`
* `mongodb`
* `gin`
* `jwt`

# 数据字典
## 用户集合
```json
{
    "account":"账号",
    "password":"密码",
    "nickname":"昵称",
    "sex" : 1, // 0-未知 1-男 2-女
    "email": "邮箱",
    "avatar":"头像",
    "created_at": 1, // 创建时间
    "updated_at": 1, // 更新时间
}
```

## 消息集合
```json
{
    "user_identity": "用户的唯一标识",
    "room_identity": "房间的唯一标识",
    "data": "发送的数据",
    "created_at": 1, // 创建时间
    "updated_at": 1, // 更新时间
}
```

## 房间集合
```json
{
    "number":"房间号",
    "name":"房间名称",
    "info":"房间简介",
    "user_identity": "房间创建者的唯一标识",
    "created_at": 1, // 房间的创建时间
    "updated_at": 1, // 房间的更新时间
}
```

## 用户房间关联集合
```json
{
    "user_identity": "用户的唯一标识",
    "room_identity": "房间的唯一标识",
    "message_identity": "消息的唯一标识",
    "created_at": 1, // 创建时间
    "updated_at": 1, // 更新时间
}
```