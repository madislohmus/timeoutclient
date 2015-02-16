/*
 * -----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <margus@stfu.ee> wrote this file.  As long as you retain this notice you can
 * do whatever you want with this stuff. If we meet some day, and you think this
 * stuff is worth it, you can buy me a beer in return. 
 * Original license by Poul-Henning Kamp.
 * -----------------------------------------------------------------------------
 */

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

// NewTimeoutClient returns a http.Client, that has a connectTimeout and readWriteTimeout (any time.Duration will do)
func NewTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(connectTimeout, readWriteTimeout),
		},
	}
}
