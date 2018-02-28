package main

import (
	"os"
	"fmt"
	"log"
	"io"
)

func generate(b []byte, targetPath string) {
	var f *os.File

	defer func() {
		f.Close()
		fmt.Println("翻译结束:",targetPath)
	}()

	/*判断文件是否存在，如果存在则覆盖，没有则创建（没有考虑文件夹是否存在）*/
	if Exists(targetPath) {
		if t, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0); err != nil {
			log.Print(err)
			fmt.Println("打开文件失败！")
		} else {
			f = t
		}

	} else {

		if t, err := os.Create(targetPath); err != nil {
			log.Print(err)
			fmt.Println("创建文件失败！")
		} else {
			f = t
		}
	}

	_, err := f.Write(b)

	if err != nil {
		log.Print(err)
		fmt.Println("写入文件失败！")
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

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
