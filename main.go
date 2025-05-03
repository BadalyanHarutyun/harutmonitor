package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	var monitorLog string
	flag.StringVar(&monitorLog, "monitor-log", "", "Log file to write process monitoring data")
	flag.Parse()

	// Get command and args
	args := flag.Args()
	if len(args) < 1 {
		color.Red("Usage: harutmonitor [--monitor-log=file] <command> [args...]")
		return
	}

	command := args[0]
	cmdArgs := args[1:]

	// Start the command
	start := time.Now()
	cmd := exec.Command(command, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	color.Cyan("Executing command: %s %s", command, strings.Join(cmdArgs, " "))

	err := cmd.Start()
	if err != nil {
		color.Red("Error starting command: %v", err)
		return
	}

	proc, err := process.NewProcess(int32(cmd.Process.Pid))
	if err != nil {
		color.Red("Failed to get process info: %v", err)
		return
	}

	displayProcessDetails(proc)

	// Monitor in a separate goroutine
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	if monitorLog != "" {
		go monitorToFile(proc, ticker, done, monitorLog)
	} else {
		go monitorToConsole(proc, ticker, done)
	}

	err = cmd.Wait()
	close(done) // stop the monitoring goroutine

	if err != nil {
		color.Red("Command terminated with error: %v", err)
	} else {
		color.Green("Command executed successfully in %d ms", time.Since(start).Milliseconds())
	}

	// Final display of process details
	displayProcessDetails(proc)
}

// Display process details (one-time)
func displayProcessDetails(p *process.Process) {
	memPercent, _ := p.MemoryPercent()
	cpuPercent, _ := p.CPUPercent()
	cmdline, _ := p.Cmdline()
	createTime, _ := p.CreateTime()
	elapsed := time.Since(time.UnixMilli(createTime))

	color.Yellow("Process Details:\nPID: %d\nCMD: %s\nCPU: %.2f%%\nMemory: %.2f%%\nElapsed: %s\n",
		p.Pid, cmdline, cpuPercent, memPercent, elapsed.Truncate(time.Second))
}

// Monitor process and print to console
func monitorToConsole(p *process.Process, ticker *time.Ticker, done <-chan struct{}) {
	for {
		select {
		case <-ticker.C:
			displayProcessDetails(p)
		case <-done:
			return
		}
	}
}

// Monitor process and write to a log file
func monitorToFile(p *process.Process, ticker *time.Ticker, done <-chan struct{}, logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		color.Red("Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	for {
		select {
		case <-ticker.C:
			memPercent, _ := p.MemoryPercent()
			cpuPercent, _ := p.CPUPercent()
			cmdline, _ := p.Cmdline()
			createTime, _ := p.CreateTime()
			elapsed := time.Since(time.UnixMilli(createTime))

			logLine := fmt.Sprintf("%s | PID: %d | CMD: %s | CPU: %.2f%% | Mem: %.2f%% | Elapsed: %s\n",
				time.Now().Format(time.RFC3339), p.Pid, cmdline, cpuPercent, memPercent, elapsed.Truncate(time.Second))
			file.WriteString(logLine)
		case <-done:
			return
		}
	}
}
