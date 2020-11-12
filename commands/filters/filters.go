package filters

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdFilters creates the version command
func NewCmdFilters(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "filters",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
