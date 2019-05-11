package tool

import "fmt"

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

// IotPubHandler
var IotInfoPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[IotInfoPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
	// 查到设备信息后，pub给user，所以用户的请求结构，必须带clientId， todo
	//func GetIotInfo(args NickForm) (DeviceForm, error)

}