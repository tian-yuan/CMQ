package main

import (
	"fmt"
	"CMQ/hub/router"
)

func main() {
	trie := router.NewTrie()
	//trie.Parse("/productKey/deviceName/*")
	trie.Parse("/topic/aaa").Handle("GET", "node1,node2")
	trie.Parse("/topic/*").Handle("GET", "node3,node4")
	var matched *router.Matched
	var err error
	matched, err = trie.Match("/topic/aaa")
	if err != nil {
		fmt.Printf("err %s\n", err)
	}

	fmt.Printf("matched path : %s\n", matched.Path)
	if matched.Node == nil {
		fmt.Printf("matched path : %s\n", matched.Node)
	}
	fmt.Printf(" handler route : %s\n", matched.Node.GetHandler("GET").(string))

	matched, err = trie.Match("/topic/bbb")
	fmt.Printf(" handler route : %s\n", matched.Node.GetHandler("GET").(string))
	matched, err = trie.Match("/topic/bbb")
	fmt.Printf(" handler route : %s\n", matched.Node.GetHandler("GET").(string))
	fmt.Println("finished.")
}
