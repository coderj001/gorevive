package cmd

import (
	"fmt"

	"github.com/coderj001/GoRevive/internal/helpers"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project]",
	Short: "Create a new project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a filename.")
			return
		}
		project := args[0]
		err := helpers.NewFile(project)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", project, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
