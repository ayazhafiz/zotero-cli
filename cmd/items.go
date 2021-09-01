package cmd

import (
	"fmt"
	api "github.com/ayazhafiz/zotero-cli/api/v3"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Search for items in your Zotero library",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		items := client.GetItems()
		idx, err := fuzzyfinder.Find(items, func(i int) string { return items[i].Title },
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i < 0 {
					return ""
				}
				item := &items[i]
				return fmt.Sprintf("%s\n%s (%s)", item.Title,
					api.PrintAuthors(item.Authors), item.ParsedDate)
			}),
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "You found a fatal error. Please report this.\n%s", err)
			os.Exit(101)
		}
		open.Start(items[idx].LibraryUrl)
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)
}
