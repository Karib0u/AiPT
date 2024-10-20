package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"AiPT/types"
)

// ExecuteCommand runs the appropriate command based on the OS
func ExecuteCommand(windowsCmd, unixCmd string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", windowsCmd)
	} else {
		cmd = exec.Command("sh", "-c", unixCmd)
	}
	out, err := cmd.Output()
	return string(out), err
}

// FetchCommandsToExecute retrieves commands from a remote URL or local file
func FetchCommandsToExecute(source string) ([]types.CommandToExecute, error) {
	var data []byte
	var err error

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
	} else {
		data, err = os.ReadFile(source)
	}

	if err != nil {
		return nil, err
	}

	var commands []types.CommandToExecute
	err = json.Unmarshal(data, &commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

// ListDirectory lists the contents of a directory
func ListDirectory(path string, recursive bool) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !recursive && filePath != path && info.IsDir() {
			return filepath.SkipDir
		}
		relPath, err := filepath.Rel(path, filePath)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil
		}
		indent := strings.Repeat("  ", strings.Count(relPath, string(os.PathSeparator)))
		fileType := "F"
		if info.IsDir() {
			fileType = "D"
		}
		fmt.Printf("%s[%s] %s (%d bytes)\n", indent, fileType, info.Name(), info.Size())
		return nil
	})
}
