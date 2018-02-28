package main

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
)

func main() {

	var excludeArr []string

	if Conf.Exclude != "" {
		//如果有需要排除的文件后缀名，则对其进行切割形成数组，稍后使用
		excludeArr = strings.Split(Conf.Exclude, ",")
	}

	// 扫描input文件夹
	filepath.Walk(Conf.Input,
		//逐个读取文件
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}

			// 如果读取的文件路径是文件夹，则在output文件夹下创建对应语言的文件夹，
			targetPath := strings.Replace(path, Conf.Input, Conf.Output, -1)

			if f.IsDir() {

				err := os.MkdirAll(targetPath, 0777)

				if err != nil {
					fmt.Printf("%s", err)
				}

			} else {
				//如果不是文件夹，则开始读取文件

				if !strings.Contains(path, ".DS_Store") &&
					!strings.Contains(path, ".gitkeep") {

					//fmt.Println("file:", path)
					//fmt.Println("name:", f.Name())

					//如果没有设置排除
					if len(excludeArr) != 0 {
						//遍历排除项
						for _, value := range excludeArr {
							//如果当前的文件符合排除条件
							if strings.Contains(path, value) {
								fmt.Println("排除不翻译的文件：", f.Name())
								//直接拷贝文件
								CopyFile(targetPath, path)
								//结束，进行下一个文件的处理
								return nil
							}
						}

					}

					fmt.Println("开始翻译文件：", f.Name())

					//如果没有需要排除的文件后缀名
					//读取文件，用正则匹配所有的中文片段，每个片段逐个调用翻译函数
					Grab(path, targetPath)

				}

			}

			return nil
		})

	fmt.Println("所有翻译完成!")

}
