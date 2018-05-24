package main

import (
	"os/exec"
	"fmt"
)

func main() {
	proc := exec.Command("sleep", "1")
	proc.Start()
	proc.Wait()
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)

	fmt.Printf("PID: %d\n\n", proc.Process.Pid)
	fmt.Printf("Process took: %dμs\n", proc.ProcessState.SystemTime())
	fmt.Printf("Exit successfully: %t\n", proc.ProcessState.Success())
}
