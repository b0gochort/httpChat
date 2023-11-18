package model

type NewChatReq struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type NewChatItem struct {
	ID            int64       `json:"id" reindex:"id,hash,pk"`
	Time          int64       `json:"time" reindex:"time,tree"`
	Message       MessageType `json:"message"`
	LastHostStaff bool        `json:"last_host_staff" reindex:"last_host_staff,-"`
	UID           int64       `json:"uid" reindex:"uid,hash"`
	IP            string      `json:"ip" reindex:"ip,hash"`
	Category      string      `json:"category" reindex:"category,hash"`
}

type MessageType struct {
	NumberOfUnread int    `json:"number_of_unread" reindex:"number_of_unread,-"`
	LastMessage    string `json:"last_message" reindex:"last_message,-"`
}

type NewChatRes struct {
	ID            int64       `json:"id"`
	Time          int64       `json:"time"`
	Message       MessageType `json:"message"`
	LastHostStaff bool        `json:"last_host_staff"`
	UID           int64       `json:"uid"`
	IP            string      `json:"ip"`
	Category      string      `json:"category"`
}

type NewMessageReq struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MessageItem struct {
	ID     int64    `json:"id" reindex:"id,hash,pk"`
	ChatId int64    `json:"chat_id" reindex:"chat_id,hash"`
	Time   int64    `json:"time" reindex:"time,tree"`
	Host   HostType `json:"host"`
	Text   string   `json:"text" reindex:"text,-"`
}

type HostType struct {
	UserId int64  `json:"user_id" reindex:"user_id,hash"`
	Sub    string `json:"sub" reindex:"sub,hash"`
}

type NewMessageRes struct {
	ID int64 `json:"id" `
}

type Message struct {
	ID     int64    `json:"id"`
	ChatId int64    `json:"chat_id"`
	Time   int64    `json:"time"`
	Host   HostType `json:"host"`
	Text   string   `json:"text"`
}

type GetRoomsReq struct {
	Token string `json:"token"`
}
