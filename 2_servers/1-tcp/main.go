package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Test with telnet localhost 8080
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	defer conn.Close()
}


func client() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	_, err = fmt.Fprintln(conn, "I dialed you.")
	if err != nil {
		log.Panic(err)
	}

}