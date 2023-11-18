package service

import (
	"fmt"
	"github.com/b0gochort/httpChat/internal/api_db"
	"github.com/b0gochort/httpChat/internal/model"

	"time"
)

type ChatServiceImpl struct {
	chatAPI api_db.ChatAPI
}

func NewChatService(chatAPI api_db.ChatAPI) *ChatServiceImpl {
	return &ChatServiceImpl{
		chatAPI: chatAPI,
	}
}

func (s *ChatServiceImpl) NewChat(userId int64, message, userIp string) (model.NewChatRes, error) {
	chat := model.NewChatItem{
		Time: time.Now().Unix(),
		Message: model.MessageType{
			NumberOfUnread: 1,
			LastMessage:    message,
		},
		LastHostStaff: false,
		UID:           userId,
		IP:            userIp,
	}

	chat, err := s.chatAPI.CreateChat(chat)
	if err != nil {
		return model.NewChatRes{}, fmt.Errorf("service.NewChat.%v", err)
	}

	chatRes := model.NewChatRes{
		ID:            chat.ID,
		Time:          chat.Time,
		Message:       chat.Message,
		LastHostStaff: chat.LastHostStaff,
		UID:           chat.UID,
		IP:            chat.IP,
	}
	return chatRes, nil
}

func (s *ChatServiceImpl) NewMessage(chatId, userID int64, inputMessage, sub string) (model.NewMessageRes, error) {
	message := model.MessageItem{
		ChatId: chatId,
		Time:   time.Now().Unix(),
		Host: model.HostType{
			UserId: userID,
			Sub:    sub,
		},
		Text: inputMessage,
	}

	message, err := s.chatAPI.NewMessage(message)
	if err != nil {
		return model.NewMessageRes{}, fmt.Errorf("service.NewMessage.%v", err)
	}

	messageRes := model.NewMessageRes{
		ID: message.ID,
	}

	return messageRes, nil
}

func (s *ChatServiceImpl) GetMessage(chatId int64) ([]model.Message, error) {
	res := make([]model.Message, 0)
	messages, err := s.chatAPI.GetMessage(chatId)
	if err != nil {
		return nil, fmt.Errorf("service.GetMessage.%v", err)
	}

	for i := range messages {
		msg := model.Message{
			ID:     messages[i].ID,
			ChatId: messages[i].ChatId,
			Time:   messages[i].Time,
			Host:   messages[i].Host,
			Text:   messages[i].Text,
		}
		res = append(res, msg)
	}

	return res, nil
}

func (s *ChatServiceImpl) GetRooms() ([]model.NewChatRes, error) {
	res := make([]model.NewChatRes, 0)

	rooms, err := s.chatAPI.GetRooms()
	if err != nil {
		return nil, fmt.Errorf("service.chatAPI.GetRooms.%v", err)
	}
	for _, room := range rooms {
		room := model.NewChatRes{
			ID:            room.ID,
			Time:          room.Time,
			Message:       room.Message,
			LastHostStaff: room.LastHostStaff,
			UID:           room.UID,
			IP:            room.IP,
			Category:      room.Category,
		}

		res = append(res, room)
	}

	return res, nil
}
