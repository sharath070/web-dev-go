package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\n IN-MEMORY DATABASE\n\n"+
		"USE:\n"+
		"SET key value\n"+
		"GET key\n"+
		"DEL key\n\n\n")

	data := make(map[string]string)
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		line := scan.Text()
		fs := strings.Fields(line)

		switch fs[0] {
		case "GET":
			key := fs[1]
			val := data[key]
			fmt.Fprintln(conn, val)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "UNEXPECTED VALUE")
				continue
			}
			key := fs[1]
			val := fs[2]
			data[key] = val
		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}
}
