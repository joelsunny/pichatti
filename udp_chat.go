package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	go uchatClient()
	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 10001, Zone: ""})
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, _ := ServerConn.ReadFromUDP(buf)
		fmt.Println(addr.IP, ": ", string(buf[0:n]))
	}
}

func uchatClient() {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{255, 255, 255, 255}, Port: 10001, Zone: ""})
	defer Conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		Conn.Write([]byte(msg))
	}
}
