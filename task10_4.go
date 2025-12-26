package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", ":9443", cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("TLS server :9443")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			msg, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("client: ", msg)
			c.Write([]byte("SECURE OK\n"))
		}(conn)
	}
}
