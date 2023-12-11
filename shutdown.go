package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Program PID:", pid)

	err := syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Got to the end")
}
