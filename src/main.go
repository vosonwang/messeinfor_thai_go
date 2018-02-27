package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func main() {

	Grap()

	Req.Query = "测试"
	Req.From = "zh"
	Req.To = "th"
	Req.Salt = 1435660288

	/*获取Sign*/
	str1 := strconv.Itoa(Conf.AppID) + Req.Query + strconv.Itoa(Req.Salt) + Conf.Key;

	//fmt.Print(str1+"\n")

	//计算Sign的md5
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str1))
	cipherStr := md5Ctx.Sum(nil)

	Req.Sign = hex.EncodeToString(cipherStr)

	t := fmt.Sprintf("?q=%s&from=%s&to=%s&appid=%d&salt=%d&sign=%s", Req.Query, Req.From, Req.To, Conf.AppID, Req.Salt, Req.Sign)

	resp, err := http.Get(Conf.URL + t)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Print(string(body)+"\n")

	err = json.Unmarshal(body, &Res)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(Res)
}
