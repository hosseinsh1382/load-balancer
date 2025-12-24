package server

import (
	"net"
	url2 "net/url"
)

type DefaultHealthChecker struct {
}

func NewDefaultHealthChecker() *DefaultHealthChecker {
	return &DefaultHealthChecker{}
}

func (hc *DefaultHealthChecker) CheckHealth(url string) bool {
	host, err := url2.Parse(url)
	conn, err := net.Dial("tcp", host.Host)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
