package cmd

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/spf13/cobra"
)

func init() {
	storage.CreateTable()
}

var rootCmd = &cobra.Command{
	Use:   "filemark",
	Short: "filemark is a tool",
	Long:  "Can add tag to file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
