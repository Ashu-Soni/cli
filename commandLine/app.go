package commandLine

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

var CUSTOM_USAGE_TEXT = `DESCRIPTION:
   {{.Name}} - {{.UsageText}}

   USAGE:
   	{{.Usage}}

   COMMANDS:
		init				Internally call "Terraform init"
   		plan             	Internally call "Terraform plan"
   		apply            	Internally call "Terraform apply"
	   
	AUTHOR(S):
		Ashutosh Soni
`

func CreateCommandLine(version string, writer io.Writer, errWriter io.Writer) *cli.App {
	cli.AppHelpTemplate = CUSTOM_USAGE_TEXT

	app := cli.NewApp()

	app.Name = "cli"
	app.Author = "Ashutosh Soni"
	app.Version = version
	app.Action = runApp
	app.Writer = writer
	app.ErrWriter = errWriter
	app.UsageText = "To check how the development is done for terragrunt"

	return app
}

func runApp(context *cli.Context) (err error) {
	if !context.Args().Present() {
		return cli.ShowAppHelp(context)
	}

	givenCommand := context.Args().First()

	fmt.Println("You applied for the command : ", givenCommand)

	return runCommand(givenCommand)
}

func runCommand(command string) error {
	// comm := fmt.Sprintf("terraform %s", command)

	cmd := exec.Command("terraform", command, "--auto-approve")
	cmd.Stdin = strings.NewReader("some input")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())
	return nil
}
