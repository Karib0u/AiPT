package types

// Command represents a command that can be executed
type Command struct {
	Name        string
	Description string
	Execute     func(...string) error
}

// CommandToExecute represents a command to be executed
type CommandToExecute struct {
	Name   string   `json:"name"`
	Params []string `json:"params"`
}
