package main

import (
	"os/exec"
	"bytes"
	"fmt"
)

func main() {
	prc := exec.Command("ls", "-l")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if err := prc.Wait(); err != nil {
		panic(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
