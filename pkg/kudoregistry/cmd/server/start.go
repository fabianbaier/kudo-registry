package server

import (
	"fmt"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/cmd/server/start"
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/config"

	"github.com/spf13/cobra"
)

var (
	serverExample = `
		# Start the KUDO Registry with a local storage backend
		kudo-registry server start --backend=local`
)

// NewInstallCmd creates the install command for the CLI
func NewStartCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:          "start",
		Short:        "-> Start the KUDO registry.",
		Long:         `Start the server.`,
		Example:      serverExample,
		RunE:         start.CmdErrorProcessor,
		SilenceUsage: true,
	}

	serverCmd.Flags().StringVar(&vars.StorageBackend, "backend", "local", "Kudo Storage Backend. (default \"local\")")
	serverCmd.Flags().StringVar(&vars.StorageBackendRootDir, "root-dir", "", "Kudo Storage Backend Root Dir. (default \"\")")

	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	serverCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(serverCmd.OutOrStderr(), usageFmt, serverCmd.UseLine(), serverCmd.Flags().FlagUsages())
		return nil
	})

	return serverCmd
}
