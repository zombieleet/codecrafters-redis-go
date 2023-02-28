package main

import (
	"bufio"
	"fmt"
	_ "log"
	"net"
	"os"
)

func main() {
	fmt.Println("logs for server.")
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		_, _ = reader.ReadString('\n')
		writer.WriteString("+PONG\r\n")
		writer.Flush()
	}

}
