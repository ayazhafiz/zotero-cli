package v3

import (
	"encoding/json"
)

// A Top-Level Item in a Zotero library, which can have items (e.g. attachments)
// under it.
type Item struct {
	Key ItemKey
	// Title of the item.
	Title string
	// Type of the item.
	Type ItemType
	// The note in the "abstract" field of the item, if any.
	AbstractNote string
	// The authors of the item. May be empty.
	Authors []Author
	// The publication date of the item.
	Date string
	// User-friendly "meta parsed data" of the item, provided by Zotero.
	ParsedDate string
	// The web URL the item was recovered from.
	Url string
	// Tags attached to the item, if any.
	Tags []string
	// Collections the item belongs to, if any.
	Collections []ItemKey
	// URL to view the item in Zotero online library.
	LibraryUrl string
}

type rawItem struct {
	Key   string `json:"key"`
	Links struct {
		Alternate struct {
			Href string `json:"href"`
		} `json:"alternate"`
	} `json:"links"`
	Meta struct {
		ParsedDate string `json:"parsedDate"`
	} `json:"meta"`
	Data struct {
		ItemType     string `json:"itemType"`
		Title        string `json:"title"`
		AbstractNote string `json:"abstractNote"`
		Creators     []struct {
			CreatorType string `json:"creatorType"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
		} `json:"creators"`
		Date string `json:"date"`
		Url  string `json:"url"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
		Collections []string `json:"collections"`
	} `json:"data"`
}

func convertRawItem(ri *rawItem) Item {
	itemType, _ := ItemTypeOfString(ri.Data.ItemType)
	var authors []Author
	for _, creator := range ri.Data.Creators {
		if creator.CreatorType == "author" {
			authors = append(authors, Author{creator.FirstName, creator.LastName})
		}
	}
	tags := make([]string, len(ri.Data.Tags))
	for i, tag := range ri.Data.Tags {
		tags[i] = tag.Tag
	}
	collections := make([]ItemKey, len(ri.Data.Tags))
	for i, collectionKey := range ri.Data.Collections {
		collections[i] = ItemKey{collectionKey}
	}
	return Item{
		Key:          ItemKey{ri.Key},
		Title:        ri.Data.Title,
		Type:         itemType,
		AbstractNote: ri.Data.AbstractNote,
		Authors:      authors,
		Date:         ri.Data.Date,
		ParsedDate:   ri.Meta.ParsedDate,
		Url:          ri.Data.Url,
		Tags:         tags,
		Collections:  collections,
		LibraryUrl:   ri.Links.Alternate.Href,
	}
}

// Retrieves all top-level items (i.e. those that do not have a parent item) in
// a user's library.
func (c *Client) GetItems() []Item {
	allRawBytes := c.getWithPagination("items/top")
	var allRawItems []rawItem
	for _, bytes := range allRawBytes {
		var raw []rawItem
		json.Unmarshal(bytes, &raw)
		allRawItems = append(allRawItems, raw...)
	}
	allItems := make([]Item, len(allRawItems))
	for i, rawItem := range allRawItems {
		allItems[i] = convertRawItem(&rawItem)
	}
	return allItems
}
