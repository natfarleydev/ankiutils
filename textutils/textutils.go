package textutils

import "strings"

// FromGURPSBookToString transforms text copied from a GURPS book into a continuous string.
//
// NOTE: does not intentional line breaks like paragraphs.
func FromGURPSBookToString(GURPSBookText string) string {
	GURPSBookText = strings.TrimSpace(GURPSBookText)
	GURPSBookText = strings.ReplaceAll(GURPSBookText, "-\n", "")
	GURPSBookText = strings.ReplaceAll(GURPSBookText, "\n", " ")

	return GURPSBookText
}
