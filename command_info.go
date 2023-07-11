package console

type CommandInfo struct {
	Name            string
	Description     string
	FlagsInfo       []FlagInfo
	ExecuteCallback func(FlagParser) error
}

type FlagInfo struct {
	Name        string
	IsRequired  bool
	Description string
	ValueType   string
}

var (
	CommandInfoMap = make(map[string]CommandInfo)
)

func RegisterCommand(commandInfo CommandInfo) {
	CommandInfoMap[commandInfo.Name] = commandInfo
}
