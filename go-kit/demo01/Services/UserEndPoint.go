package Services

import (
	"context"
	"github.com/ go-kit/kit/endpoint"
)

type UserRequest struct { //封装User请求结构体
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndPoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		r := request.(UserRequest)
		result := userService.GetName(r.Uid)
		return UserResponse{Result: result}, nil
	}
}
