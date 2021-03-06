package routers

import (
	"go-websocket/api/bind2group"
	"go-websocket/api/login"
	"go-websocket/api/register"
	"go-websocket/api/send2client"
	"go-websocket/api/send2group"
	"go-websocket/servers"
	"net/http"
)

func Init() {
	registerHandler := &register.Controller{}
	loginHandler := &login.Controller{}
	pushToClientHandler := &send2client.Controller{}
	pushToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}

	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/login", loginHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(pushToClientHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(pushToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))

	servers.StartWebSocket()

	go servers.WriteMessage()
}
