package cmd

import (
	"fmt"

	"github.com/Ysoding/xkcd/comic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch all xkcd comics",
	Run: func(cmd *cobra.Command, args []string) {
		err := comic.NewDownloader().Download()
		if err != nil {
			fmt.Printf("Download error: %v\n", err)
		}
	},
}
