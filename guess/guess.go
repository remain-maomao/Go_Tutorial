package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 生成一个1-100的随机数
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	// fmt.Println(secretNumber)

	// 读取用户的输入
	fmt.Println("请输入你的猜测")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("读取输入时发生错误，请重试", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("无效的输入，请重试")
			continue
		}

		fmt.Println("你的猜测是", guess)

		// 实现判断逻辑
		if guess > secretNumber {
			fmt.Println("你猜的数字太大了，再小一点")
		} else if guess < secretNumber {
			fmt.Println("你猜的数字太小了，再大一点")
		} else {
			fmt.Println("你猜的数字刚刚好")
			break
		}
	}
}
