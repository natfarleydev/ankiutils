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
	"log"

	ac "github.com/natfarleydev/ankiutils/ankiconnect"
	"github.com/spf13/cobra"
)

// addAnkiCardCmd represents the addAnkiCard command
var addAnkiCardCmd = &cobra.Command{
	Use:   "addAnkiCard",
	Short: "Brings up the Add Card dialog from a currently open Anki session",
	Run: func(cmd *cobra.Command, args []string) {

		_, err := ac.GuiAddCards(ac.GuiAddCardsParams{
			Note: ac.GuiAddCardsNote{
				DeckName:  "General::Random",
				ModelName: "Context Basic (optional reversed card)",
				Options: map[string]bool{
					"closeAfterAdding": true,
				},
			},
		})
		if err != nil {
			log.Fatalf("Problem with guiAddCards: %q", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addAnkiCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addAnkiCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addAnkiCardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
