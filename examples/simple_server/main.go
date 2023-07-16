package main

import (
	"fmt"

	"github.com/Kankeran/console"
)

func main() {
	var cmdInfo = console.RegisterCommand("something", "Test command", OnExec)
	cmdInfo.OptionalInt("asd", "Getting int value", 123)
	fmt.Println(console.ListenCommands())
}

func OnExec(in console.Input, out console.Output) error {
	fmt.Println("Called")

	fmt.Fprint(out, "Hello")
	return nil
}
