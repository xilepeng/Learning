package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("连接成功！")

	ok, err := conn.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)

	v, err := redis.String(conn.Do("GET", "hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	defer conn.Close()
}
