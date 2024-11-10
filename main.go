package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

type Datablock struct {
	key        string
	data_block string
	flags      uint16
	expiry     int
	byte_count int
}

var port *string = flag.String("port", "11211", "Server will be running on the given port if specified defaults to 1121")
var ip string = "localhost"
var cache map[string]Datablock = make(map[string]Datablock)

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
			go handle_connection(Connection)
		}
	}
}

func handle_connection(Connection net.Conn) {
	var (
		command string
		block   string
		error   error
	)

	reader := bufio.NewReader(Connection)
	writer := bufio.NewWriter(Connection)

	for {
		command, error = reader.ReadString('\n')
		command = strings.TrimRight(command, "\r\n")

		if !strings.HasPrefix(command, "get") {
			block, error = reader.ReadString('\n')
			block = strings.TrimRight(block, "\r\n")
		}

		if error != nil {
			fmt.Println("Error:", error)
			return
		} else {
			result := parseCommand(command, block)
			writer.WriteString(result)
			writer.Flush()
		}
	}
}

func parseCommand(command string, data_block string) string {
	var (
		block  Datablock
		result string
	)

	block.data_block = data_block
	name, args, _ := strings.Cut(command, " ")

	if name == "set" {
		fmt.Sscanf(args, "%v %v %v %v %v", &block.key, &block.flags, &block.expiry, &block.byte_count)
		result = putBlock(block)
	} else {
		key, _, _ := strings.Cut(args, " ")
		result = getBlock(key)
	}

	return result
}

func getBlock(key string) string {
	value, present := cache[key]

	if present {
		return fmt.Sprintf("VALUE %v %v %v\r\n%v\r\nEND\r\n", value.key, value.flags, value.byte_count, value.data_block)
	} else {
		return "END\r\n"
	}
}

func putBlock(block Datablock) string {
	cache[block.key] = block
	return "STORED\r\n"
}

func main() {
	flag.Parse()
	start_server()
}
