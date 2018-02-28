package main

type Response struct {
	From string
	To   string
	TransResult `json:"trans_result"`
}

type TransResult []Result

type Result struct {
	Src string
	Dst string
}

type Config struct {
	Input  string
	Output string
	From   string
	To     string
	URL    string `json:"url"`
	SSL    string `json:"ssl"`
	AppID  int    `json:"app_id"`
	Key    string `json:"key"`
}
