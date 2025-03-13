package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/aztecher/vellun/pkg/logging"
	"github.com/aztecher/vellun/pkg/version"
	"github.com/spf13/cobra"
)

func NewAgentCmd() *cobra.Command {
	return NewDefaultAgentCmd()
}

func NewDefaultAgentCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "vellun-agent",
		Short: "Run the vellun agent",
		Run: func(cmd *cobra.Command, args []string) {
			if v, _ := cmd.Flags().GetBool("version"); v {
				fmt.Printf("%s %s\n", cmd.Name(), version.Version)
				os.Exit(0)
			}
			// viper init
			// option validation
			logger := logging.DefaultSlogLogger
			// Run
			if err := run(logger); err != nil {
				logger.Error(err.Error())
			}
		},
	}
	// FlagSets

	// AddCmd

	// InitGlobalFlags

	// OnInitialize
	return rootCmd
}

func run(log *slog.Logger) error {
	return fmt.Errorf("daemon not implemented yet")
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
