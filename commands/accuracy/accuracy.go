package accuracy

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdAccuracy creates the version command
func NewCmdAccuracy(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "accuracy",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
