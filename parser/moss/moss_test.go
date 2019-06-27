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
		lines int
	}{
		{
			0,
			"output/9633094_1/final_project (75%)",
			"output/9633094_1/final_project (75%)",
			1001,
		}, {
			130,
			"output/9731100_1/final_project2/cmake-build-debug/CMakeFiles/ (89%)",
			"output/9731016_1/final_project/cmake-build-debug/CMakeFiles/ (89%)",
			33,
		},
	}

	for _, c := range cases {
		match := matches[c.index]
		if match.File1 != c.file1 {
			t.Fatalf("")
		}
		if match.File2 != c.file2 {
			t.Fatalf("")
		}
		if match.LinesMatch != c.lines {
			t.Fatalf("")
		}
	}
}
