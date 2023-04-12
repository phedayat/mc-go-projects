package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main(){
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("[ERROR] [Listening] ", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("[ERROR] [Accepting] ", err.Error())
			os.Exit(1)
		}
		go processClient(conn)
	}
}

func processClient(conn net.Conn){
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("[ERROR] [Reading] ", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:n]))
	_, err = conn.Write([]byte("Message received!"))
	conn.Close()
}

