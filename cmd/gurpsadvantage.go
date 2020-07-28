/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// gurpsadvantageCmd represents the gurpsadvantage command
var gurpsadvantageCmd = &cobra.Command{
	Use:   "gurpsadvantage",
	Short: "Opens an Anki add card dialog prepopulated with a GURPS advantage",
	Run: func(cmd *cobra.Command, args []string) {
		// resp, err := http.Post("localhost:8765")
		// if err != nil {
		// 	log.Fatalf("Something went wrong with the command: %q", err)
		// }
		// log.Printf("%#v", resp)
		// fmt.Println("gurpsadvantage called")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(gurpsadvantageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gurpsadvantageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gurpsadvantageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
