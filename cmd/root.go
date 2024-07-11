/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"

	"github.com/spf13/cobra"
)

//go:embed quotes.json
var fileContent []byte

type Quote struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terry",
	Short: "Get quotes from computer genius Terry A. Davis",
	Long:  `Get real quotes from computer genius and creator of TempleOS Terry A. Davis.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		cat, _ := cmd.Flags().GetString("category")
		big, _ := cmd.Flags().GetBool("big")

		byteResult := fileContent

		var quotes []Quote
		err := json.Unmarshal(byteResult, &quotes)
		if err != nil {
			log.Fatalln(err)
		}

		if cat != "" && !catValid(cat) {
			fmt.Println("Category:", cat, "doesn't exist")
			return
		}

		if catValid(cat) {
			var filtered []Quote
			for _, quote := range quotes {
				if quote.Category == cat {
					filtered = append(filtered, quote)
				}
			}
			quotes = filtered
		}

		if len(quotes) == 0 {
			fmt.Println("No quotes found")
			return
		}

		index := rand.Int() % len(quotes)

		fmt.Print("\"")
		fmt.Print(quotes[index].Text)
		fmt.Println("\"")

		if big == false {
			fmt.Println("― Terry A. Davis")
			return
		}
		fmt.Println(beautify())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("category", "c", "", "Filter quotes by category")
	rootCmd.Flags().BoolP("big", "b", false, "Write Terry A. Davis' Name in big ASCII art")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terry.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize()
}

// Helper function:
func catValid(cat string) bool {
	if cat == "crazy" || cat == "cia" || cat == "confused" || cat == "racist" || cat == "deep" || cat == "technical" {
		return true
	}
	return false
}

func beautify() string {
	return "  _______                                  _____              _     \n |__   __|                       /\\       |  __ \\            (_)    \n    | | ___ _ __ _ __ _   _     /  \\      | |  | | __ ___   ___ ___ \n    | |/ _ \\ '__| '__| | | |   / /\\ \\     | |  | |/ _` \\ \\ / / / __|\n    | |  __/ |  | |  | |_| |  / ____ \\ _  | |__| | (_| |\\ V /| \\__ \\\n    |_|\\___|_|  |_|   \\__, | /_/    \\_(_) |_____/ \\__,_| \\_/ |_|___/\n                       __/ |                                        \n                      |___/                                         "
}
