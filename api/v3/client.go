package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
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

// Like `get`, but for Zotero APIs that may be paginated. Returns a list of raw
// byte reponses, one per page.
func (c *Client) getWithPagination(subroute string) [][]byte {
	uri := func(query string) string {
		return fmt.Sprintf("%s/users/%d/%s?%s", baseApi, c.userId, subroute, query)
	}

	const limitPerPage = 50
	_, headers := c.getUriWithHeaders(uri("limit=1"))
	numItems, _ := strconv.Atoi(headers.Get("Total-Results"))

	ch := make(chan []byte, 10)
	var wg sync.WaitGroup

	for start := 0; start < numItems; start += limitPerPage {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			data, _ := c.getUriWithHeaders(uri(fmt.Sprintf("start=%d&limit=%d", start, limitPerPage)))
			ch <- data
		}(start)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var allBytes [][]byte
	for bytes := range ch {
		allBytes = append(allBytes, bytes)
	}
	return allBytes
}
