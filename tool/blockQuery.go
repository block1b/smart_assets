package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// http get 查询 bigchaindb 的查询接口

func GetAsset(){
	params := url.Values{}

	Url, err := url.Parse("http://192.168.113.6:9984/api/v1/assets")
	if err != nil {
		panic(err.Error())
	}
	params.Set("search", "余额2")

	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}