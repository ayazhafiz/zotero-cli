package cmd

import (
	"errors"
	"fmt"
	api "github.com/ayazhafiz/zotero-cli/api/v3"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Zotero CLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct{ Client api.Client }{}
		survey.Ask(confiqS, &answers)
		viper.Set("api_key", answers.Client.ApiKey())
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing configuration file: %s\n", err)
		} else {
			fmt.Printf("Configuration written to %s.\n", cfgFile)
		}
	},
}

var confiqS = []*survey.Question{
	{
		Name:   "client",
		Prompt: &survey.Input{Message: "Please provide a Zotero API key (https://www.zotero.org/settings/keys):"},
		Validate: func(key interface{}) error {
			str, ok := key.(string)
			if !ok {
				return errors.New("API key must be a string.")
			}
			if !api.ApiKeyExists(str) {
				return errors.New("API key does not exist.")
			}
			perms := api.GetApiKeyPermissions(str)
			if !perms.Library || !perms.Files || !perms.Notes || !perms.Write {
				return errors.New("API key must have library, files, notes, and write permission.")
			}
			return nil
		},
		Transform: func(key interface{}) interface{} {
			str, _ := key.(string)
			return api.ClientOfApiKey(str)
		},
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
}
