package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)



func main() {

	Req.Query = "apple"
	Req.From = "en"
	Req.To = "zh"
	Req.Salt = 1435660288
	Req.Sign = "f89f9594663708c1605f3d736d01d2d4"

	t := fmt.Sprintf("?q=%s&from=%s&to=%s&appid=%d&salt=%d&sign=%s", Req.Query, Req.From, Req.To, C.AppID, Req.Salt, Req.Sign)


	resp, err := http.Get(C.URL+t)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &Res)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(Res)
}
