package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"AiPT/types"
	"AiPT/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

// AvailableCommands is a map of all available commands
var AvailableCommands = map[string]types.Command{
	"shell":   shellCommand,
	"file":    fileCommand,
	"process": processCommand,
	"sysinfo": sysinfoCommand,
}

var shellCommand = types.Command{
	Name:        "shell",
	Description: "Execute a shell command",
	Execute: func(args ...string) error {
		if len(args) == 0 {
			return fmt.Errorf("no command provided")
		}
		fmt.Println("Executing shell command...")
		out, err := utils.ExecuteCommand(strings.Join(args, " "), strings.Join(args, " "))
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	},
}

var fileCommand = types.Command{
	Name:        "file",
	Description: "Perform file operations",
	Execute: func(args ...string) error {
		if len(args) < 2 {
			return fmt.Errorf("insufficient arguments")
		}
		action := args[0]
		path := args[1]
		switch action {
		case "create", "update":
			if len(args) < 3 {
				return fmt.Errorf("content required for %s", action)
			}
			return os.WriteFile(path, []byte(args[2]), 0644)
		case "read":
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Println(string(data))
		case "delete":
			return os.Remove(path)
		case "list":
			return utils.ListDirectory(path, false)
		case "tree":
			return utils.ListDirectory(path, true)
		default:
			return fmt.Errorf("unknown file action: %s", action)
		}
		return nil
	},
}

var processCommand = types.Command{
	Name:        "process",
	Description: "Manage processes",
	Execute: func(args ...string) error {
		if len(args) == 0 {
			return fmt.Errorf("no action specified")
		}
		action := args[0]
		switch action {
		case "list":
			processes, err := process.Processes()
			if err != nil {
				return err
			}
			for _, p := range processes {
				name, _ := p.Name()
				fmt.Printf("PID: %d, Name: %s\n", p.Pid, name)
			}
		case "create":
			if len(args) < 2 {
				return fmt.Errorf("command required for create")
			}
			cmd := exec.Command(args[1], args[2:]...)
			return cmd.Start()
		case "kill":
			if len(args) < 2 {
				return fmt.Errorf("PID required for kill")
			}
			pid := args[1]
			_, err := utils.ExecuteCommand(fmt.Sprintf("taskkill /F /PID %s", pid), fmt.Sprintf("kill -9 %s", pid))
			return err
		default:
			return fmt.Errorf("unknown process action: %s", action)
		}
		return nil
	},
}

var sysinfoCommand = types.Command{
	Name:        "sysinfo",
	Description: "Display system information",
	Execute: func(...string) error {
		hostInfo, _ := host.Info()
		cpuInfo, _ := cpu.Info()
		memInfo, _ := mem.VirtualMemory()
		diskInfo, _ := disk.Usage("/")

		fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
		fmt.Printf("OS: %s %s\n", hostInfo.Platform, hostInfo.PlatformVersion)
		fmt.Printf("Kernel: %s\n", hostInfo.KernelVersion)
		fmt.Printf("CPU: %s (Cores: %d)\n", cpuInfo[0].ModelName, cpuInfo[0].Cores)
		fmt.Printf("Memory: %.2f GB / %.2f GB\n", float64(memInfo.Used)/1024/1024/1024, float64(memInfo.Total)/1024/1024/1024)
		fmt.Printf("Disk: %.2f GB / %.2f GB\n", float64(diskInfo.Used)/1024/1024/1024, float64(diskInfo.Total)/1024/1024/1024)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		return nil
	},
}
