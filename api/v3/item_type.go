package v3

import (
	"errors"
)

const (
	artwork = iota
	audioRecording
	bill
	blogPost
	book
	bookSection
	caseItem
	computerProgram
	conferencePaper
	dictionaryEntry
	document
	email
	encyclopediaArticle
	film
	forumPost
	hearing
	instantMessage
	interview
	journalArticle
	letter
	magazineArticle
	manuscript
	mapItem
	newspaperArticle
	note
	patent
	podcast
	presentation
	radioBroadcast
	report
	statute
	tvBroadcast
	thesis
	videoRecording
	webpage
)

type ItemType struct {
	i int
}

// Converts a string `itemType` from the Zotero API into the internal `ItemType`
// representation.
func ItemTypeOfString(s string) (ItemType, error) {
	i := -1
	switch s {
	case "artwork":
		i = artwork
	case "audioRecording":
		i = audioRecording
	case "bill":
		i = bill
	case "blogPost":
		i = blogPost
	case "book":
		i = book
	case "bookSection":
		i = bookSection
	case "case":
		i = caseItem
	case "computerProgram":
		i = computerProgram
	case "conferencePaper":
		i = conferencePaper
	case "dictionaryEntry":
		i = dictionaryEntry
	case "document":
		i = document
	case "email":
		i = email
	case "encyclopediaArticle":
		i = encyclopediaArticle
	case "film":
		i = film
	case "forumPost":
		i = forumPost
	case "hearing":
		i = hearing
	case "instantMessage":
		i = instantMessage
	case "interview":
		i = interview
	case "journalArticle":
		i = journalArticle
	case "letter":
		i = letter
	case "magazineArticle":
		i = magazineArticle
	case "manuscript":
		i = manuscript
	case "map":
		i = mapItem
	case "newspaperArticle":
		i = newspaperArticle
	case "note":
		i = note
	case "patent":
		i = patent
	case "podcast":
		i = podcast
	case "presentation":
		i = presentation
	case "radioBroadcast":
		i = radioBroadcast
	case "report":
		i = report
	case "statute":
		i = statute
	case "tvBroadcast":
		i = tvBroadcast
	case "thesis":
		i = thesis
	case "videoRecording":
		i = videoRecording
	case "webpage":
		i = webpage
	default:
		return ItemType{i}, errors.New("unknown item type \"" + s + "\"")
	}
	return ItemType{i}, nil
}

// Converts an `ItemType` to a string `itemType` consumable by the Zotero web
// API.
func (it *ItemType) ApiStringOfItemType() string {
	switch it.i {
	case artwork:
		return "artwork"
	case audioRecording:
		return "audioRecording"
	case bill:
		return "bill"
	case blogPost:
		return "blogPost"
	case book:
		return "book"
	case bookSection:
		return "bookSection"
	case caseItem:
		return "case"
	case computerProgram:
		return "computerProgram"
	case conferencePaper:
		return "conferencePaper"
	case dictionaryEntry:
		return "dictionaryEntry"
	case document:
		return "document"
	case email:
		return "email"
	case encyclopediaArticle:
		return "encyclopediaArticle"
	case film:
		return "film"
	case forumPost:
		return "forumPost"
	case hearing:
		return "hearing"
	case instantMessage:
		return "instantMessage"
	case interview:
		return "interview"
	case journalArticle:
		return "journalArticle"
	case letter:
		return "letter"
	case magazineArticle:
		return "magazineArticle"
	case manuscript:
		return "manuscript"
	case mapItem:
		return "map"
	case newspaperArticle:
		return "newspaperArticle"
	case note:
		return "note"
	case patent:
		return "patent"
	case podcast:
		return "podcast"
	case presentation:
		return "presentation"
	case radioBroadcast:
		return "radioBroadcast"
	case report:
		return "report"
	case statute:
		return "statute"
	case tvBroadcast:
		return "tvBroadcast"
	case thesis:
		return "thesis"
	case videoRecording:
		return "videoRecording"
	case webpage:
		return "webpage"
	default:
		panic("unreachable")
	}
}

// Converts an `ItemType` to an english-localized string.
func (it *ItemType) EnglishStringOfItemType() string {
	switch it.i {
	case artwork:
		return "Artwork"
	case audioRecording:
		return "Audio Recording"
	case bill:
		return "Bill"
	case blogPost:
		return "Blog Post"
	case book:
		return "Book"
	case bookSection:
		return "Book Section"
	case caseItem:
		return "Case"
	case computerProgram:
		return "Computer Program"
	case conferencePaper:
		return "Conference Paper"
	case dictionaryEntry:
		return "Dictionary Entry"
	case document:
		return "Document"
	case email:
		return "E-mail"
	case encyclopediaArticle:
		return "Encyclopedia Article"
	case film:
		return "Film"
	case forumPost:
		return "Forum Post"
	case hearing:
		return "Hearing"
	case instantMessage:
		return "Instant Message"
	case interview:
		return "Interview"
	case journalArticle:
		return "Journal Article"
	case letter:
		return "Letter"
	case magazineArticle:
		return "Magazine Article"
	case manuscript:
		return "Manuscript"
	case mapItem:
		return "Map"
	case newspaperArticle:
		return "Newspaper Article"
	case note:
		return "Note"
	case patent:
		return "Patent"
	case podcast:
		return "Podcast"
	case presentation:
		return "Presentation"
	case radioBroadcast:
		return "Radio Broadcast"
	case report:
		return "Report"
	case statute:
		return "Statute"
	case tvBroadcast:
		return "TV Broadcast"
	case thesis:
		return "Thesis"
	case videoRecording:
		return "Video Recording"
	case webpage:
		return "Web Page"
	default:
		panic("unreachable")
	}
}

// Checks an `ItemType` for a certain type.
func (it *ItemType) Is(what int) bool {
	return it.i == what
}
