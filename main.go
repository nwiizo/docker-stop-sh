package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed command.txt
var static embed.FS

func main() {
	// sudo docker ps -aq | xargs sudo docker stop
	// sudo docker ps -aq | xargs sudo docker rm
	// sudo docker network ls -q | xargs sudo docker network rm
	// sudo service docker stop
	b, err := static.ReadFile("command.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(b))
	Command_output(string(b))
}

func Command_output(cmds string) {
	arr := strings.Split(cmds, "\n")
	for _, cmd := range arr {
		fmt.Printf("cmd: %s \n", cmd)
		output, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("output: %s \n", output)
	}
}
