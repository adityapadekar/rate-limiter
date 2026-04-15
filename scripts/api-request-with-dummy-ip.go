package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

var routes = []string{"/route1"}
var sourceIPs = []string{"127.0.0.1"}

func clientWithIP(ip string) *http.Client {
	localAddr := &net.TCPAddr{IP: net.ParseIP(ip)}
	dialer := &net.Dialer{LocalAddr: localAddr}
	return &http.Client{
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}
}

func main() {
	var wg sync.WaitGroup
	for _, route := range routes {
		for _, ip := range sourceIPs {
			c := clientWithIP(ip)
			for i := 1; i <= 20; i++ {
				wg.Add(1)
				go func(route, ip string, n int) {
					defer wg.Done()
					resp, err := c.Get("http://127.0.0.1:8000" + route)
					if err != nil {
						fmt.Printf("[ERR] %s ip=%s req#%d: %v\n", route, ip, n, err)
						return
					}
					resp.Body.Close()
					fmt.Printf("[%d] %s ip=%s req#%d\n", resp.StatusCode, route, ip, n)
				}(route, ip, i)
			}
		}
	}
	wg.Wait()
}
