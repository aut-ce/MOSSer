package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

// Match represents a single match in moss result
type Match struct {
	File1 string
	File2 string
}

// Execute run the root command
func Execute() {
	var cmd = &cobra.Command{
		Use:   "parser [moss result url]",
		Short: "parser for moss results",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			z := html.NewTokenizer(resp.Body)
			var tr bool
			var a bool
			var td int
			var match Match
			var matches []Match
			for {
				tt := z.Next()
				if tt == html.ErrorToken {
					break
				}
				t := z.Token()
				switch tt {
				case html.StartTagToken:
					switch t.Data {
					case "tr":
						tr = true
					case "td":
						if tr {
							td++
							match = Match{}
						}
					case "a":
						if td > 0 {
							a = true
						}
					}
				case html.EndTagToken:
					switch t.Data {
					case "tr":
						tr = false
						td = 0
						matches = append(matches, match)
					case "a":
						if td > 0 {
							a = false
						}
					}
				case html.TextToken:
					if tr && a {
						switch td {
						case 1:
							fmt.Println(t.Data)
							match.File1 = t.Data
						case 2:
							fmt.Println(t.Data)
							match.File2 = t.Data
						}
					}
				}
			}
			fmt.Println(matches)
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
