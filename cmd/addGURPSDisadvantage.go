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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	ac "github.com/natfarleydev/ankiutils/ankiconnect"
	"github.com/natfarleydev/ankiutils/textutils"
	"github.com/spf13/cobra"
)

// addGURPSDisadvantageCmd represents the addGURPSDisadvantage command
var addGURPSDisadvantageCmd = &cobra.Command{
	Use:   "addGURPSDisadvantage",
	Short: "Sets up an Add Card dialog to add a GURPS Disadvantage",
	Run: func(cmd *cobra.Command, args []string) {
		var front, back string
		if len(args) == 2 {
			front = args[0]
			back = args[1]
		} else if len(args) == 0 {
			fmt.Println("No arguments detected, getting values from STDIN.")
			fmt.Print("Enter Front__pk (\\n\\n^D to finish input): ")
			frontBytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf("Failed to parse input for Front__pk: %q", err)
			}
			front = string(frontBytes)

			fmt.Print("\nEnter Back__pk (\n\n^D to finish input): ")
			backBytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf("Failed to parse input for Back__pk: %q", err)
			}
			back = string(backBytes)
		} else if len(args) != 2 {
			fmt.Printf("Expected two arguments: front of card, back of card; Got: %v\n", args)
			os.Exit(1)
		}

		_, err := ac.GuiAddCards(ac.GuiAddCardsParams{
			Note: ac.GuiAddCardsNote{
				DeckName:  "General::GURPS",
				ModelName: "Context Basic (optional reversed card)",
				Options: map[string]bool{
					"closeAfterAdding": true,
				},
				Fields: map[string]string{
					"Context":     `<img src="gurps-basic-set-cover-image.png"><br><div class="context-alt-text">GURPS</div><div>Disadvantages</div>`,
					"Front__pk":   textutils.FromGURPSBookToString(front),
					"Back__pk":    textutils.FromGURPSBookToString(back),
					"Add Reverse": "y",
				},
			},
		})
		if err != nil {
			log.Fatalf("Problem with guiAddCards: %q", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addGURPSDisadvantageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addGURPSDisadvantageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addGURPSDisadvantageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
