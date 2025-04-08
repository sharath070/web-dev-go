package main

import (
	"bufio"
	"fmt"
	"log"
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
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)

		if i == 0 {
			// request line (shows what kind of request is made)
			m := strings.Fields(ln)[0]   // method
			url := strings.Fields(ln)[2] // url
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	fmt.Printf("ends here..............\n\n\n\n")
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body><strong>Hello World</strong></body></html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
