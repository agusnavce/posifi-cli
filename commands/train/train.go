package train

import (
	"fmt"
	"net/http"

	"github.com/agusnavce/posifi-cli/pkg"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

// TrainOptions - the options for training
type TrainOptions struct {
	IO         *iostreams.IOStreams
	HTTPClient func() (*http.Client, error)
}

// NewCmdTrain creates the version command
func NewCmdTrain(f *pkg.Factory) *cobra.Command {
	tOpts := &TrainOptions{
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

func trainRun(opts *TrainOptions) error {
	// client, err := opts.HttpClient()
	// if err != nil {
	// 	return err
	// }

	// err := trainModels(client)
	// if err != nil {
	// 	return err

	return nil
}
