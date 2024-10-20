# AiPT (AI-Powered Terminal)

AiPT is a powerful command-line tool that executes a series of commands defined in a JSON file. It supports both local and remote command sources, making it flexible for various use cases.

## Features

- Execute commands from a local JSON file or a remote URL
- Cross-platform support (Windows and Unix-based systems)
- Built-in commands for shell execution, file operations, process management, and system information
- Recursive and non-recursive directory listing
- Compact binary size with optional UPX compression

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/Karib0u/AiPT.git
   cd AiPT
   ```

2. Build the project:
   ```
   go build -ldflags="-s -w" -o AiPT
   ```

   This command will create an optimized binary with reduced size.

3. (Optional) Further compress the binary using UPX:
   - Download and install UPX from [https://upx.github.io/](https://upx.github.io/)
   - Run the following command:
     ```
     upx --best --lzma AiPT
     ```

## Usage

Run AiPT with a command source:

```
./AiPT [command_source]
```

If no command source is provided, it will use the default remote URL: `https://example.com/commands_to_execute.json`

To use a local test file:

```
./AiPT test
```

This will use the local file `commands_to_execute.json` instead of the remote URL.

## Available Commands

AiPT supports the following built-in commands:

1. `shell`: Execute a shell command
   - Usage: `shell [command]`

2. `file`: Perform file operations
   - Usage: `file [action] [path] [content]`
   - Actions: create, update, read, delete, list, tree

3. `process`: Manage processes
   - Usage: `process [action] [args]`
   - Actions: list, create, kill

4. `sysinfo`: Display system information
   - Usage: `sysinfo`

## Command Source Format

The command source should be a JSON file containing an array of objects with the following structure:

```json
[
  {
    "name": "command_name",
    "params": ["param1", "param2", ...]
  },
  ...
]
```

## Dependencies

AiPT uses the following external libraries:

- `github.com/shirou/gopsutil/v3`: For system and process information

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]
