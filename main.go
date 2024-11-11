package main

import (
	"flag"
	"MEM-CACHED-SERVER/server_methods"
)

func main() {
	flag.Parse()
	server_methods.Start_server()
}
