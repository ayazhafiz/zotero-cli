package v3

import (
	. "gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

type testcase struct {
	authors []Author
	fmt     string
}

func (tc *testcase) check(t *testing.T) {
	res := PrintAuthors(tc.authors)
	Assert(t, is.Equal(res, tc.fmt))
}

func tc(fmt string, authors ...Author) *testcase {
	return &testcase{authors, fmt}
}

func TestPrintAuthors(t *testing.T) {
	tc("(unknown)").check(t)
	tc("Doe", Author{"Jane", "Doe"}).check(t)
	tc("Doe and Eve", Author{"Jane", "Doe"}, Author{"Serine", "Eve"}).check(t)
	tc("Doe, Eve, and Foil",
		Author{"Jane", "Doe"}, Author{"Serine", "Eve"}, Author{"Alum", "Foil"}).check(t)
	tc("Doe et al.",
		Author{"Jane", "Doe"}, Author{"Serine", "Eve"}, Author{"Alum", "Foil"},
		Author{"Boris", "Gale"}).check(t)
	tc("Doe et al.",
		Author{"Jane", "Doe"}, Author{"Serine", "Eve"}, Author{"Alum", "Foil"},
		Author{"Boris", "Gale"}, Author{"Hash", "Hope"}).check(t)
}
