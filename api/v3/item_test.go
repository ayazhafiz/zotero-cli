package v3

import (
	"github.com/jarcoal/httpmock"
	. "gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

const page1 = `[{
        "key": "KEYA",
        "version": 18,
        "library": {
            "type": "user",
            "id": 123,
            "name": "janedoe",
            "links": {
                "alternate": {
                    "href": "https://www.zotero.org/janedoe",
                    "type": "text/html"
                }
            }
        },
        "links": {
            "self": {
                "href": "https://api.zotero.org/users/123/items/KEYA",
                "type": "application/json"
            },
            "alternate": {
                "href": "https://www.zotero.org/janedoe/items/KEYA",
                "type": "text/html"
            },
            "attachment": {
                "href": "https://api.zotero.org/users/123/items/VUNPABHE",
                "type": "application/json",
                "attachmentType": "application/pdf",
                "attachmentSize": 1291508
            }
        },
        "meta": {
            "creatorSummary": "Landin",
            "parsedDate": "1966",
            "numChildren": 1
        },
        "data": {
            "key": "KEYA",
            "version": 18,
            "itemType": "journalArticle",
            "title": "The next 700 programming languages",
            "creators": [
                {
                    "creatorType": "author",
                    "firstName": "P. J.",
                    "lastName": "Landin"
                }
            ],
            "abstractNote": "ABSTRACTA",
            "publicationTitle": "Communications of the ACM",
            "volume": "9",
            "issue": "3",
            "pages": "157-166",
            "date": "03/1966",
            "series": "",
            "seriesTitle": "",
            "seriesText": "",
            "journalAbbreviation": "Commun. ACM",
            "language": "en",
            "DOI": "10.1145/365230.365257",
            "ISSN": "0001-0782, 1557-7317",
            "shortTitle": "",
            "url": "https://dl.acm.org/doi/10.1145/365230.365257",
            "accessDate": "2000-01-01T00:12:00Z",
            "archive": "",
            "archiveLocation": "",
            "libraryCatalog": "DOI.org (Crossref)",
            "callNumber": "",
            "rights": "",
            "extra": "",
            "tags": [
								{
										"tag": "MYTAG1"
								}
						],
            "collections": [
                "COLLECTIONA"
            ],
            "relations": {},
            "dateAdded": "2000-01-01T00:12:48Z",
            "dateModified": "2000-02-01T15:37:05Z"
        }
    }
]`

const page2 = `[{
        "key": "KEYB",
        "version": 25,
        "library": {
            "type": "user",
            "id": 123,
            "name": "janedoe",
            "links": {
                "alternate": {
                    "href": "https://www.zotero.org/janedoe",
                    "type": "text/html"
                }
            }
        },
        "links": {
            "self": {
                "href": "https://api.zotero.org/users/123/items/KEYB",
                "type": "application/json"
            },
            "alternate": {
                "href": "https://www.zotero.org/janedoe/items/KEYB",
                "type": "text/html"
            },
            "attachment": {
                "href": "https://api.zotero.org/users/123/items/UX7DDNBG",
                "type": "application/json",
                "attachmentType": "application/pdf",
                "attachmentSize": 395812
            }
        },
        "meta": {
            "creatorSummary": "Eisenberg et al.",
            "parsedDate": "2021-08-22",
            "numChildren": 1
        },
        "data": {
            "key": "KEYB",
            "version": 25,
            "itemType": "journalArticle",
            "title": "An existential crisis resolved: type inference for first-class existential types",
            "creators": [
                {
                    "creatorType": "author",
                    "firstName": "Richard A.",
                    "lastName": "Eisenberg"
                },
                {
                    "creatorType": "author",
                    "firstName": "Guillaume",
                    "lastName": "Duboc"
                },
                {
                    "creatorType": "author",
                    "firstName": "Stephanie",
                    "lastName": "Weirich"
                },
                {
                    "creatorType": "author",
                    "firstName": "Daniel",
                    "lastName": "Lee"
                }
            ],
            "abstractNote": "ABSTRACTB",
            "publicationTitle": "Proceedings of the ACM on Programming Languages",
            "volume": "5",
            "issue": "ICFP",
            "pages": "1-29",
            "date": "2021-08-22",
            "series": "",
            "seriesTitle": "",
            "seriesText": "",
            "journalAbbreviation": "Proc. ACM Program. Lang.",
            "language": "en",
            "DOI": "10.1145/3473569",
            "ISSN": "2475-1421",
            "shortTitle": "An existential crisis resolved",
            "url": "https://dl.acm.org/doi/10.1145/3473569",
            "accessDate": "2021-08-01T22:19:41Z",
            "archive": "",
            "archiveLocation": "",
            "libraryCatalog": "DOI.org (Crossref)",
            "callNumber": "",
            "rights": "",
            "extra": "",
            "tags": [
                {
                    "tag": "MYTAG2"
                },
                {
                    "tag": "MYTAG3"
                }
            ],
            "collections": [
                "COLLECTIONA",
                "COLLECTIONB"
            ],
            "relations": {},
            "dateAdded": "2021-08-01T22:19:41Z",
            "dateModified": "2021-08-01T22:20:00Z"
        }
    }
]`

func TestGetItems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	headResp := httpmock.NewStringResponse(200, "")
	headResp.Header.Set("Total-Results", `80`)

	httpmock.RegisterResponder("GET", "https://api.zotero.org/users/123/items/top?limit=1",
		httpmock.ResponderFromResponse(headResp))

	httpmock.RegisterResponder("GET", "https://api.zotero.org/users/123/items/top?start=0&limit=50",
		httpmock.NewStringResponder(200, page1))

	httpmock.RegisterResponder("GET", "https://api.zotero.org/users/123/items/top?start=50&limit=50",
		httpmock.NewStringResponder(200, page2))

	c := Client{"test-api-key", "TSTUSR", 123}
	items := c.GetItems()

	Assert(t, is.Equal(httpmock.GetTotalCallCount(), 3))

	Assert(t, is.Len(items, 2))

	next700, extCrisis := items[0], items[1]

	Assert(t, is.Equal(next700.Key.Value, "KEYA"))
	Assert(t, is.Equal(next700.Title, "The next 700 programming languages"))
	Assert(t, next700.Type.Is(journalArticle))
	Assert(t, is.Equal(next700.AbstractNote, "ABSTRACTA"))
	Assert(t, is.DeepEqual(next700.Authors, []Author{
		{"P. J.", "Landin"},
	}))
	Assert(t, is.Equal(next700.Date, "03/1966"))
	Assert(t, is.Equal(next700.ParsedDate, "1966"))
	Assert(t, is.Equal(next700.Url, "https://dl.acm.org/doi/10.1145/365230.365257"))
	Assert(t, is.DeepEqual(next700.Tags, []string{"MYTAG1"}))
	Assert(t, is.DeepEqual(next700.Collections, []ItemKey{{"COLLECTIONA"}}))
	Assert(t, is.DeepEqual(next700.LibraryUrl, "https://www.zotero.org/janedoe/items/KEYA"))

	Assert(t, is.Equal(extCrisis.Key.Value, "KEYB"))
	Assert(t, is.Equal(extCrisis.Title, "An existential crisis resolved: type inference for first-class existential types"))
	Assert(t, extCrisis.Type.Is(journalArticle))
	Assert(t, is.Equal(extCrisis.AbstractNote, "ABSTRACTB"))
	Assert(t, is.DeepEqual(extCrisis.Authors, []Author{
		{"Richard A.", "Eisenberg"},
		{"Guillaume", "Duboc"},
		{"Stephanie", "Weirich"},
		{"Daniel", "Lee"},
	}))
	Assert(t, is.Equal(extCrisis.Date, "2021-08-22"))
	Assert(t, is.Equal(extCrisis.ParsedDate, "2021-08-22"))
	Assert(t, is.Equal(extCrisis.Url, "https://dl.acm.org/doi/10.1145/3473569"))
	Assert(t, is.DeepEqual(extCrisis.Tags, []string{"MYTAG2", "MYTAG3"}))
	Assert(t, is.DeepEqual(extCrisis.Collections, []ItemKey{{"COLLECTIONA"}, {"COLLECTIONB"}}))
	Assert(t, is.DeepEqual(extCrisis.LibraryUrl, "https://www.zotero.org/janedoe/items/KEYB"))
}
