package wol

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var (
	bemfaBroker   string
	bemfaPort     int
	bemfaTopic    string
	bemfaClientID string
)

func MqttHandle(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	fmt.Printf("Received message: %s from topic: %s\n", payload, msg.Topic())
	if payload == "on" {
		_ = Shutdown()
	} else if payload == "off" {
		_ = Shutdown()
	}
}

func ConnectBemfa() mqtt.Client {
	//opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", "bemfa.com", 9501)).SetClientID("8dfc34c2dffe4948b031b724e7192e19")
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", "xxxxxx", 9501)).SetClientID("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	opts.SetDefaultPublishHandler(MqttHandle)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("xxxxxxx", 1, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	return client
}
