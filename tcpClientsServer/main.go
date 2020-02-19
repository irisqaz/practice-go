package main

import (
	"fmt"
	"net"
)

const addr = ":8080"

func main() {
	go server(addr)
	go client(addr)
	go client(addr)

	fmt.Scanln()
}

func server(addr string) {
	listener, _ := net.Listen("tcp", addr)
	defer listener.Close()
	fmt.Println("server: listening on port", addr)
	for {
		conn, _ := listener.Accept()
		fmt.Println("server: got connection ...")

		go handle(conn)

		//conn.Close()
	}
}

func client(addr string) {
	fmt.Println("client: starting ...")
	conn, _ := net.Dial("tcp", addr)
	defer conn.Close()

	msg := "Hello there\n Bye"
	conn.Write([]byte(msg))
	fmt.Println("client: send message:", msg)

	resp := make([]byte, 256)
	n, _ := conn.Read(resp)
	fmt.Println("client: got:", string(resp[:n]))
	fmt.Println("client: exiting ...")
}

func handle(conn net.Conn) {
	//io.Copy(conn, conn)
	req := make([]byte, 256)
	n, _ := conn.Read(req)
	conn.Write(req[:n])

	conn.Close()
}
