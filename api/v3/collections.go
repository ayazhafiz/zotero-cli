package v3

import (
	"encoding/json"
	"fmt"
)

// Represents a Zotero library collection.
type Collection struct {
	// The human-readable name of this collection.
	Name string

	// The Zotero API-internal key.
	Key ItemKey

	// Whether the collection is on the toplevel of the user library.
	// When true, `Parent == EmptyItemKey`.
	IsTopLevel bool

	// The parent collection, if it exists.
	// When this collection has no parent, this is `EmptyItemKey`.
	Parent ItemKey

	// Number of direct subcollections.
	NumSubCollections int

	// Number of non-collection items in the collections.
	NumItems int
}

type rawCollection struct {
	Key  string `json:"key"`
	Meta struct {
		NumSubCollections int `json:"numCollections"`
		NumItems          int `json:"numItems"`
	} `json:"meta"`
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
}

func convertRaw(rc rawCollection) Collection {
	return Collection{
		Name:              rc.Data.Name,
		Key:               ItemKey{Value: rc.Key},
		IsTopLevel:        true,         // overridden in GetCollections
		Parent:            EmptyItemKey, // overridden in GetCollections
		NumSubCollections: rc.Meta.NumSubCollections,
		NumItems:          rc.Meta.NumItems,
	}
}

func unmarshalCollections(bytes []byte) []Collection {
	var raw []rawCollection
	json.Unmarshal(bytes, &raw)
	collections := make([]Collection, len(raw))
	for i, rc := range raw {
		collections[i] = convertRaw(rc)
	}
	return collections
}

func (c *Client) GetCollections() []Collection {
	var allCollections []Collection

	visitQueue := unmarshalCollections(c.get("collections/top"))

	for len(visitQueue) > 0 {
		head := visitQueue[0]
		visitQueue = visitQueue[1:]
		allCollections = append(allCollections, head)
		if head.NumSubCollections > 0 {
			subCollections := unmarshalCollections(
				c.get(fmt.Sprintf("collections/%s/collections", head.Key.Value)),
			)
			for i := range subCollections {
				subCollections[i].IsTopLevel = false
				subCollections[i].Parent = head.Key
			}
			visitQueue = append(visitQueue, subCollections...)
		}
	}

	return allCollections
}
