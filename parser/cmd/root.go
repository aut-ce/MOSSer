package cmd

import (
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/aut-ceit/mosser/parser/moss"
	"github.com/spf13/cobra"
)

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
			matches, err := moss.ExtractMatches(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			sort.Slice(matches, func(i, j int) bool { return matches[i].File1 < matches[j].File1 })
			for _, match := range matches {
				fmt.Printf("%s\t%s\t%d\n", match.File1, match.File2, match.LinesMatch)
			}

		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
