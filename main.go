package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 2 {
		color.Red("Usage: go run command.go <command> [args...]")
		return
	}

	// Get the command and its arguments
	command := os.Args[1]
	args := os.Args[2:]

	// Build the command
	start := time.Now()
	cmd := exec.Command(command, args...)

	// Redirect the output to the Go program's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	color.Cyan("Executing command: %s %s", command, strings.Join(args, " "))
	err := cmd.Start()
	if err != nil {
		color.Red("Error starting command: %v", err)
		return
	}

	// Display process details
	displayProcessDetails(cmd)

	// Start a goroutine to monitor process details
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go monitorProcess(cmd, ticker)

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		color.Red("Command terminated with error: %v", err)
	} else {
		color.Green("Command executed successfully in %d ms", time.Since(start).Milliseconds())
	}
}

// Display process details for the given command
func displayProcessDetails(cmd *exec.Cmd) {
	psCmd := exec.Command("ps", "-p", fmt.Sprintf("%d", cmd.Process.Pid), "-o", "pid,ppid,cmd,%mem,%cpu,etime")
	psOutput, err := psCmd.CombinedOutput()
	if err != nil {
		color.Red("Error getting process details: %v", err)
		return
	}
	color.Yellow("Process details for PID %d:\n%s", cmd.Process.Pid, string(psOutput))
}

// Monitor the process periodically and display details
func monitorProcess(cmd *exec.Cmd, ticker *time.Ticker) {
	for range ticker.C {
		psCmd := exec.Command("ps", "-p", fmt.Sprintf("%d", cmd.Process.Pid), "-o", "pid,ppid,cmd,%mem,%cpu,etime")
		psOutput, err := psCmd.CombinedOutput()
		if err != nil {
			color.Red("Error getting process details: %v", err)
			return
		}
		color.Yellow("Process details for PID %d:\n%s", cmd.Process.Pid, string(psOutput))
	}
}
