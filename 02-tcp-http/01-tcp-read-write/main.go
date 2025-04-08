package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleReadWriteConn(conn)
	}
}

func writeToConn(conn net.Conn) {
	io.WriteString(conn, "Hello from TCP server!!!")
	fmt.Fprintln(conn, "Hey how are you...?")

	defer conn.Close()
}

/**************************************/
/******* Terminal read write *********/
/*************************************/

func handleReadWriteConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you said: %s\n\n", line)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}

/*********************************/
/******* Web Connection *********/
/*********************************/

func handleWebConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}

//  The `fmt.Println(line)` from handleConn gives this output

//  GET / HTTP/1.1
//  Host: localhost:8080
//  User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:137.0) Gecko/20100101 Firefox/137.0
//  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
//  Accept-Language: en-US,en;q=0.5
//  Accept-Encoding: gzip, deflate, br, zstd
//  DNT: 1
//  Sec-GPC: 1
//  Connection: keep-alive
//  Upgrade-Insecure-Requests: 1
//  Sec-Fetch-Dest: document
//  Sec-Fetch-Mode: navigate
//  Sec-Fetch-Site: cross-site
//  Priority: u=0, i
