package commands

import "github.com/spf13/cobra"

func newServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve fiber webserver",
		Run: func(cmd *cobra.Command, args []string) {
			VexoraApp.Start()
		},
	}

	return cmd
}
