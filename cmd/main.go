package main

import (
	"fmt"
	"os"

	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/agusnavce/posifi-cli/commands/root"
	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/mgutz/ansi"
)

func main() {

	cmdFactory := pkg.NewFactory()
	//stderr := cmdFactory.IOStreams.ErrOut

	if !cmdFactory.IOStreams.ColorEnabled() {
		surveyCore.DisableColor = true
	} else {
		// override survey's poor choice of color
		surveyCore.TemplateFuncsWithColor["color"] = func(style string) string {
			switch style {
			case "white":
				if cmdFactory.IOStreams.ColorSupport256() {
					return fmt.Sprintf("\x1b[%d;5;%dm", 38, 242)
				}
				return ansi.ColorCode("default")
			default:
				return ansi.ColorCode(style)
			}
		}
	}

	rootCmd := root.NewCmdRoot(cmdFactory)
	stderr := cmdFactory.IOStreams.ErrOut
	cs := cmdFactory.IOStreams.ColorScheme()

	hasConfig := false
	if !hasConfig {
		fmt.Fprintln(stderr, cs.Bold("Welcome to Posifi CLI!"))
		fmt.Fprintln(stderr)
		fmt.Fprintln(stderr, "To authenticate, please run `gh auth login`.")
		fmt.Fprintln(stderr, "You can also set the GITHUB_TOKEN environment variable, if preferred.")
		os.Exit(4)
	}

	// c := credentials.NewSharedCredentials("", "default")
	// creds, err := c.Get()
	// if err != nil {
	// 	fmt.Fprintf(stderr, "Failed to read  AWS configuration:  %s\n", err)
	// 	os.Exit(2)
	// }
	//fmt.Println(creds)
	rootCmd.Execute()
}
