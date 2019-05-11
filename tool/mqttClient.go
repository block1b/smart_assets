package tool

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 全局mqtt client

var MqttClient MQTT.Client

const CLIENTID  = "smartServer"

func Init()  {
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.1.107:1883")
	opts.SetClientID(CLIENTID)
	opts.SetDefaultPublishHandler(DefaultPublishHandler)
	opts.SetOnConnectHandler(OnConnectHandler)
	opts.SetConnectionLostHandler(ConnectionLostHandler)
	opts.SetAutoReconnect(true)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	MqttClient = c
}

func GetClient() MQTT.Client {
	return MqttClient
}

// publish
func Pub(pubTopic string, payload []byte) error {
	c := GetClient()
	pubToken := c.Publish(pubTopic, 0, false, payload)
	pubToken.Wait()
	return nil
}
// Subscribe !!! 订阅 CLIENTID+"xxx" = smartServer/#
func Sub(subTopic string, callback MQTT.MessageHandler) error {
	c := GetClient()
	if token := c.Subscribe(CLIENTID+subTopic, 0, callback); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return token.Error()
	}
	return nil
}

// SubPub
func SubPub(subTopic string, callback MQTT.MessageHandler, pubTopic string, payload []byte) error {
	var err error
	err = Sub(subTopic, callback)
	err = Pub(pubTopic, payload)
	if err != nil {
		return err
	}
	return nil
}

