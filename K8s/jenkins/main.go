package main

import "fmt"

func main() {
	var lixi int
	lixi = 50000 * 0.045 * 2
	total := lixi + 50000
	fmt.Printf("5W元2年总利息 = 50000 * 0.045 * 2 = %d\n", lixi)
	fmt.Printf("本金+利息 = 50000 + %d = %d 元人民币\n", lixi, total)
	fmt.Printf("截止日期：2023年12月30日 需还清 = %d + 5000 = %d 元人民币\n", total, total+5000)
}


5W元2年总利息 = 50000 * 0.045 * 2 = 4500
本金+利息 = 50000 + 4500 = 54500 元人民币
截止日期：2023年12月30日 需还清 = 54500 + 5000 = 59500 元人民币