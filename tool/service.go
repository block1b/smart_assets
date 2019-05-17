package tool

import (
	"encoding/json"
	"fmt"
)

// 提供的服务
// 类似于api的定义
func InitApi() error {
	var err error
	// 获取设备信息
	err = Sub(CLIENTID+"/iotInfo", IotInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 创建设备
	err = Sub(CLIENTID+"/newIot", NewIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 修改设备
	err = Sub(CLIENTID+"/updateIot", UseIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 租用设备
	err = Sub(CLIENTID+"/rentIot", UseIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 归还设备
	err = Sub(CLIENTID+"/returnIot", UseIotPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 查看账单
	err = Sub(CLIENTID+"/billInfo", BillInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 查看余额
	err = Sub(CLIENTID+"/balanceInfo", BalanceInfoPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 充值提现
	err = Sub(CLIENTID+"/useBalance", UserBalancePubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}

	// +/post
	err = Sub("+/post", PostPubHandler)
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
func PostWork(clientId string, transfer TransferPrepare) error {
	ct := ClientIdTransfer{
		ClientId:clientId,
		TransferData:transfer,
	}
	ctByte, err := json.Marshal(ct)
	if err != nil {
		fmt.Println("marshal", err)
		return err
	}
	err = SubPub(clientId+"/post", PostPubHandler, "postServer/post", ctByte)
	if err != nil {
		fmt.Println("SubPub", err)
		return err
	}
	return nil
}
