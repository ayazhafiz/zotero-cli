package cmd

import (
	"fmt"
	api "github.com/ayazhafiz/zotero-cli/api/v3"
	"os"

	"github.com/spf13/viper"
)

func getClient() api.Client {
	viper.ReadInConfig()
	apiKey := viper.GetString("api_key")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "No Zotero API key found.\nPlease run `zotero config init'.\n")
		os.Exit(1)
	}
	return api.ClientOfApiKey(apiKey)
}
