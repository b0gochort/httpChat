package main

import (
	"fmt"
	"github.com/b0gochort/httpChat/internal/api_db"
	"github.com/b0gochort/httpChat/internal/handler"
	"github.com/b0gochort/httpChat/internal/service"
	"github.com/restream/reindexer/v3"

	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/valyala/fasthttp"
)

func main() {

	db := reindexer.NewReindex("cproto://192.168.10.111:6534/center")

	if db.Status().Err != nil {
		panic(db.Status().Err)
	}

	apiDB := api_db.NewAPIDB(db)
	s := service.NewService(apiDB)
	h := handler.NewHandler(s)

	server := fasthttp.Server{
		Handler: h.InitRoutes,
	}

	fmt.Println("server was started!")

	err := server.ListenAndServe(":3305")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
