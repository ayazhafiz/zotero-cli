package v3

import (
	"encoding/json"
	. "gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"io/ioutil"
	"net/http"
	"testing"
)

type zoteroItemType struct {
	ItemType  string `json:"itemType"`
	Localized string `json:"localized"`
}

func getZoteroItemTypes() []zoteroItemType {
	res, _ := http.DefaultClient.Get("https://api.zotero.org/itemTypes?locale=en-US")
	bytes, _ := ioutil.ReadAll(res.Body)
	var itemTypes []zoteroItemType
	json.Unmarshal(bytes, &itemTypes)
	return itemTypes
}

func TestZoteroItemTypes(t *testing.T) {
	for _, zotIt := range getZoteroItemTypes() {
		it, err := ItemTypeOfString(zotIt.ItemType)
		Assert(t, is.Nil(err))
		Assert(t, is.Equal(it.ApiStringOfItemType(), zotIt.ItemType))
		Assert(t, is.Equal(it.EnglishStringOfItemType(), zotIt.Localized))
	}
}
