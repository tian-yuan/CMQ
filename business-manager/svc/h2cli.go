package svc

import (
	"crypto/tls"
	"net"
	"net/http"
	"golang.org/x/net/http2"
)

type HttpClientFactory struct {
	ClientMap map[string]*http.Client
}

var httpClientFactory HttpClientFactory

// key: "ip:port"
func GetH2client(key string) *http.Client {
	cli, exist := httpClientFactory.ClientMap[key]
	if !exist {
		cli = &http.Client{
			// Skip TLS dial
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}
		httpClientFactory.ClientMap[key] = cli
	}
	return cli
}

/*
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
 */
