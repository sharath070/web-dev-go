package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	writeToConn(conn)
}

func readFromConn(conn net.Conn) {
	byteSlice, err := io.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(byteSlice))

	defer conn.Close()
}

func writeToConn(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintln(conn, "I dailed you...")
}
