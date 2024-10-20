package main

import (
	"fmt"
	"log"
	"os"

	"AiPT/commands"
	"AiPT/types"
	"AiPT/utils"
)

const (
	defaultCommandSource = "https://example.com/commands_to_execute.json"
	testCommandSource    = "commands_to_execute.json"
)

func main() {
	// Check if we're in test mode
	commandSource := defaultCommandSource
	if len(os.Args) > 1 && os.Args[1] == "test" {
		commandSource = testCommandSource
		os.Args = append(os.Args[:1], os.Args[2:]...) // Remove "test" from args
	}

	// Fetch commands to execute
	commandsToExecute, err := utils.FetchCommandsToExecute(commandSource)
	if err != nil {
		log.Fatalf("Error fetching commands to execute: %v", err)
	}

	// Execute the fetched commands
	for _, cmd := range commandsToExecute {
		if err := executeCommand(cmd); err != nil {
			log.Printf("Error executing command %s: %v\n", cmd.Name, err)
		}
	}
}

func executeCommand(cmd types.CommandToExecute) error {
	command, exists := commands.AvailableCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	fmt.Printf("Executing command: %s\n", cmd.Name)
	return command.Execute(cmd.Params...)
}
