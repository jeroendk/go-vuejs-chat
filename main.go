package main

import (
	"context"
	"flag"
	"log"
	"net/http"

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

	wsServer := NewWebsocketServer(&repository.RoomRepository{Db: db}, &repository.UserRepository{Db: db})
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
