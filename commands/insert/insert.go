package insert

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdInsert creates the version command
func NewCmdInsert(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "insert",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
