package main

import (
	"context"
	."demo01/Services"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"os"
)
func main()  {
	tag,_ := url.Parse("http://127.0.0.1:8080")
	//创建一个直连client，这里我们必须写两个func,一个是如何请求,一个是响应我们怎么处理
	client := httptransport.NewClient("GET", tag,GetUserInfo_Request, GetUserInfo_Response)

	//通过这个拿到了定义在服务端的endpoint也就是上面这段代码return出来的函数，直接在本地就可以调用服务端的代码
	getUserInfo := client.Endpoint()
	//创建一个上下文
	ctx := context.Background()
	//执行
	res,err := getUserInfo(ctx, UserRequest{Uid: 52})//使用go-kit插件来直接调用服务
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := res.(UserResponse)
	fmt.Println(userinfo.Result)
}
