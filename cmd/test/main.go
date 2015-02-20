package main

import (
	"fmt"
	"log"
	"net"

	"github.com/docker/libnetwork"
	_ "github.com/docker/libnetwork/bridge"
)

func main() {
	_, net, _ := net.ParseCIDR("192.168.100.1/24")

	options := libnetwork.DriverParams{"Subnet": *net}
	netw, err := libnetwork.NewNetwork("bridgednetwork", options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Network=%#v\n", netw)
}