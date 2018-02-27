package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)



var (
	Conf Config
	Req Request
	Res Response
)

func init() {
	byteConf, err := ioutil.ReadFile("conf/baidu.json")

	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(byteConf, &Conf)

	if err != nil {
		fmt.Print(err)
	}
}
