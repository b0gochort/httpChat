package api_db

import (
	"fmt"
	"github.com/b0gochort/httpChat/internal/model"
	"github.com/restream/reindexer/v3"
)

type ChatAPIImpl struct {
	db *reindexer.Reindexer
}

func NewChatApi(db *reindexer.Reindexer) *ChatAPIImpl {
	return &ChatAPIImpl{
		db: db,
	}
}

func (a *ChatAPIImpl) CreateChat(room model.NewChatItem) (model.NewChatItem, error) {

	err := a.db.OpenNamespace("support_chat", reindexer.DefaultNamespaceOptions(), model.NewChatItem{})
	if err != nil {
		return model.NewChatItem{}, fmt.Errorf("chatApi.CreateChat.OpenNamespace: %v", err)
	}

	ok, err := a.db.Insert("support_chat", &room, "id=serial()")
	if err != nil {
		return model.NewChatItem{}, fmt.Errorf("chatApi.CreateChat.db.Insert: %v", err)
	}

	if ok == 0 {
		return model.NewChatItem{}, fmt.Errorf("nil insert")
	}

	return room, nil
}

func (a *ChatAPIImpl) NewMessage(message model.MessageItem) (model.MessageItem, error) {
	err := a.db.OpenNamespace("support_message", reindexer.DefaultNamespaceOptions(), model.MessageItem{})
	if err != nil {
		return model.MessageItem{}, fmt.Errorf("chatApi.NewMessage.OpenNamespace: %v", err)
	}

	err = a.db.OpenNamespace("support_chat", reindexer.DefaultNamespaceOptions(), model.NewChatItem{})
	if err != nil {
		return model.MessageItem{}, fmt.Errorf("chatApi.NewMessage.OpenNamespace: %v", err)
	}

	ok, err := a.db.Insert("support_message", &message, "id=serial()")
	if err != nil {
		return model.MessageItem{}, fmt.Errorf("chatApi.NewMessage.db.Insert: %v", err)
	}

	if ok == 0 {
		return model.MessageItem{}, fmt.Errorf("nil insert")
	}

	query := a.db.Query("support_chat").Where("id", reindexer.EQ, message.ChatId).Set("message.last_message", message.Text).Update()
	if query.Error() != nil {
		return model.MessageItem{}, fmt.Errorf("db.Query.Update: %v", query.Error())
	}

	return message, nil
}

func (a *ChatAPIImpl) GetMessage(chatId int64) ([]model.MessageItem, error) {
	err := a.db.OpenNamespace("support_message", reindexer.DefaultNamespaceOptions(), model.MessageItem{})
	if err != nil {
		return nil, fmt.Errorf("chatApi.GetMessage.OpenNamespace: %v", err)
	}
	elem := a.db.Query("support_message").Sort("time", false).Where("chat_id", reindexer.EQ, chatId)

	var response []model.MessageItem

	iter := elem.Exec()
	if iter.Error() != nil {
		return nil, fmt.Errorf("chatApi.GetMessage.Exec: %v", iter.Error())
	}

	for iter.Next() {
		elem := iter.Object().(*model.MessageItem)
		response = append(response, model.MessageItem{
			ID:     elem.ID,
			ChatId: elem.ChatId,
			Time:   elem.Time,
			Host:   elem.Host,
			Text:   elem.Text,
		})
	}

	return response, nil
}

func (a *ChatAPIImpl) GetRooms() ([]model.NewChatItem, error) {
	err := a.db.OpenNamespace("support_chat", reindexer.DefaultNamespaceOptions(), model.MessageItem{})
	if err != nil {
		return nil, fmt.Errorf("chatApi.GetMessage.OpenNamespace: %v", err)
	}
	elem := a.db.Query("support_chat").Sort("time", false)

	var response []model.NewChatItem

	iter := elem.Exec()
	if iter.Error() != nil {
		return nil, fmt.Errorf("chatApi.GetMessage.Exec: %v", iter.Error())
	}

	for iter.Next() {
		elem := iter.Object().(*model.NewChatItem)
		response = append(response, model.NewChatItem{
			ID:            elem.ID,
			Time:          elem.Time,
			Message:       elem.Message,
			LastHostStaff: elem.LastHostStaff,
			UID:           elem.UID,
			IP:            elem.IP,
			Category:      elem.Category,
		})
	}

	return response, nil
}
