/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand/v2"
	"strconv"
)

// timezoneCmd represents the timezone command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		nFlag, _ := cmd.Flags().GetString("num")
		randomInt := -1
		if nFlag != "" {
			n, err := strconv.Atoi(nFlag)
			if err != nil {
				panic(err)
			}
			randomInt = rand.IntN(n)
		} else {
			randomInt = rand.IntN(100)
		}
		fmt.Println(randomInt)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().StringP("num", "n", "", "Range size")
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
