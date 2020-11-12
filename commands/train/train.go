package train

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/spf13/cobra"
)

// NewCmdTrain creates the version command
func NewCmdTrain(f *pkg.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "train",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test")
		},
	}

	return cmd
}
