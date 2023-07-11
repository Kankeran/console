package console

type CommandMessage struct {
	Name  string
	Flags []CommandFlagMessage
}

type CommandFlagMessage struct {
	Name   string
	Values []string
}
