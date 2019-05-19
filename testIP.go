package main

import (
	"fmt"
	"net"
)

func main() {


	ip := net.IP{127, 0, 0, 1}

	ipstr := "127.0.0.1"
	fmt.Println(ip.String() == ipstr)
	fmt.Println(net.ParseIP(ipstr), ip)
}

