package help

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootHelpFunc - function to see help correctly
func RootHelpFunc(command *cobra.Command, args []string) {
	for _, c := range command.Commands() {
		fmt.Println(c.Use)
	}
}
