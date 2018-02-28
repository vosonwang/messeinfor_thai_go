package main

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
)

func main() {

	// 扫描input文件夹
	filepath.Walk(Conf.Input,
		//逐个读取文件
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				// 如果读取的文件路径是文件夹，则在output文件夹下创建对应语言的文件夹，
				targetPath := strings.Replace(path, Conf.Input, Conf.Output, -1)

				err := os.MkdirAll(targetPath, 0777)

				if err != nil {
					fmt.Printf("%s", err)
				} else {
					fmt.Print("Create Directory OK!")
				}

			} else {
				//如果不是文件夹，则开始读取文件

				if !strings.Contains(path, ".DS_Store") &&
					!strings.Contains(path, ".gitkeep") {

					//fmt.Println("file:", path)
					//fmt.Println("name:", f.Name())

					//读取文件，用正则匹配所有的中文片段，每个片段逐个调用翻译函数
					Grab(path, f.Name())
				}
			}

			return nil
		})

}
