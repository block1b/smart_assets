package tool

import (
	"encoding/json"
	"fmt"
	"testing"
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