package main

type Request struct {
	Query    string
	From string
	To   string
	Salt int
	Sign string
}

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
	URL   string `json:"url"`
	SSL   string `json:"ssl"`
	AppID int    `json:"app_id"`
	Key   string `json:"key"`
}