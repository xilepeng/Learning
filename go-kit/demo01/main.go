package main

import (
	."demo01/Services"
	httptranspsrt "github.com/ go-kit/kit/transport/http"
	"net/http"
)

func main() {
	user := UserService{}
	endp := GenUserEndPoint(user)

	serverHanlder := httptranspsrt.NewServer(endp, DecodeUserRequest,EncodeUserResponse)

	http.ListenAndServe(":8080",serverHanlder)
	//localhost:8080/?uid=522
}

