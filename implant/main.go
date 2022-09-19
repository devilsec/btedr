package main

// go build -o btedr-implant -ldflags "-X \"main.ServerIP=$SERVERIP\" -X main.ServerPort=$SERVERPORT" ./implant

import (
	"strconv"

	"github.com/devilsec/btedr/implant/client"
)

// Strings embedded at compile time
var ServerIP string
var ServerPort string

func main() {
	port, err := parsePort(ServerPort)
	if err != nil {
		panic(err)
	}

	client, err := client.New(ServerIP, port)
	if err != nil {
		panic(err)
	}

	defer client.Dial.Close()
	client.Register()

	// TODO:
	// While not done...
	// fetch tasks
}

// Convert string to int16
func parsePort(port string) (int16, error) {
	parse, err := strconv.ParseInt(ServerPort, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(parse), nil
}
