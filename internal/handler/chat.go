package handler

import (
	"encoding/json"
	"fmt"
	"github.com/b0gochort/httpChat/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
	"resenje.org/logging"
	"strconv"
	"time"
)

func (h *Handler) NewChat(ctx *fasthttp.RequestCtx) {
	var req model.NewChatReq

	start := time.Now().Unix()

	if !ctx.IsPost() {
		logging.Info("")
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	cToken := ctx.Request.Header.Cookie("token")

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.Unmarshal: %v", err))
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(cToken), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bot"), nil
	})

	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.ParseWithClaims: %v", err))
		ctx.Error("ParseWithClaims", fasthttp.StatusInternalServerError)
		return
	}

	strUserId, err := claims.GetAudience()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.GetAudience: %v", err))
		ctx.Error("GetAudience", fasthttp.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(strUserId[0])
	if err != nil {
		logging.Info("handler.strconv.Atoi: %v", err)
		ctx.Error("unprocessable entity", fasthttp.StatusInternalServerError)
	}
	time.Sleep(1 * time.Second)
	//CHAT
	timeRequest := float64(time.Now().Unix()) - float64(start)

	userIp := string(ctx.Request.Header.Peek("x-forwarded-for"))
	chat, err := h.services.ChatService.NewChat(int64(userId), timeRequest, req.Message, userIp)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.%v", err))
		ctx.Error("error creating chat", fasthttp.StatusInternalServerError)
		return
	}
	chat.Name = claims["name"].(string)
	chat.Surname = claims["surname"].(string)

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
		ctx.Error("error jwt.Parse", fasthttp.StatusInternalServerError)
		return
	}

	claims, ok := token.Claims.(*model.JWTCustomClaims)
	if !ok {
		logging.Info(fmt.Sprintf("handler.token.Claims.(*jwt.StandardClaims).%v", err))
		ctx.Error("error token.Claims.(*jwt.StandardClaims)", fasthttp.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(claims.RegisteredClaims.ID)
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

func (h *Handler) GetRooms(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	cToken := ctx.Request.Header.Cookie("token")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(cToken), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bot"), nil
	})

	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.ParseWithClaims: %v", err))
		ctx.Error("ParseWithClaims", fasthttp.StatusInternalServerError)
		return
	}

	sub, err := claims.GetSubject()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.NewChat.GetSubject: %v", err))
		ctx.Error("GetSubject", fasthttp.StatusInternalServerError)
		return
	}

	if sub == "user" {
		ctx.Error("user cant take all rooms", fasthttp.StatusForbidden)
		return
	}

	messages, err := h.services.GetRooms()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetMessages.GetRooms.%v", err))
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

//func getCategory(text string) {
//	var (
//		f model.First_end
//		s model.Second_end
//		t model.Third_end
//	)
//
//	c := &fasthttp.Client{
//		Dial: func(addr string) (net.Conn, error) {
//			return fasthttp.DialTimeout(addr, time.Second*10)
//		},
//		MaxConnsPerHost: 1,
//	}
//	code, body, err := c.Get(nil, fmt.Sprintf("127.0.0.1:5000/ai_first?text=\"%s\"", text))
//	if err != nil {
//		return
//	}
//	if code != 200 {
//		return
//	}
//	if err := json.Unmarshal(body, &f); err != nil {
//		return
//	}
//
//	code, body, err = c.Get(nil, fmt.Sprintf("127.0.0.1:5000/ml_secondary?text=\"%s\"", text))
//	if err != nil {
//		return
//	}
//	if code != 200 {
//		return
//	}
//	if err := json.Unmarshal(body, &f); err != nil {
//		return
//	}
//
//	code, body, err = c.Get(nil, fmt.Sprintf("127.0.0.1:5000/ml_two?text=\"%s\"", text))
//	if err != nil {
//		return
//	}
//	if code != 200 {
//		return
//	}
//	if err := json.Unmarshal(body, &f); err != nil {
//		return
//	}
//
//}
