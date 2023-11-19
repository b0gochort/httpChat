package service

import (
	"github.com/b0gochort/httpChat/internal/api_db"
	"github.com/b0gochort/httpChat/internal/model"
)

type ChatService interface {
	NewChat(userId, mId, priority int64, reqTime float64, message, userIp, category string) (model.NewChatRes, error)
	NewMessage(chatId, userID int64, inputMessage, sub string) (model.NewMessageRes, error)
	GetMessage(chatId int64) ([]model.Message, error)
	GetRooms() ([]model.NewChatRes, error)
	GetModerId(category string) (int64, error)
}

type Service struct {
	ChatService
}

func NewService(ApiDB *api_db.ApiDB) *Service {
	return &Service{
		ChatService: NewChatService(ApiDB.ChatAPI),
	}
}
