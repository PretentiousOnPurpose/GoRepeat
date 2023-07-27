package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			// log.Fatalln(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	defer conn.Close()

	for scanner.Scan() {
		data := scanner.Text()

		if data == "exit" {
			return
		}

		fmt.Println(data)
		fmt.Fprintf(conn, "You say: %s\n", data)
	}
}
