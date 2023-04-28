package cmd

import (
	"fmt"
	"runtime"

	"github.com/brpaz/copier-run/internal/version"
	"github.com/spf13/cobra"
)

// NewVersionCmd returns a new instance of the version command
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			fmt.Fprintf(out, "Version: %s\n", version.Version)
			fmt.Fprintf(out, "Git commit: %s\n", version.GitCommit)
			fmt.Fprintf(out, "Build date: %s\n", version.BuildDate)
			fmt.Fprintf(out, "Go version: %s\n", runtime.Version())

			return nil
		},
	}
}
