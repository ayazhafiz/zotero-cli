package v3

import (
	"errors"
)

const (
	Artwork = iota
	AudioRecording
	Bill
	BlogPost
	Book
	BookSection
	Case
	ComputerProgram
	ConferencePaper
	DictionaryEntry
	Document
	Email
	EncyclopediaArticle
	Film
	ForumPost
	Hearing
	InstantMessage
	Interview
	JournalArticle
	Letter
	MagazineArticle
	Manuscript
	Map
	NewspaperArticle
	Note
	Patent
	Podcast
	Presentation
	RadioBroadcast
	Report
	Statute
	TVBroadcast
	Thesis
	VideoRecording
	Webpage
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
		i = Artwork
	case "audioRecording":
		i = AudioRecording
	case "bill":
		i = Bill
	case "blogPost":
		i = BlogPost
	case "book":
		i = Book
	case "bookSection":
		i = BookSection
	case "case":
		i = Case
	case "computerProgram":
		i = ComputerProgram
	case "conferencePaper":
		i = ConferencePaper
	case "dictionaryEntry":
		i = DictionaryEntry
	case "document":
		i = Document
	case "email":
		i = Email
	case "encyclopediaArticle":
		i = EncyclopediaArticle
	case "film":
		i = Film
	case "forumPost":
		i = ForumPost
	case "hearing":
		i = Hearing
	case "instantMessage":
		i = InstantMessage
	case "interview":
		i = Interview
	case "journalArticle":
		i = JournalArticle
	case "letter":
		i = Letter
	case "magazineArticle":
		i = MagazineArticle
	case "manuscript":
		i = Manuscript
	case "map":
		i = Map
	case "newspaperArticle":
		i = NewspaperArticle
	case "note":
		i = Note
	case "patent":
		i = Patent
	case "podcast":
		i = Podcast
	case "presentation":
		i = Presentation
	case "radioBroadcast":
		i = RadioBroadcast
	case "report":
		i = Report
	case "statute":
		i = Statute
	case "tvBroadcast":
		i = TVBroadcast
	case "thesis":
		i = Thesis
	case "videoRecording":
		i = VideoRecording
	case "webpage":
		i = Webpage
	default:
		return ItemType{i}, errors.New("unknown item type \"" + s + "\"")
	}
	return ItemType{i}, nil
}

// Converts an `ItemType` to a string `itemType` consumable by the Zotero web
// API.
func (it *ItemType) ApiStringOfItemType() string {
	switch it.i {
	case Artwork:
		return "artwork"
	case AudioRecording:
		return "audioRecording"
	case Bill:
		return "bill"
	case BlogPost:
		return "blogPost"
	case Book:
		return "book"
	case BookSection:
		return "bookSection"
	case Case:
		return "case"
	case ComputerProgram:
		return "computerProgram"
	case ConferencePaper:
		return "conferencePaper"
	case DictionaryEntry:
		return "dictionaryEntry"
	case Document:
		return "document"
	case Email:
		return "email"
	case EncyclopediaArticle:
		return "encyclopediaArticle"
	case Film:
		return "film"
	case ForumPost:
		return "forumPost"
	case Hearing:
		return "hearing"
	case InstantMessage:
		return "instantMessage"
	case Interview:
		return "interview"
	case JournalArticle:
		return "journalArticle"
	case Letter:
		return "letter"
	case MagazineArticle:
		return "magazineArticle"
	case Manuscript:
		return "manuscript"
	case Map:
		return "map"
	case NewspaperArticle:
		return "newspaperArticle"
	case Note:
		return "note"
	case Patent:
		return "patent"
	case Podcast:
		return "podcast"
	case Presentation:
		return "presentation"
	case RadioBroadcast:
		return "radioBroadcast"
	case Report:
		return "report"
	case Statute:
		return "statute"
	case TVBroadcast:
		return "tvBroadcast"
	case Thesis:
		return "thesis"
	case VideoRecording:
		return "videoRecording"
	case Webpage:
		return "webpage"
	default:
		panic("unreachable")
	}
}

// Converts an `ItemType` to an english-localized string.
func (it *ItemType) EnglishStringOfItemType() string {
	switch it.i {
	case Artwork:
		return "Artwork"
	case AudioRecording:
		return "Audio Recording"
	case Bill:
		return "Bill"
	case BlogPost:
		return "Blog Post"
	case Book:
		return "Book"
	case BookSection:
		return "Book Section"
	case Case:
		return "Case"
	case ComputerProgram:
		return "Computer Program"
	case ConferencePaper:
		return "Conference Paper"
	case DictionaryEntry:
		return "Dictionary Entry"
	case Document:
		return "Document"
	case Email:
		return "E-mail"
	case EncyclopediaArticle:
		return "Encyclopedia Article"
	case Film:
		return "Film"
	case ForumPost:
		return "Forum Post"
	case Hearing:
		return "Hearing"
	case InstantMessage:
		return "Instant Message"
	case Interview:
		return "Interview"
	case JournalArticle:
		return "Journal Article"
	case Letter:
		return "Letter"
	case MagazineArticle:
		return "Magazine Article"
	case Manuscript:
		return "Manuscript"
	case Map:
		return "Map"
	case NewspaperArticle:
		return "Newspaper Article"
	case Note:
		return "Note"
	case Patent:
		return "Patent"
	case Podcast:
		return "Podcast"
	case Presentation:
		return "Presentation"
	case RadioBroadcast:
		return "Radio Broadcast"
	case Report:
		return "Report"
	case Statute:
		return "Statute"
	case TVBroadcast:
		return "TV Broadcast"
	case Thesis:
		return "Thesis"
	case VideoRecording:
		return "Video Recording"
	case Webpage:
		return "Web Page"
	default:
		panic("unreachable")
	}
}

// Checks an `ItemType` for a certain type.
func (it *ItemType) Is(what int) bool {
	return it.i == what
}
