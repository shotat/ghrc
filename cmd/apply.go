package cmd

import (
	"context"
	"fmt"

	"github.com/shotat/ghrc/config"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply specs to the actual state",
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

		name, _ := cmd.Flags().GetString("name")
		if name != "" {
			conf.Metadata.Name = name
		}

		if err := conf.Apply(ctx); err != nil {
			return err
		}
		fmt.Println("successfully applied!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringP("filename", "f", ".ghrc.yaml", "config file name")
	applyCmd.Flags().String("name", "", "repository full name")
}
