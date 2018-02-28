package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
)



var (
	Conf Config
	Res Response
)

func init() {
	/*log*/
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	byteConf, err := ioutil.ReadFile("conf/config.json")

	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(byteConf, &Conf)

	if err != nil {
		fmt.Print(err)
	}
}
