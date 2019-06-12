package models

import (
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"

	"github.com/tidwall/gjson"
)

var dClient = &http.Client{}

func queryApiForZs(url, tag string) (bool, *gjson.Result) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")

	resp, err := dClient.Do(request)

	if err != nil {
		return false, nil
	}

	robots, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	respJsonDatas := gjson.ParseBytes(robots)

	if respJsonDatas.Get("ok").Int() != 1 {
		var msg string
		if respJsonDatas.Get("msg").Exists() {
			msg = respJsonDatas.Get("msg").String()
			beego.Debug(msg)
		}
		return false, nil
	}

	if !respJsonDatas.Get("data").Exists() {
		return false, nil
	}
	respJsonDatas = respJsonDatas.Get("data")

	return true, &respJsonDatas
}
