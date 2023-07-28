package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个1-100的随机数
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
}
