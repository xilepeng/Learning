package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *consulapi.Client

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.1.2:8500"
	client, err := consulapi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
}

func RegService() {
	reg := consulapi.AgentServiceRegistration{}
	reg.Name = "userservice" //注册service的名字
	reg.Address = "192.168.1.2" //注册service的ip
	reg.Port = 8080//注册service的端口
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{} //创建consul的检查器
	check.Interval="5s" //设置consul心跳检查时间间隔
	check.HTTP = "http://192.168.1.2:8080/health" //设置检查使用的url

	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func UnRegService()  {
	ConsulClient.Agent().ServiceDeregister("userservice")
}