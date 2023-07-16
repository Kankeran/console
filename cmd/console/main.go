package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Kankeran/console"
)

func main() {
	conn, err := net.Dial("tcp", console.GetAdress())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := console.CommandMessage{
		Name: "something",
		Flags: map[string][]string{
			"asd2": {"12"},
		},
	}

	requestBody := msg.ToBytes()
	request := make([]byte, 0, len(requestBody)+4)
	request = binary.BigEndian.AppendUint32(request, uint32(len(requestBody)))
	fmt.Println(request)
	request = append(request, requestBody...)

	fmt.Println(len(requestBody))
	conn.Write(request)

	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	b = b[:n]

	fmt.Println(string(b))
}
