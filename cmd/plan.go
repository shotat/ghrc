package cmd

import (
	"context"

	"github.com/shotat/ghrc/config"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Check expected changes without changing the actual state",
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}

		conf, err := config.LoadFromFile(filepath)
		if err != nil {
			return err
		}

		ctx := context.Background()

		return conf.Plan(ctx)
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.Flags().StringP("filename", "f", ".ghrc.yaml", "config file name")
}
