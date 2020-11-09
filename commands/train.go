package commands

import (
	"fmt"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

// Options - the options for training
type Options struct {
	IO         *iostreams.IOStreams
	HTTPClient func() (pkg.Client, error)
}

// NewCmdTrain creates the version command
func NewCmdTrain(f *pkg.Factory) *cobra.Command {
	tOpts := &Options{
		IO:         f.IOStreams,
		HTTPClient: f.HTTPClient,
	}

	cmd := &cobra.Command{
		Use: "train",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("test")
			return trainRun(tOpts)
		},
	}

	return cmd
}

func trainRun(opts *Options) error {
	client, err := opts.HTTPClient()
	if err != nil {
		return err
	}
	fmt.Println(client)
	return nil
}
