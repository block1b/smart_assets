package main

import (
	"fmt"
	"github.com/goiot/libmqtt"
	"log"
	"time"
)

func main()  {
	fmt.Println()
	client, err := libmqtt.NewClient(
		// try MQTT 5.0 and fallback to MQTT 3.1.1
		libmqtt.WithVersion(libmqtt.V5, true),
		// server address(es) tcp://iot.eclipse.org:1883 ws://192.168.113.7:8083/mqtt
		libmqtt.WithServer("192.168.113.6:1883"),
		// enable keepalive (10s interval) with 20% tolerance
		libmqtt.WithKeepalive(10, 1.2),
		// enable auto reconnect and set backoff strategy
		libmqtt.WithAutoReconnect(true),
		libmqtt.WithBackoffStrategy(time.Second, 5*time.Second, 1.2),
		// use RegexRouter for topic routing if not specified
		// will use TextRouter, which will match full text
		libmqtt.WithRouter(libmqtt.NewRegexRouter()),
	)

	if err != nil {
		// handle client creation error
		panic("hmm, how could it failed")
	}

	// register handlers
	{
		// register net handler
		client.HandleNet(func(server string, err error) {
			if err != nil {
				log.Printf("error happened to connection to server [%v]: %v", server, err)
			}
		})
		// register persist handler, you don't need this if all your message had QoS 0
		client.HandlePersist(func(err error) {
			if err != nil {
				log.Printf("session persist error: %v", err)
			}
		})
		// register subscribe handler
		client.HandleSub(func(topics []*libmqtt.Topic, err error) {
			if err != nil {
				for _, t := range topics {
					log.Printf("subscribe to topic [%v] failed: %v", t.Name, err)
				}
			} else {
				for _, t := range topics {
					log.Printf("subscribe to topic [%v] success: %v", t.Name, err)
				}

				// publish some packet (just for example)
				client.Publish([]*libmqtt.PublishPacket{
					{TopicName: "foo", Payload: []byte("bar"), Qos: libmqtt.Qos0},
					{TopicName: "bar", Payload: []byte("foo"), Qos: libmqtt.Qos1},
				}...)
			}
		})
		// register unsubscribe handler
		client.HandleUnSub(func(topic []string, err error) {
			if err != nil {
				// handle unsubscribe failure
				for _, t := range topic {
					log.Printf("unsubscribe to topic [%v] failed: %v", t, err)
				}
			} else {
				for _, t := range topic {
					log.Printf("unsubscribe to topic [%v] failed: %v", t, err)
				}
			}
		})
		// register publish handler
		client.HandlePub(func(topic string, err error) {
			if err != nil {
				log.Printf("publish packet to topic [%v] failed: %v", topic, err)
			} else {
				log.Printf("publish packet to topic [%v] success: %v", topic, err)
			}
		})

		// handle every subscribed message (just for example)
		client.Handle(".*", func(topic string, qos libmqtt.QosLevel, msg []byte) {
			log.Printf("[%v] message: %v", topic, string(msg))
		})
	}

	// connect to server
	client.Connect(func(server string, code byte, err error) {
		if err != nil {
			log.Printf("connect to server [%v] failed: %v", server, err)
			return
		}

		if code != libmqtt.CodeSuccess {
			log.Printf("connect to server [%v] failed with server code [%v]", server, code)
			return
		}

		// connected
		go func() {
			// subscribe to some topics
			client.Subscribe([]*libmqtt.Topic{
				{Name: "foo", Qos: libmqtt.Qos0},
				{Name: "bar", Qos: libmqtt.Qos1},
			}...)

			// in this example, we publish packets right after subscribe succeeded
			// see `client.HandleSub`
		}()
	})

	client.Wait()
}
