package utils

import (
	"os/exec"
	"fmt"
    "strings"
)

func ExecuteCommand(app string, args ...string) string {
    // Create a new command using the variadic args
    cmd := exec.Command(app, args...)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err.Error()
    }

    // Print the output
    return string(stdout)
}


func CheckOS() string {
    output := ExecuteCommand("cat","/etc/os-release")
    temp := strings.Split(output,"\n")
    var osName string
    for _, line := range temp {
        if strings.HasPrefix(line, "NAME=") {
            osName = strings.TrimPrefix(line, "NAME=")
            osName = strings.Trim(osName, `"`) // Remove the double quotes
            break
        }
    }

    return osName
}