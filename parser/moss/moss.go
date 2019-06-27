/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 27-06-2019
 * |
 * | File Name:     moss.go
 * +===============================================
 */

package moss

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Match represents a single match in moss result
type Match struct {
	File1      string
	File2      string
	Link       string
	LinesMatch int
}

// ExtractMatches extracts the matches from the moss result html page.
// check .ci/moss.html as an example
func ExtractMatches(input io.Reader) ([]Match, error) {
	var matches []Match

	doc, err := html.Parse(input)
	if err != nil {
		return matches, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			/*
				The following ogly code tries to decode the html with the following structure
				<tr>
				<td>
				<a href="http://moss.stanford.edu/results/623019542/match1.html">output/9731100_1/final_project2/cmake-build-debug/CMakeFiles/3.12.2/CompilerIdC/ (70%)</a>
				</td>
				<td>
				<a href="http://moss.stanford.edu/results/623019542/match1.html">output/9731016_1/final_project/cmake-build-debug/CMakeFiles/3.13.2/CompilerIdC/ (70%)</a>
				</td>
				<td align="right">297</td>
				</tr>
			*/
			td := n.FirstChild
			if td != nil && td.Data == "td" {
				var match Match

				match.File1 = td.FirstChild.FirstChild.Data

				match.Link = td.FirstChild.Attr[0].Val

				match.File2 = td.NextSibling.FirstChild.FirstChild.Data

				var n int
				fmt.Sscanf(td.NextSibling.NextSibling.FirstChild.Data, "%d", &n)
				match.LinesMatch = n

				matches = append(matches, match)

				return // there is no point to continue iteration on td tag
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return matches, nil
}
