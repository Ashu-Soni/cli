package main

import (
	"fmt"
	"os"

	"github.com/Ashu-Soni/cli/commandLine"
)

var VERSION string

func main() {
	fmt.Print("This is my first command line application\n\n")
	fmt.Print("Author: Ashutosh Soni\n\n")
	fmt.Print("repo: github.com/Ashu-Soni/cli\n\n")

	app := commandLine.CreateCommandLine(VERSION, os.Stdout, os.Stderr)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error got : ", err)
	}
}
