package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mittwald/spacectl/spacefile"
)

var spaceFile string

var spacesCmd = &cobra.Command{
	Use:     "spaces",
	Aliases: []string{"space", "spc", "s"},
	Short:   "Manage Spaces",
	Long:    `This command can be used to manage Spaces`,
}

func init() {
	RootCmd.AddCommand(spacesCmd)

	spacesCmd.PersistentFlags().StringVarP(&spaceFile, "spacefile", "f", "./" + spacefile.DefaultFilename, "Use Space defined in this file")
}
