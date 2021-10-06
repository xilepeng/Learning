package Services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	mymux "github.com/gorilla/mux"
)

//localhost:8080/user/522
//这个函数决定了使用哪个request结构体来请求
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	vars := mymux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{Uid: uid, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

//设置响应格式为json，这样客户端接收到的值就是json，就是把我们设置的UserResponse给json化了
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response) //判断响应格式是否正确
}
