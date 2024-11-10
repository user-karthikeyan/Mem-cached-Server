package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"net"
)

var port *string = flag.String("port", "11211", "Server will be running on the given port if specified defaults to 1121")
var ip string = "localhost"

func start_server() {
	address := ip + ":" + *port
	Listener, Error := net.Listen("tcp", address)

	if Error != nil {
		fmt.Println("Error:", Error)
		return
	} else {
		fmt.Println("Starting server at port", *port)
		service_client(Listener)
	}
}

func service_client(Listener net.Listener) {
	for {
		Connection, Error := Listener.Accept()

		if Error != nil {
			fmt.Println("Error:", Error)
			continue
		} else {
			fmt.Println("Connected to client")
			handle_connection(Connection)
		}
	}
}

func handle_connection(Connection net.Conn) {
	for {
		var command string

		error := gob.NewDecoder(Connection).Decode(&command)

		if error != nil {
			fmt.Println("Error:", error)
			return
		} else {
			fmt.Println(command)
			parseCommand(command)
		}
	}
}

func parseCommand(command string) {

	var (
		name       string
		key        string
		flags      int
		expiry     int
		byte_count int
		data_block string
	)

	fmt.Sscanf(command, "%s %s %d %d %d \r\n%s\r\n", name, key, flags, expiry, byte_count, data_block)
	fmt.Println(name, key, flags, expiry, byte_count, data_block)
}

func main() {
	flag.Parse()
	start_server()
}
