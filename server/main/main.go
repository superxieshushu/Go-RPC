package main

import (
	"flag"
	"fmt"
	"net"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error Listening: ", err)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error Accept : ", err)
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleReq(conn)
	}
}

func handleReq(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [1024]byte
		size, _ := conn.Read(buf[0:])
		conn.Write(buf[0:size])
	}
}
