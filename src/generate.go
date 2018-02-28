package main

import (
	"os"
	"fmt"
	"log"
)

func generate(b []byte, targetPath, fileName string) {
	var f *os.File

	defer f.Close()

	/*判断文件是否存在，如果存在则覆盖，没有则创建（没有考虑文件夹是否存在）*/
	if Exists(targetPath) {
		if t, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0); err != nil {
			log.Print(err)
			fmt.Print("打开文件失败！")
		} else {
			f = t
		}

	} else {

		if t, err := os.Create(targetPath); err != nil {
			log.Print(err)
			fmt.Print("创建文件失败！")
		} else {
			f = t
		}
	}

	_, err := f.Write(b)

	if err != nil {
		log.Print(err)
		fmt.Print("写入文件失败！")
	}
}

/*
判断文件是否存在
 */
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
