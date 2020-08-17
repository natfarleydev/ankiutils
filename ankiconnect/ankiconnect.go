package ankiconnect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL of AnkiConnect
var URL = "http://localhost:8765"

// GuiAddCardsNote defines a note for use with guiAddCards
type GuiAddCardsNote struct {
	DeckName  string            `json:"deckName"`
	ModelName string            `json:"modelName"`
	Options   map[string]bool   `json:"options,omitempty"`
	Fields    map[string]string `json:"fields,omitempty"`
}

// GuiAddCardsParams is the parameters for the payload for the guiAddCards endpoint on AnkiConnect
type GuiAddCardsParams struct {
	Note GuiAddCardsNote `json:"note"`
}

// GuiAddCardsPayload is the payload for the guiAddCards endpoint on AnkiConnect
type GuiAddCardsPayload struct {
	Action  string            `json:"action"`
	Version int               `json:"version"`
	Params  GuiAddCardsParams `json:"params"`
}

// Response is a standard response from AnkiConnect
type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

// GuiAddCards calls the guiaddcards action on AnkiConnect
func GuiAddCards(params GuiAddCardsParams) (Response, error) {

	payload, err := json.Marshal(GuiAddCardsPayload{
		Action:  "guiAddCards",
		Version: 6,
		Params:  params,
	})
	if err != nil {
		return Response{}, fmt.Errorf("Unable to marshal payload: %q", err)
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return Response{}, fmt.Errorf("Got a problem with the request to AnkiConnect: %q", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to read JSON body of reply: %q", err)
	}

	var ankiConnectResp Response
	err = json.Unmarshal(body, &ankiConnectResp)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to parse responsen from AnkiConnect: %q", err)
	}
	if ankiConnectResp.Error != "" {
		return ankiConnectResp, fmt.Errorf("Error, AnkiConnect returned following error: %q", ankiConnectResp.Error)
	}

	return ankiConnectResp, nil

}
