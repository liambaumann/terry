/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

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

		if cat != "" {
			fmt.Println("Category:", cat, "doesn't exist")
			return
		}

		fileContent, err := os.Open("quotes.json")
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer fileContent.Close()

		byteResult, err := io.ReadAll(fileContent)
		if err != nil {
			log.Fatalln(err)
			return
		}

		var quotes []Quote
		err = json.Unmarshal(byteResult, &quotes)
		if err != nil {
			log.Fatalln(err)
		}

		if len(quotes) == 0 {
			fmt.Println("No quotes found")
			return
		}

		fmt.Println(quotes[0].Text)
		fmt.Println("― Terry A. Davis")
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
	rootCmd.Flags().StringP("category", "c", "", "Quote category")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terry.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize()
}
