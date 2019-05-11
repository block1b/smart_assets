package tool

import "fmt"

// 提供的服务
// 类似于api的定义
func InitApi() error {
	var err error
	// 获取设备信息
	err = Sub("/iotInfo", IotInfoPubHandler)
	//err = Sub("/xx", xxPubHandler)

	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	return nil
}

// event 向postServer提交transfer
func PostWork(payload []byte) error {
	err := SubPub("/post", PostPubHandler, "postServer/post", payload)
	if err != nil {
		fmt.Println("SubPub", err)
		return err
	}
	return nil
}
