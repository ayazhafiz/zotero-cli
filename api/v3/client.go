package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// A client for the Zotero API, belonging to a single user.
type Client struct {
	// A user-provisioned API key.
	apiKey string

	// The user this key belongs to.
	username string

	// The user ID.
	userId int
}

func (c *Client) ApiKey() string {
	return c.apiKey
}

// An ItemKey is the Zotero API's opaque handle to an item in a library.
type ItemKey struct {
	// The string value of the key.
	Value string
}

var EmptyItemKey = ItemKey{Value: ""}

const baseApi = "https://api.zotero.org"

// Sends a `GET` request to `subroute` and returns the raw bytes of the response
// body, along with repsonse headers.
func (c *Client) getUriWithHeaders(uri string) ([]byte, http.Header) {
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Add("Zotero-API-Version", "3")
	req.Header.Add("Zotero-API-Key", c.apiKey)

	res, _ := http.DefaultClient.Do(req)
	bytes, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return bytes, res.Header
}

// Sends a `GET` request to `subroute` and returns the raw bytes of the response
// body.
func (c *Client) get(subroute string) []byte {
	uri := fmt.Sprintf("%s/users/%d/%s", baseApi, c.userId, subroute)
	bytes, _ := c.getUriWithHeaders(uri)
	return bytes
}

// https://www.zotero.org/support/dev/web_api/v3/basics#link_header
var reNextPage = regexp.MustCompile(`.*<(.*)>; rel="next"`)

// Like `get`, but for Zotero APIs that may be paginated. Returns a list of raw
// byte reponses, one per page.
func (c *Client) getWithPagination(subroute string) [][]byte {
	var allBytes [][]byte
	nextPage := fmt.Sprintf("%s/users/%d/%s", baseApi, c.userId, subroute)
	for true {
		bytes, headers := c.getUriWithHeaders(nextPage)
		allBytes = append(allBytes, bytes)
		matchNextPage := reNextPage.FindStringSubmatch(headers.Get("Link"))
		if matchNextPage == nil {
			break
		}
		nextPage = matchNextPage[1]
	}
	return allBytes
}
