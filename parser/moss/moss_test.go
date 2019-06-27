package moss_test

import (
	"os"
	"testing"

	"github.com/aut-ceit/mosser/parser/moss"
)

func TestExtractMatches(t *testing.T) {
	f, err := os.Open("../.ci/moss.html")
	if err != nil {
		t.Fatal(err)
	}

	matches, err := moss.ExtractMatches(f)
	if err != nil {
		t.Fatal(err)
	}

	// check the total number of matches
	if len(matches) != 250 {
		t.Fatalf("some matches are missing 250 != %d", len(matches))
	}

	// check some specific match
	cases := []struct {
		index int
		file1 string
		file2 string
		link  string
		lines int
	}{
		{
			0,
			"output/9633094_1/final_project (75%)",
			"output/9633094_1/final_project (75%)",
			"http://moss.stanford.edu/results/623019542/match0.html",
			1001,
		}, {
			130,
			"output/9731100_1/final_project2/cmake-build-debug/CMakeFiles/ (89%)",
			"output/9731016_1/final_project/cmake-build-debug/CMakeFiles/ (89%)",
			"http://moss.stanford.edu/results/623019542/match130.html",
			33,
		},
	}

	for _, c := range cases {
		match := matches[c.index]
		if match.File1 != c.file1 {
			t.Fatalf("record [%d] has missmatched on file1 %s != %s", c.index, c.file1, match.File1)
		}
		if match.File2 != c.file2 {
			t.Fatalf("record [%d] has missmatched on file2 %s != %s", c.index, c.file2, match.File2)
		}
		if match.Link != c.link {
			t.Fatalf("record [%d] has missmatched on link %s != %s", c.index, c.link, match.Link)
		}
		if match.LinesMatch != c.lines {
			t.Fatalf("record [%d] has missmatched on lines matched %d != %d", c.index, c.lines, match.LinesMatch)
		}
	}
}
