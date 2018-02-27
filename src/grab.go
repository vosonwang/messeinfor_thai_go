package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
)

// 将中文替换成泰文
func transfer(b []byte) []byte{
	// TODO 调用翻译API  将中文替换成泰文
	return []byte(string(b)+"1")
}

func Grap() {
	byteStr, err := ioutil.ReadFile("input/zh-CN/AccountSecurity.html")

	if err != nil {
		fmt.Print(err)
	}

	r, _ := regexp.Compile("([\u4e00-\u9fa5])+")


	out:= r.ReplaceAllFunc(byteStr, transfer)

	fmt.Println(string(out))
}
