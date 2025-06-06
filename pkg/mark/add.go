package mark

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/spf13/cobra"
)

var (
	markCmd = &cobra.Command{
		Use:   "mark",
		Short: "list all marks",
		Long:  "list all marks",
		Run: func(cmd *cobra.Command, args []string) {
			marks := storage.QueryMarks()
			for index, mark := range marks {
				fmt.Printf("[%d] %s\n", index, mark.Mark)
			}
		},
	}

	create_mark_name        string
	create_mark_color       string
	create_mark_description string
	create_mark_icon        string

	addMarkCmd = &cobra.Command{
		Use:   "add",
		Short: "add mark",
		Long:  "add mark",
		Run: func(cmd *cobra.Command, args []string) {
			storage.CreateTable()
			flag := storage.InsertMark(storage.CreateMark{
				Mark:        create_mark_name,
				Description: create_mark_description,
				Color:       create_mark_color,
				Icon:        create_mark_icon,
			})
			if flag {
				fmt.Println("add mark success")
			} else {
				fmt.Println("add mark failed")
			}
		},
	}
)

func init() {
	markCmd.AddCommand(addMarkCmd)
	addMarkCmd.Flags().StringVar(&create_mark_name, "name", "", "mark name")
	addMarkCmd.Flags().StringVar(&create_mark_color, "color", "", "mark color")
	addMarkCmd.Flags().StringVar(&create_mark_description, "description", "", "mark description")
	addMarkCmd.Flags().StringVar(&create_mark_icon, "icon", "", "mark icon")
}

func GetMarkCmd() *cobra.Command {
	return markCmd
}
