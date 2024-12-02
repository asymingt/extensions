package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
    {
		var stdout bytes.Buffer
		cmd := exec.Command("nvidia-smi", "-i", "0", "-c", "EXCLUSIVE_PROCESS")
		cmd.Stdout = &stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error: %v\n", stdout.String())
		} else {
			fmt.Printf("success: %v\n", stdout.String())
		}
	}
	{
	    var stdout bytes.Buffer
		cmd := exec.Command("nvidia-cuda-mps-control", "-d")
		cmd.Stdout = &stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error: %v\n", stdout.String())
		} else {
			fmt.Printf("success: %v\n", stdout.String())
		}
	}
}