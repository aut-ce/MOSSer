package cmd

import (
	"fmt"
	"net/http"
	"os"

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
			fmt.Println(matches)
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
