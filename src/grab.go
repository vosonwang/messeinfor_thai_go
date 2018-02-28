package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"strings"
)

// 抓取文本中需要翻译的中文
func Grab(filePath, fileName string) {
	//读取文件
	byteStr, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Print(err)
	}

	//用正则匹配所有的中文片段
	r, _ := regexp.Compile("([\u4e00-\u9fa5])+")

	//每个片段逐个调用翻译函数，接收翻译结果，替换中文片段
	out := r.ReplaceAllFunc(byteStr, transfer)
	//fmt.Print(string(out))

	//用output文件夹替换input的文件夹名，获取目标路径
	targetPath := strings.Replace(filePath, Conf.Input, Conf.Output, -1)

	generate(out, targetPath, fileName)
}
