package model

import (
	"github.com/golang-jwt/jwt/v5"
)

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
	RequestTime   float64     `json:"request_time" reindex:"request_time,tree"`
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
	Name          string      `json:"name"`
	Surname       string      `json:"surname"`
	RequestTime   float64     `json:"request_time"`
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

type JWTCustomClaims struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	jwt.RegisteredClaims
}

type GetInfo struct {
	First_end
	Second_end
	Third_end
}

type First_end struct {
	Class string  `json:"class"`
	Score float64 `json:"score"`
}

type Second_end struct {
	Priorate string  `json:"priorate"`
	Score    float64 `json:"score"`
}

type Third_end struct {
	Category string  `json:"category"`
	Score    float64 `json:"score"`
}
