package main

import (
	. "demo01/Services"
	"demo01/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httptranspsrt "github.com/ go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
)

func main() {
	user := UserService{}
	endp := GenUserEndPoint(user)

	serverHanlder := httptranspsrt.NewServer(endp, DecodeUserRequest, EncodeUserResponse)

	r := mymux.NewRouter()

	// r.Handle(`/user/{uid:\d+}`, serverHanlder)
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHanlder)

	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})


	//localhost:8080/user/522

	errChan := make(chan error)

	 go func(){
		util.RegService()//注册服务
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Println(err)
			errChan<-err
		}
	} ()

	 go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan<-fmt.Errorf("%s", <-sigChan)
	}()

	getErr := <-errChan //只要报错 或者service关闭阻塞在这里的会进行下去
	util.UnRegService()
	log.Println(getErr)
}


