package handler

import (
	"encoding/json"
	"fmt"
	"github.com/b0gochort/httpChat/internal/model"

	"github.com/valyala/fasthttp"
	"resenje.org/logging"
	"strconv"
)

func (h *Handler) NewChat(ctx *fasthttp.RequestCtx) {
	var req model.NewRoomReq

	if !ctx.IsPost() {
		logging.Info("")
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.Unmarshal: %v", err))
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	userId := int64(1)
	userIp := string(ctx.Request.Header.Peek("x-forwarded-for"))
	chat, err := h.services.ChatService.NewChat(userId, req.Message, userIp)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.%v", err))
		ctx.Error("error creating chat", fasthttp.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(chat)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)
		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) SendMessage(ctx *fasthttp.RequestCtx) {
	var req model.NewMessageReq

	if !ctx.IsPost() {
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		logging.Info(fmt.Sprintf("handler.SendMessage.Unmarshal: %v", err))
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	userId := int64(1)
	sub := "user"
	messageId, err := h.services.NewMessage(req.ChatID, userId, req.Message, sub)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.SendMessage.NewMessage.%v", err))
		ctx.Error("error creating message", fasthttp.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(messageId)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) GetMessages(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	q := ctx.QueryArgs()

	chatId, err := strconv.Atoi(string(q.Peek("chat-id")))
	if err != nil {
		logging.Info("handler.strconv.Atoi: %v", err)
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
	}

	messages, err := h.services.GetMessage(int64(chatId))
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetMessages.GetMessage.%v", err))
		ctx.Error("error creating message", fasthttp.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(messages)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetMessages.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}
