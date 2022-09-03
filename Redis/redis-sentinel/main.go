package main
 
import (
        "fmt"
        "github.com/ go-redis/redis"
        "time"
)
 
func main() {
        client := redis.NewFailoverClient(&redis.FailoverOptions{
                MasterName:    "mymaster",
                SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
                Password:      "",
                DB:            0,
        })
 
        for {
                reply, err := client.Incr("pvcount").Result()
                fmt.Printf("reply=%v err=%v\n", reply, err)
                time.Sleep(1 * time.Second)
        }
}

