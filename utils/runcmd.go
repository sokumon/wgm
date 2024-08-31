package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(app string, args ...string) {
	// Create a new command using the variadic args
	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
