package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
	What is ROT 13 encryption?
	It is a simple encryption algo that takes only english alphabets and
	rotates it to next 13 letters

	ROT-N shifts each letter N places forward in the alphabet.
	If it goes past Z, it wraps around to A.

	So in ROT13:
		A → N
    	B → O
    	...
    	Z → M
*/

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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		bs := []byte(line)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s -> %s\n\n", line, r)
	}
}

func rot13(bs []byte) string {
	for i, c := range bs {
		switch {
		case c >= 'A' && c <= 'Z':
			bs[i] = 'A' + (c-'A'+13)%26
		case c >= 'a' && c <= 'z':
			bs[i] = 'a' + (c-'a'+13)%26
		}
	}
	return string(bs)
}
