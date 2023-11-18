package api_db

import (
	"github.com/b0gochort/httpChat/internal/model"
	"github.com/restream/reindexer/v3"
)

type ChatAPI interface {
	CreateChat(room model.NewChatItem) (model.NewChatItem, error)
	NewMessage(message model.MessageItem) (model.MessageItem, error)
	GetMessage(chatId int64) ([]model.MessageItem, error)
	GetRooms() ([]model.NewChatItem, error)
}

type ApiDB struct {
	ChatAPI
}

func NewAPIDB(db *reindexer.Reindexer) *ApiDB {
	return &ApiDB{
		ChatAPI: NewChatApi(db),
	}
}
