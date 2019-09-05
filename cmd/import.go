package cmd

import (
	"context"
	"fmt"

	"bytes"

	"github.com/shotat/ghrc/config"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		owner, err := cmd.Flags().GetString("owner")
		if err != nil {
			return err
		}
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			return err
		}
		ctx := context.Background()
		conf, err := config.Import(ctx, owner, repo)
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(nil)
		err = yaml.NewEncoder(buf).Encode(conf)
		fmt.Println(buf.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().String("repo", "", "repository name")
	importCmd.Flags().String("owner", "", "owner name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
