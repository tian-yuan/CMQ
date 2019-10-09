package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://localhost:8972")
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
}
