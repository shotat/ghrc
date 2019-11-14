package cmd

import (
	"context"
	"fmt"

	"github.com/shotat/ghrc/config"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import <repository name>",
	Short: "Import an existing repository state",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ctx := context.Background()
		conf, err := config.Import(ctx, name)
		if err != nil {
			return err
		}

		yaml, err := conf.ToYAML()
		if err != nil {
			return err
		}
		fmt.Println(yaml)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
