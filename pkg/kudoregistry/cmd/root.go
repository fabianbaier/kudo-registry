package cmd

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/version"
	"github.com/spf13/cobra"
)

// NewKudoctlCmd creates a new root command for kudoctl
func NewKudoRegistryCmd() *cobra.Command {
	cmd := &cobra.Command{
		// Workaround or Compromise as "kubectl kudo" would result in Usage print out "kubectl install <name> [flags]"
		Use:   "kudo-registry",
		Short: "CLI to start the KUDO Registry server.",
		Long: `
KUDO Registry.
`,
		Example: `
	# Run the KUDO Registry via
	kudo-registry server start

`,
		Version: version.Get().GitVersion,
	}

	cmd.AddCommand(NewServerCmd())

	return cmd
}
