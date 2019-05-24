package tool

import (
	"encoding/json"
	"fmt"
	"strings"
)

import MQTT  "github.com/eclipse/paho.mqtt.golang"

//define a function for the default message handler
var DefaultPublishHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[DefaultPublishHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
}

var OnConnectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Printf("connected \n")
}

var ConnectionLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("lost connect , err: %v\n", err)
}

// pub pkg handler +/post
var PostPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[PostPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	// 等待postServer的响应，理论上应该再转发给user的，但是异步的需要追踪
	// +/post
	subTopic := fmt.Sprintf("%v",msg.Topic())
	sts := strings.Split(subTopic, "/")
	pubClientId := sts[0]
	// 只处理create transfer的响应 Transaction
	var createResult TilfilledTransaction
	err := json.Unmarshal(msg.Payload(), &createResult)
	if err != nil {
		fmt.Println("[PostPubHandler] unmarshal", err)
		return
	}
	if createResult.Operation == "CREATE"{
		ss := strings.Split(createResult.Asset.Data.Sn,".")
		assetType := ss[1]
		if assetType == "balance"{
			// 跟新 +/ClientId/balanceAssetId {balanceAssetId:createResult.Id}
			// pub
			balanceId := struct {
				BalanceAssetId string `json:"balance_asset_id"`
			}{createResult.Id}
			bb, err := json.Marshal(balanceId)
			err = Pub(pubClientId+"/balanceAssetId", bb)
			if err != nil{
				fmt.Println(err)
			}
		}
		if assetType == "iot"{
			// 跟新 +/iotAssetId {iotAssetId:createResult.Id}
			// pub
			iotId := struct {
				IotAssetId string `json:"iot_asset_id"`
			}{createResult.Id}
			ib, err := json.Marshal(iotId)
			err = Pub(pubClientId+"/iotAssetId", ib)
			if err != nil{
				fmt.Println(err)
			}
		}
	}
}

// 查看设备信息
// IotInfoPubHandler
var IotInfoPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[IotInfoPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	// 查到设备信息后，pub给user，所以用户的请求结构，必须带clientId
	//func GetIotInfo(args NickForm) (DeviceForm, error)
	var getIotInfoForm GetIotInfoForm
	err := json.Unmarshal(msg.Payload(), &getIotInfoForm)
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
	iotInfo, err := GetIotInfo(getIotInfoForm.Iot)
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
	iotInfoByte, err := json.Marshal(iotInfo)
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, getIotInfoForm.ClientId, 1)
	err = Pub(pubTopic, iotInfoByte)
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
}

// 创建设备
//NewIotPubHandler
var NewIotPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[NewIotPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	//func CreateDevice(deviceForm DeviceForm) error
	repPayload := "ok"
	var newDeviceForm NewDeviceForm
	err := json.Unmarshal(msg.Payload(), &newDeviceForm)
	if err != nil{
		fmt.Println("NewIotPubHandler :", err)
		repPayload = "err"
	}
	createTransfer,err := CreateDevice(newDeviceForm.DeviceForm)
	if err != nil{
		fmt.Println("NewIotPubHandler :", err)
		repPayload = "err"
	}
	// post
	err = PostWork(newDeviceForm.ClientId,createTransfer)
	if err != nil {
		repPayload = "err"

	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, newDeviceForm.ClientId, 1)
	err = Pub(pubTopic, []byte(repPayload))
	if err != nil{
		fmt.Println("NewIotPubHandler :", err)
	}
}

// 租用设备（归还）
//RentIotPubHandler
var UseIotPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[RentIotPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	// UseIot(user NickForm, iotForm DeviceForm) error
	repPayload := "ok"
	var userIotForm UserIotForm
	err := json.Unmarshal(msg.Payload(), &userIotForm)
	if err != nil{
		fmt.Println("RentIotPubHandler :", err)
		repPayload = "err"
		return
	}
	balanceTransfer,iotTransfer, err := UseIot(userIotForm.User, userIotForm.Iot)
	if err != nil{
		fmt.Println("RentIotPubHandler :", err)
		repPayload = "err"
	}
	fmt.Println(balanceTransfer, iotTransfer)
	if userIotForm.Iot.Status == "Return"{
		fmt.Println("归还支付")
		// 只有归还操作，才有金钱交易
		err = PostWork(userIotForm.ClientId,balanceTransfer)
		if err != nil {
			repPayload = "err"
		}
	}
	// post
	err = PostWork(userIotForm.ClientId,iotTransfer)
	if err != nil {
		repPayload = "err"
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, userIotForm.ClientId, 1)
	err = Pub(pubTopic, []byte(repPayload))
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
}

// 修改设备状态
// 同租用


// 查看个人账单
// BillInfoPubHandler
var BillInfoPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[BillInfoPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	var getUerBillsForm GetUerBillsForm
	err := json.Unmarshal(msg.Payload(), &getUerBillsForm)
	if err != nil{
		fmt.Println("BillInfoPubHandler :", err)
	}
	userBills, err := GetPersonBills(getUerBillsForm.User)
	if err != nil{
		fmt.Println("BillInfoPubHandler :", err)
	}
	userBillsByte, err := json.Marshal(userBills)
	if err != nil{
		fmt.Println("marshal :", err)
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, getUerBillsForm.ClientId, 1)
	err = Pub(pubTopic, userBillsByte)
	if err != nil{
		fmt.Println("BillInfoPubHandler :", err)
	}
}

// 查看余额
//BalanceInfoPubHandler
var BalanceInfoPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[BalanceInfoPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	var getUerBalanceForm GetUerBalanceForm
	err := json.Unmarshal(msg.Payload(), &getUerBalanceForm)
	if err != nil{
		fmt.Println("BalanceInfoPubHandler :", err)
	}
	//userBalanceOutput,userBalanceOutputResult, err := OutputQuery(getUerBalanceForm.User)
	userBalanceOutput,userBalanceOutputResult, err := GetBalanceOutputs(getUerBalanceForm.User)
	if err != nil{
		fmt.Println("BalanceInfoPubHandler :", err)
		errType := fmt.Sprint(err)
		if errType == "noWallet"{
			// 主钱包初始化
			transferPrepare, err := InitWallet()
			if err != nil{
				fmt.Println("merge balance", err)
			}
			// post
			err = PostWork(getUerBalanceForm.ClientId,transferPrepare)
			if err != nil {
				fmt.Println("post", err)
			}
		}

		if errType == "unMerge"{
			// 合并balance
			//func MergeBalanceAsset(args NickForm,outPutResults []GetOutputResult) (TransferPrepare,error)
			transferPrepare, err := MergeBalanceAsset(getUerBalanceForm.User,userBalanceOutputResult)
			if err != nil{
				fmt.Println("merge balance", err)
			}
			// post
			err = PostWork(getUerBalanceForm.ClientId,transferPrepare)
			if err != nil {
				fmt.Println("post", err)
			}
		}

	}
	// 封装 amount
	amount := BalanceAmount{Amount:userBalanceOutput.Amount}
	amountByte, err := json.Marshal(amount)
	if err != nil{
		fmt.Println("marshal", err)
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, getUerBalanceForm.ClientId, 1)
	err = Pub(pubTopic, amountByte)
	if err != nil{
		fmt.Println("BalanceInfoPubHandler :", err)
	}
}

// 充值/提现
//UserBalancePubHandler
var UserBalancePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[UserBalancePubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	//func UseBalance(args UseBalanceForm, bUser NickForm) error
	repPayload := "ok"
	var useMoneyForm UseMoneyForm
	err := json.Unmarshal(msg.Payload(), &useMoneyForm)
	if err != nil{
		fmt.Println("UserBalancePubHandler :", err)
		repPayload = "err"
	}
	balanceTransfer, err := UseBalance(useMoneyForm.AUser, useMoneyForm.BUser)
	if err != nil{
		fmt.Println("UserBalancePubHandler :", err)
		repPayload = "err"
	}
	// post
	err = PostWork(useMoneyForm.ClientId,balanceTransfer)
	if err != nil {
		return
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, useMoneyForm.ClientId, 1)
	err = Pub(pubTopic, []byte(repPayload))
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
}
