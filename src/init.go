package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)



var (
	C Config
	Req Request
	Res Response
)

func init() {
	byteConf, err := ioutil.ReadFile("conf/baidu.json")

	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(byteConf, &C)

	if err != nil {
		fmt.Print(err)
	}
}
