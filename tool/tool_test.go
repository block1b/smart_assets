package tool

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGetAsset(t *testing.T) {
	asset, _:= GetAsset("余额2")
	fmt.Println(string(asset))
}

type User struct {
	Id string `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
} 

func TestStruct(t *testing.T )  {
	byteData := []byte(`{"name":"han"}`)
	var user User
	_ = json.Unmarshal(byteData, &user)

	fmt.Println(user)
	if user.Id == "1"{
		fmt.Println("aaa")
	}else {
		fmt.Println("undefine")
	}

}

func TestTime(t *testing.T)  {
	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
	fmt.Println(tm.Format("02/01/2006 15:04:05"))
	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())

	ft := time.Now().Format("2006-01-02 03:04:05")
	fmt.Println(ft)
	ft2, _ := time.Parse("2006-01-02 03:04:05", ft)
	fmt.Println(ft2)

	a, _ := time.Parse("2006-01-02 03:04:05", "2019-01-02 03:04:05")
	fmt.Println(a)
	b, _ := time.Parse("2006-01-02 03:04:05", "2019-01-02 03:05:05")
	fmt.Println(b)
	c := b.Unix() - a.Unix()
	fmt.Println(c, string(c))
	d := time.Unix(c, 0)
	fmt.Println(d, " ", d.Format("2006-01-02 03:04:05"))


}