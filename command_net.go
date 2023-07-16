package console

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func GetAdress() (address string) {
	address = os.Getenv("CMD_ADDRESS")
	if address == "" {
		address = "localhost:51005"
	}
	return
}

func ListenCommands() error {
	return NewCommandListener(GetAdress()).ListenCommands()
}

type CommandListener struct {
	Address string
}

func NewCommandListener(address string) *CommandListener {
	return &CommandListener{
		Address: address,
	}
}

func (c *CommandListener) ListenCommands() error {
	l, err := net.Listen("tcp", c.Address)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go c.handleConnection(conn)
	}
}

func (c *CommandListener) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Błąd odczytu danych:", err.Error())
		return
	}
	n := binary.BigEndian.Uint32(buffer)

	fmt.Println("Received: ", n)

	buffer = make([]byte, n)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Błąd odczytu danych:", err.Error())
		return
	}

	fmt.Println("Received: ", buffer)
	msg := MessageFromBytes(buffer)

	in := &FlagParser{flags: msg.Flags}

	commandInfoMap[msg.Name].ExecuteCallback(in, conn)

}
