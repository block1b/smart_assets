package tool

import "fmt"

import MQTT  "github.com/eclipse/paho.mqtt.golang"
// 模拟路由
//define a function for the default message handler
var DefaultPublishHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	// 手工撸路由
	switch msg.Topic() {
	case "smartServer/transaction/":  // 用于联调测试postServer
		fmt.Println("提交chain事务,weChat0")
		// todo 组装事务参数
		token := client.Publish("smartServer/transaction", 0, false, "")
		token.Wait()

	case "smartServer/balance":
		fmt.Println("余额相关 查，增，改")  // todo

	case "smartServer/iot":
		fmt.Println("设备相关 查，增，改")  // todo

	case "smartServer/bill":
		fmt.Println("账单相关 查")  // todo

	default:
		fmt.Println("undefined topic")
	}

}
