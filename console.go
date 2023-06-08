package console

type CommandOptionMode uint32

const (
	ValueNone = 1 << iota
	ValueRequired
	ValueOptional
	ValueIsArray
)

type CommandOption struct {
	Name        string
	Shortcut    string
	Mode        CommandOptionMode
	Description string
}

type Command struct {
	Name        string
	Description string
	Options     []CommandOption
}

func RegisterCommand(command Command) {

}
