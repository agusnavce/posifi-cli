package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootHelpFunc - function to see help correctly
func rootHelpFunc(command *cobra.Command, args []string) {
	for _, c := range command.Commands() {
		fmt.Println(c.Use)
	}
}
