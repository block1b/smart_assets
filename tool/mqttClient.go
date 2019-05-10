package tool

import MQTT  "github.com/eclipse/paho.mqtt.golang"

// 全局mqtt client

var MqttClient MQTT.Client

func Init()  {
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.18.128:1883")
	opts.SetClientID("smartServer")
	opts.SetDefaultPublishHandler(DefaultPublishHandler)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	MqttClient = c
}

// publish
// Subscribe

// SubPub todo

