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

// pub pkg handler
var PostPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[PostPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	// 等待postServer的响应，理论上应该再转发给user的，但是异步的需要追踪 todo

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
//NewInfoPubHandler
var NewInfoPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[NewInfoPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	//func CreateDevice(deviceForm DeviceForm) error
	repPayload := "ok"
	var newDeviceForm NewDeviceForm
	err := json.Unmarshal(msg.Payload(), &newDeviceForm)
	if err != nil{
		fmt.Println("NewInfoPubHandler :", err)
		repPayload = "err"
		return
	}
	err = CreateDevice(newDeviceForm.DeviceForm)
	if err != nil{
		fmt.Println("NewInfoPubHandler :", err)
		repPayload = "err"
		return
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, newDeviceForm.ClientId, 1)
	err = Pub(pubTopic, []byte(repPayload))
	if err != nil{
		fmt.Println("NewInfoPubHandler :", err)
		return
	}
}

// 租用设备
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
	}
	err = UseIot(userIotForm.User, userIotForm.Iot)
	if err != nil{
		fmt.Println("RentIotPubHandler :", err)
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
	userBillsByte, err := GetPersonBills(getUerBillsForm.User)
	if err != nil{
		fmt.Println("BillInfoPubHandler :", err)
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
	userBalanceOutput,_, err := OutputQuery(getUerBalanceForm.User)
	if err != nil{
		fmt.Println("BalanceInfoPubHandler :", err)
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, getUerBalanceForm.ClientId, 1)
	err = Pub(pubTopic, []byte(userBalanceOutput.Amount))
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
	err = UseBalance(useMoneyForm.AUser, useMoneyForm.BUser)
	if err != nil{
		fmt.Println("UserBalancePubHandler :", err)
		repPayload = "err"
	}
	// pub
	pubTopic := strings.Replace(msg.Topic(), CLIENTID, useMoneyForm.ClientId, 1)
	err = Pub(pubTopic, []byte(repPayload))
	if err != nil{
		fmt.Println("IotInfoPubHandler :", err)
	}
}
