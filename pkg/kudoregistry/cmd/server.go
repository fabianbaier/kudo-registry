package cmd

import (
	"fmt"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/cmd/server"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/config"
	"github.com/spf13/cobra"
)

var (
	serverExample = `
		# Start the KUDO Registry with a local storage backend
		kudo-registry server start --backend=local`
)

// NewInstallCmd creates the install command for the CLI
func NewServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:          "server",
		Short:        "-> Server KUDO registry.",
		Long:         `Start the server.`,
		Example:      serverExample,
		SilenceUsage: true,
	}

	serverCmd.Flags().StringVar(&vars.StorageBackend, "backend", "local", "Kudo Storage Backend. (default \"local\")")

	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	serverCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(serverCmd.OutOrStderr(), usageFmt, serverCmd.UseLine(), serverCmd.Flags().FlagUsages())
		return nil
	})

	serverCmd.AddCommand(server.NewStartCmd())

	return serverCmd
}
