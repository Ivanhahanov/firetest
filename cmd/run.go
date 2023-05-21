package cmd

import (
	"github.com/PoteeDev/firetest/fire"
	"github.com/spf13/cobra"
)

var c string

func init() {
	rootCmd.AddCommand(run)
	rootCmd.PersistentFlags().StringVarP(&c, "config", "c", "fire.yaml", "config path")
}

var run = &cobra.Command{
	Use:   "run",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		fire.Run(c)
		return nil
	},
}
