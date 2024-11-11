package server_methods

import (
	"MEM-CACHED-SERVER/commands"
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

var port *string = flag.String("port", "11211", "Server will be running on the given port if specified defaults to 1121")
var ip string = "localhost"

func Start_server() {
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
			go handle_connection(Connection)
		}
	}
}

func handle_connection(Connection net.Conn) {

	defer Connection.Close()

	var (
		command string
		err     error
	)

	reader := bufio.NewReader(Connection)
	writer := bufio.NewWriter(Connection)

	readFromClient := func() (string, error) {
		return readLine(reader)
	}

	for {
		command, err = readFromClient()

		if err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			result := commands.ParseCommand(command, readFromClient)
			writer.WriteString(result)
			writer.Flush()
		}
	}
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n") + " $"
	return line, err
}
