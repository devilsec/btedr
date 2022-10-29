package main

// go build -o btedr-agent -ldflags "-X \"main.ServerIP=$SERVERIP\" -X main.ServerPort=$SERVERPORT" ./agent

import (
	"fmt"
	"strconv"
	"time"

	"github.com/devilsec/btedr/agent/client"
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

	// Keep trying to register if registration fails
	for err := client.Register(); err != nil; err = client.Register() {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
	}

	// TODO:
	// While not done...
	// fetch tasks
}

// Convert string to int16
func parsePort(port string) (uint16, error) {
	parse, err := strconv.ParseUint(ServerPort, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(parse), nil
}
