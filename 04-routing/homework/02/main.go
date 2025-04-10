package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleHttp(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}

	/*
		Code will never reach here cuz, the scan above will be open since the
		connection to browser (http) in and stream... even if the req data is sent
		it will still read the open stream.
		the value of ln = "" (after the request data is sent due to open stream)
		we need to explicitly break the loop if we want to write somthing...
	*/
	fmt.Println("Code go here")
	io.WriteString(conn, "I see you are connected")
}

func handleHttp(conn net.Conn) {
	defer conn.Close()

	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}

	fmt.Println("Code go here")

	// fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	// fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	// fmt.Fprintf(conn, "\r\n")

	// still write in browser (no need of above stuff.. ig we need that only for text/html)
	io.WriteString(conn, "I see you are connected")
}
