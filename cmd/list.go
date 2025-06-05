package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	file_path string

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all file",
		Long:  "list all file",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := os.Stat(file_path)
			if err != nil {
				println("err", err)
			}
			if info.IsDir() {
				dirs, _ := os.ReadDir(file_path)
				for _, item := range dirs {
					prefix := "[F]"
					if item.IsDir() {
						prefix = "[D]"
					}
					println(prefix, item.Name())
				}
			} else {
				println("[F]", info.Name())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&file_path, "path", "", "path to list")
}
