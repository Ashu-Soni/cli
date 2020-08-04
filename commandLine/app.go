package commandLine

import "fmt"

func CreateCommandLine() error {
	if true {
		fmt.Println("into the command line mode")
		return nil
	}
	return fmt.Errorf("error catching the function")
}
