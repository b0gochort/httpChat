package handler

import (
	"encoding/json"
	"fmt"
	"github.com/b0gochort/httpChat/internal/model"
	"github.com/golang-jwt/jwt/v4"
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

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("salt"), nil
	})
	if err != nil {
		logging.Info(fmt.Sprintf("handler.jwt.Parse.%v", err))
		ctx.Error("error cjwt.Parse", fasthttp.StatusInternalServerError)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		logging.Info(fmt.Sprintf("handler.token.Claims.(*jwt.StandardClaims).%v", err))
		ctx.Error("error token.Claims.(*jwt.StandardClaims)", fasthttp.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(claims.Id)
	if err != nil {
		logging.Info("handler.strconv.Atoi: %v", err)
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
	}

	userIp := string(ctx.Request.Header.Peek("x-forwarded-for"))
	chat, err := h.services.ChatService.NewChat(int64(userId), req.Message, userIp)
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

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("salt"), nil
	})
	if err != nil {
		logging.Info(fmt.Sprintf("handler.jwt.Parse.%v", err))
		ctx.Error("error cjwt.Parse", fasthttp.StatusInternalServerError)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		logging.Info(fmt.Sprintf("handler.token.Claims.(*jwt.StandardClaims).%v", err))
		ctx.Error("error token.Claims.(*jwt.StandardClaims)", fasthttp.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(claims.Id)
	if err != nil {
		logging.Info("handler.strconv.Atoi: %v", err)
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
	}
	sub := claims.Subject
	messageId, err := h.services.NewMessage(req.ChatID, int64(userId), req.Message, sub)
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
