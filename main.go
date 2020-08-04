package main

import (
	"fmt"

	"github.com/Ashu-Soni/cli/commandLine"
)

func main() {
	fmt.Println("This is my first command line application")
	fmt.Println("Author: Ashutosh Soni")
	fmt.Println("repo: github.com/Ashu-Soni/cli")

	err := commandLine.CreateCommandLine()
	if err != nil {
		fmt.Println("error got : ", err)
	}
}
