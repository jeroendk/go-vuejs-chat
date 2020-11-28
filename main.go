package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/jeroendk/chatApplication/auth"
	"github.com/jeroendk/chatApplication/config"
	"github.com/jeroendk/chatApplication/repository"
)

var addr = flag.String("addr", ":8080", "http server address")
var ctx = context.Background()

func main() {
	flag.Parse()

	config.CreateRedisClient()
	db := config.InitDB()
	defer db.Close()

	userRepository := &repository.UserRepository{Db: db}

	wsServer := NewWebsocketServer(&repository.RoomRepository{Db: db}, userRepository)
	go wsServer.Run()

	api := &API{UserRepository: userRepository}

	http.HandleFunc("/ws", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	}))

	http.HandleFunc("/api/login", api.HandleLogin)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
