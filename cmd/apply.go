package cmd

import (
	"context"

	"github.com/shotat/ghrc/config"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply specs to the actual state",
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		conf, err := config.LoadFromFile(filepath)
		if err != nil {
			return err
		}

		ctx := context.Background()

		return conf.Apply(ctx)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
