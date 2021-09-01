package v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Permissions alloted to a Zotero API key.
type KeyPermissions struct {
	Library bool
	Files   bool
	Notes   bool
	Write   bool
}

type rawKeyData struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Access   struct {
		User struct {
			Library bool `json:"library"`
			Files   bool `json:"files"`
			Notes   bool `json:"notes"`
			Write   bool `json:"write"`
		} `json:"user"`
	} `json:"access"`
}

func getApiKeyData(apiKey string) rawKeyData {
	res, _ := http.DefaultClient.Get(fmt.Sprintf("%s/keys/%s", baseApi, apiKey))
	bytes, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var keyData rawKeyData
	json.Unmarshal(bytes, &keyData)
	return keyData
}

func ApiKeyExists(apiKey string) bool {
	res, _ := http.DefaultClient.Get(fmt.Sprintf("%s/keys/%s", baseApi, apiKey))
	return res.StatusCode == 200
}

func GetApiKeyPermissions(apiKey string) KeyPermissions {
	keyData := getApiKeyData(apiKey)
	return KeyPermissions{
		Library: keyData.Access.User.Library,
		Files:   keyData.Access.User.Files,
		Notes:   keyData.Access.User.Notes,
		Write:   keyData.Access.User.Write,
	}
}

func ClientOfApiKey(apiKey string) Client {
	keyData := getApiKeyData(apiKey)
	return Client{
		apiKey:   apiKey,
		username: keyData.Username,
		userId:   keyData.UserID,
	}
}
