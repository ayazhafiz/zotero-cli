package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// A client for the Zotero API, belonging to a single user.
type Client struct {
	// A user-provisioned API key.
	apiKey string

	// The user ID.
	userId ItemKey
}

// An ItemKey is the Zotero API's opaque handle to an item in a library.
type ItemKey struct {
	// The string value of the key.
	Value string
}

var EmptyItemKey = ItemKey{Value: ""}

const baseApi = "https://api.zotero.org"

func (c *Client) get(subroute string) []byte {
	uri := fmt.Sprintf("%s/users/%s/%s", baseApi, c.userId.Value, subroute)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Add("Zotero-API-Version", "3")
	req.Header.Add("Zotero-API-Key", c.apiKey)

	fmt.Fprintf(os.Stderr, "URI: %s\n\n", uri)

	res, _ := http.DefaultClient.Do(req)
	bytes, _ := ioutil.ReadAll(res.Body)
	return bytes
}

func NewClient(apiKey string, userId string) Client {
	return Client{apiKey: apiKey, userId: ItemKey{Value: userId}}
}
