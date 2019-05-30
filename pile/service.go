package pile

import (
	"fmt"
	. "smart_assets/tool"

	MQTT  "github.com/eclipse/paho.mqtt.golang"
)


// 提供的服务
// 类似于api的定义
func InitPileApi() error {
	var err error
	// 开锁
	err = Sub("clock0/openClock", OpenClockPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}
	// 关锁
	err = Sub("clock0/closeClock", CloseClockPubHandler)
	if err != nil{
		fmt.Println("build api :", err)
		return err
	}

	return nil
}

var OpenClockPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[OpenClockPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
}

var CloseClockPubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[CloseClockPubHandler] TOPIC: %s MSG: %s\n", msg.Topic(), msg.Payload())
}
