package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"flag"
)

var (
	Conf Config
	Res  Response

)

func init() {
	/*log*/
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)


	var (
		byteConf []byte
		err      error
	)

	Debug := flag.Bool("debug", false, "-debug")

	flag.Parse()

	//fmt.Println(*Debug)

	if *Debug {
		/*读取配置文件*/
		byteConf, err = ioutil.ReadFile("conf/dev.json")
		fmt.Println("使用百度测试账户：")
	} else {
		byteConf, err = ioutil.ReadFile("conf/prod.json")
		fmt.Println("使用百度正式账户：")
	}

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteConf, &Conf)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("百度账户AppID：",Conf.AppID)

}
