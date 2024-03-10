package main

import (
	"fmt"
	"grpc/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "x",
		Age:      18,
	}
	// 序列化的过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	// 反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())
}
