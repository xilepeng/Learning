package models

import (
	"fmt"
	"time"
)

// model 公共的方法
// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println("时间戳：", timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
