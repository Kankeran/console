package console

import (
	"net"
)

type CommandListener struct {
	Host string
	Port string
}

func NewCommandListener(host string, port string) *CommandListener {
	return &CommandListener{
		Host: host,
		Port: port,
	}
}

func (c *CommandListener) ListenCommands() error {
	l, err := net.Listen("tcp", c.Host+":"+c.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	// TODO acceptTCP

	return nil
}

func ListenCommands() error {
	return NewCommandListener("localhost", "51005").ListenCommands()
}
