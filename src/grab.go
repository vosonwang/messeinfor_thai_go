package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
)

// 抓取文本中需要翻译的中文
func Grab(filePath, targetPath string) {
	//读取文件
	byteStr, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	var out []byte

	//用正则匹配所有的中文片段
	r, _ := regexp.Compile("[\u4e00-\u9fa5]+")

	//每个片段逐个调用翻译函数，接收翻译结果，替换中文片段
	out = r.ReplaceAllFunc(byteStr, transfer)
	//fmt.Println(string(out))

	generate(out, targetPath)
}
