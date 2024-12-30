package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

type AuthorPaper struct {
	author string
	paper  string
}
type MostCitations struct {
	citationsPerAuthor map[string]Citations //Each paper for a single author with its number of citations
	author             string
	paper              string
	citations          int
}
type Citations struct { //Map of authors and Citations with a notation of the paper with the most Citations
	citations map[string]int //Each paper with its number of citations
}

// Determine in Big O(N) the maximum number of Citations for authors.  It must be how group-by is implemented in a database.
func TestCountCitations(t *testing.T) {

	f := func(mostCitations MostCitations, citation AuthorPaper) MostCitations {
		c, ok := mostCitations.citationsPerAuthor[citation.author]
		if !ok { //author not yet in MostCitations map
			e := map[string]int{citation.paper: 1}

			mostCitations.citationsPerAuthor[citation.author] = Citations{
				citations: e,
			}
		} else { //Found author in MostCitations map
			d, ok := c.citations[citation.paper]
			//Lookup if there are already Citations for this author's paper
			if !ok { //There are none so make a new hashtable entry
				c.citations[citation.paper] = 1
				mostCitations.citationsPerAuthor[citation.author] = c
			} else { //There is already a citation for this author's paper so increment it
				c.citations[citation.paper] = d + 1
				mostCitations.citationsPerAuthor[citation.author] = c
			}
		}
		currentAuthorPaper, _ := mostCitations.citationsPerAuthor[citation.author]
		currentAuthorCitations, _ := currentAuthorPaper.citations[citation.paper]
		if currentAuthorCitations > mostCitations.citations {
			mostCitations.citations = currentAuthorCitations
			mostCitations.paper = citation.paper
			mostCitations.author = citation.author
		}
		return mostCitations
	}

	allCitations := []AuthorPaper{{"fred", "covid1"}, {"fred", "covid2"}, {"fred", "covid3"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid5"}, {"fred", "covid1"}, {"fred", "covid2"}, {"fred", "covid3"},
		{"fred", "covid1"}, {"fred", "covid2"}, {"fred1", "covid3"}, {"fred1", "covid1"},
		{"fred1", "covid3"}, {"fred1", "covid1"}, {"fred1", "covid3"}, {"fred1", "covid1"},
		{"fred1", "covid3"}, {"fred1", "covid3"}, {"fred1", "covid3"}, {"fred1", "covid1"},
		{"fred1", "covid2"}, {"fred", "covid3"},
		{"fred", "covid4"}, {"fred1", "covid3"}, {"fred1", "covid1"}, {"fred1", "covid3"},
		{"fred1", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"},
		{"fred", "covid4"}, {"fred", "covid1"}, {"fred", "covid1"}, {"fred", "covid1"}}

	c := make(map[string]Citations, 0)
	citationsPerAuthor := MostCitations{c, "", "", 0} //make(map[string]Citations, 0)

	actual := arrays.FoldLeft(allCitations, citationsPerAuthor, f)
	if !(actual.citations == 66 && actual.author == "fred" && actual.paper == "covid1") {
		t.Errorf("Expected Author: fred, Paper: covid1, citations: 66 but was Author: %v, Paper: %v, citations: %v ", actual.author, actual.paper, actual.citations)
	}
}
