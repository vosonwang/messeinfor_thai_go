package main

import (
	"time"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

// 翻译函数
func transfer(b []byte) []byte {

	query := string(b)
	salt := time.Now().Unix()

	/*获取Sign*/
	str1 := strconv.Itoa(Conf.AppID) + query + strconv.FormatInt(salt, 10) + Conf.Key

	//fmt.Println(str1+"\n")

	//计算Sign的md5
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str1))
	cipherStr := md5Ctx.Sum(nil)

	Sign := hex.EncodeToString(cipherStr)

	t := fmt.Sprintf("?q=%s&from=%s&to=%s&appid=%d&salt=%d&sign=%s", query, Conf.From, Conf.To, Conf.AppID, salt, Sign)


	resp, err := http.Get(Conf.URL + t)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body)+"\n")

	err = json.Unmarshal(body, &Res)

	if err != nil {
		fmt.Println(err)
	}

	return []byte(Res.TransResult[0].Dst)
}
