package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// sudo docker ps -aq | xargs sudo docker stop
	Command_output("sudo docker ps -aq | xargs sudo docker stop")
	// sudo docker ps -aq | xargs sudo docker rm
	Command_output("sudo docker ps -aq | xargs sudo docker rm")
	// sudo docker network ls -q | xargs sudo docker network rm
	Command_output("sudo docker network ls -q | xargs sudo docker network rm")
	// sudo service docker stop
	Command_output("sudo service docker stop")
}

func Command_output(cmd string) {
	fmt.Printf("cmd: %s \n", cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("output: %s \n", output)
}
