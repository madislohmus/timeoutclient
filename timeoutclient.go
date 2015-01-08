package timeoutclient

import (
	"log"
	"net"
	"net/http"
	"time"
)

func timeoutDialer(connectTimeout time.Duration, readWriteTimeout time.Duration) func(network, address string) (c net.Conn, err error) {
	return func(network, address string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, address, connectTimeout)
		if err != nil {
			return nil, err
		}
		if err := conn.SetDeadline(time.Now().Add(readWriteTimeout)); err != nil {
			log.Println(err)
		}
		return conn, nil
	}
}

func NewTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(connectTimeout, readWriteTimeout),
		},
	}
}
