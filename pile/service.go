package tool

import (
	"fmt"
	"time"

	. "smart_assets/tool"
)


// 提供的服务
// 类似于api的定义
func InitPileApi() error {
	time.Now()
	var err error
	// 获取设备信息
	err = Sub("/iotInfo", IotInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 创建设备
	err = Sub("/newInfo", NewIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}

	// 租用设备
	err = Sub("/rentIot", UseIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 归还设备
	err = Sub("/returnIot", UseIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}

	// 修改设备状态
	// todo 跟租用差不多

	// 查看账单
	err = Sub("/billInfo", BillInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 查看余额
	err = Sub("/balanceInfo", BalanceInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 充值提现
	err = Sub("/userBalance", UserBalancePubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}

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
