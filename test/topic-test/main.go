package main

import (
	"CMQ/hub/topic"
	"fmt"
)

func main() {
	m := topic.NewCSTrieMatcher()
	s0 := "productKey/deviceName/lamb/+"
	m.Subscribe(s0, "lamb")
	m.Subscribe("productKey/deviceName/ipc/+", "ipc")
	m.Subscribe("productKey/deviceName/+", "all")
	var sub []topic.Subscriber
	sub = m.Lookup("productKey/deviceName/ipc/temprature")
	fmt.Printf("match topic len : %d\n", len(sub))
	for key, value := range sub {
		fmt.Printf("topic key : %d\n", key)
		fmt.Printf("topic value : %s\n", value)
	}
	subs := topic.NewSubscription(1, "productKey/deviceName/ipc/+", "ipc")
	m.Unsubscribe(subs)
	sub = m.Lookup("productKey/deviceName/ipc/temprature")
	fmt.Printf("match topic len : %d\n", len(sub))
}
