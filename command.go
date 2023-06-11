package console

type CommandFlag struct {
	Name         string
	DefaultValue string
	Description  string
}

type Command struct {
	Name        string
	Description string
	Flags       []CommandFlag
}

func RegisterCommand(command Command) {

}
