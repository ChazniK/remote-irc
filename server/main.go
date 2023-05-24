package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go connectionHandler(conn)
	}
}

func connectionHandler(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
}
