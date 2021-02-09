package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("docker ps -aq | xargs sudo docker stop")
	// sudo docker ps -aq | xargs sudo docker stop
	dockerstop := "sudo docker ps -aq | xargs sudo docker stop"
	outstop, err := exec.Command("sh", "-c", dockerstop).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("output: %s", outstop)
	// sudo docker ps -aq | xargs sudo docker rm
	dockerrm := "sudo docker ps -aq | xargs sudo docker rm"
	outrm, err := exec.Command("sh", "-c", dockerrm).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("output: %s", outrm)
	// sudo docker network ls -q | xargs sudo docker network rm
	dockernr := "sudo docker network ls -q | xargs sudo docker network rm"
	outnr, err := exec.Command("sh", "-c", dockernr).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("output: %s", outnr)
	// sudo service docker stop
	dockerservicestop := "sudo service docker stop"
	outservicestop, err := exec.Command("sh", "-c", dockerservicestop).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("output: %s", outservicestop)
}
