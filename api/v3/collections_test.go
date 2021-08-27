package v3

import (
	"github.com/jarcoal/httpmock"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

const toplevels = `[
    {
        "key": "KEYA",
        "version": 3,
        "library": {
            "type": "user",
            "id": "000",
            "name": "TSTUSR",
            "links": {
                "alternate": {
                    "href": "https://www.zotero.org/TSTUSR",
                    "type": "text/html"
                }
            }
        },
        "links": {
            "self": {
                "href": "https://api.zotero.org/users/000/collections/KEYA",
                "type": "application/json"
            },
            "alternate": {
                "href": "https://www.zotero.org/TSTUSR/collections/KEYA",
                "type": "text/html"
            }
        },
        "meta": {
            "numCollections": 1,
            "numItems": 2
        },
        "data": {
            "key": "KEYA",
            "version": 3,
            "name": "Fashion",
            "parentCollection": false,
            "relations": {}
        }
    },
    {
        "key": "KEYB",
        "version": 3,
        "library": {
            "type": "user",
            "id": "000",
            "name": "TSTUSR",
            "links": {
                "alternate": {
                    "href": "https://www.zotero.org/TSTUSR",
                    "type": "text/html"
                }
            }
        },
        "links": {
            "self": {
                "href": "https://api.zotero.org/users/000/collections/KEYB",
                "type": "application/json"
            },
            "alternate": {
                "href": "https://www.zotero.org/TSTUSR/collections/KEYB",
                "type": "text/html"
            }
        },
        "meta": {
            "numCollections": 0,
            "numItems": 4
        },
        "data": {
            "key": "KEYB",
            "version": 3,
            "name": "Cologne",
            "parentCollection": false,
            "relations": {}
        }
    }
]`

const sub_fashion = `
[
    {
        "key": "KEYC",
        "version": 3,
        "library": {
            "type": "user",
            "id": "000",
            "name": "TSTUSR",
            "links": {
                "alternate": {
                    "href": "https://www.zotero.org/TSTUSR",
                    "type": "text/html"
                }
            }
        },
        "links": {
            "self": {
                "href": "https://api.zotero.org/users/000/collections/KEYC",
                "type": "application/json"
            },
            "alternate": {
                "href": "https://www.zotero.org/TSTUSR/collections/KEYC",
                "type": "text/html"
            }
        },
        "meta": {
            "numCollections": 0,
            "numItems": 5
        },
        "data": {
            "key": "KEYC",
            "version": 3,
            "name": "French Vintage",
            "parentCollection": false,
            "relations": {}
        }
    }
]`

func TestGetCollections(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.zotero.org/users/TSTUSR/collections/top",
		httpmock.NewBytesResponder(200, []byte(toplevels)))
	httpmock.RegisterResponder("GET", "https://api.zotero.org/users/TSTUSR/collections/KEYA/collections",
		httpmock.NewStringResponder(200, sub_fashion))

	c := NewClient("test-api-key", "TSTUSR")
	collections := c.GetCollections()

	assert.Assert(t, is.Equal(httpmock.GetTotalCallCount(), 2))

	assert.Assert(t, is.Len(collections, 3))

	fashion, cologne, french_vintage := collections[0], collections[1], collections[2]

	assert.Assert(t, is.Equal(fashion.Name, "Fashion"))
	assert.Assert(t, is.Equal(fashion.Key.Value, "KEYA"))
	assert.Assert(t, fashion.IsTopLevel)
	assert.Assert(t, is.Equal(fashion.Parent, EmptyItemKey))
	assert.Assert(t, is.Equal(fashion.NumSubCollections, 1))
	assert.Assert(t, is.Equal(fashion.NumItems, 2))

	assert.Assert(t, is.Equal(cologne.Name, "Cologne"))
	assert.Assert(t, is.Equal(cologne.Key.Value, "KEYB"))
	assert.Assert(t, cologne.IsTopLevel)
	assert.Assert(t, is.Equal(cologne.Parent, EmptyItemKey))
	assert.Assert(t, is.Equal(cologne.NumSubCollections, 0))
	assert.Assert(t, is.Equal(cologne.NumItems, 4))

	assert.Assert(t, is.Equal(french_vintage.Name, "French Vintage"))
	assert.Assert(t, is.Equal(french_vintage.Key.Value, "KEYC"))
	assert.Assert(t, is.Equal(french_vintage.IsTopLevel, false))
	assert.Assert(t, is.Equal(french_vintage.Parent, fashion.Key))
	assert.Assert(t, is.Equal(french_vintage.NumSubCollections, 0))
	assert.Assert(t, is.Equal(french_vintage.NumItems, 5))
}
