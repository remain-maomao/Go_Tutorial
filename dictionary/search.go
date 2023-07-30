package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserId    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}

	// 创建一个请求对象
	request := DictRequest{TransType: "en2zh", Source: word}
	// 将请求对象序列化为json字符串
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.757739313.1690550639; _ga_65TZCJSDBD=GS1.1.1690550638.1.0.1690550639.0.0.0; _ga_R9YPR75N68=GS1.1.1690550638.1.0.1690550639.59.0.0; _ga=GA1.2.2096512254.1690550639; _gid=GA1.2.1395040821.1690550642")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "02c95aaa9ec45900cfdd21cdaa76323b")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 防御式编程，检查响应的状态码
	if resp.StatusCode != 200 {
		fmt.Println("bad status code", resp.StatusCode, "body", string(bodyText))
	}
	// 将从服务器接受到的字节流形式的json数据反序列化到DictResponse对象中
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	// 打印结果
	fmt.Println(dictResponse.Dictionary.Entry, "en-us", dictResponse.Dictionary.Prons.EnUs, "en", dictResponse.Dictionary.Prons.En)
	for _, explanation := range dictResponse.Dictionary.Explanations {
		fmt.Println(explanation)
	}
}

func main() {
	if len(os.Args) != 2 {
		_, err := fmt.Fprintf(os.Stderr, `usage: search <WORD>
example: search pretty
`)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	word := os.Args[1]
	query(word)
}
