package main

import (
	"fmt"
	"sort"
	"strings"
)

var original = []string{
	"Nonmetals",
    "    Hydrogen",
    "    Carbon",
    "    Nitrogen",
    "    Oxygen",
    "Inner Transitionals",
    "    Lanthanides",
    "        Europium",
    "        Cerium",
    "    Actinides",
    "        Uranium",
    "        Plutonium",
    "        Curium",
    "Alkali Metals",
    "    Lithium",
    "    Sodium",
    "    Potassium",
}

func main() {
	fmt.Println("|     Original      |       Sorted      |")
    fmt.Println("|-------------------|-------------------|")
    sorted := SortedIndentedStrings(original)
    for i := range original {
    	fmt.Printf("|%-19s|%-19s|\n", original[i], sorted[i])
    }
}

type Entry struct {
	key string
	value string
	children Entries
}

type Entries []Entry

func SortedIndentedStrings(slice []string) []string {
	entries := populateEntries(slice)

}

func populateEntries(slice []string) Entries {
	indent, indentSize := computeIndent(slice)

	entries := make(Entries, 0)
}

func computeIndent(slice []string) (string, int) {
	for _, item := range slice {
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whitespace := rune(item[0])
			for i, char := range item[1:] {
				if char != whitespace {
					i ++
					return strings.Repeat(string(whitespace), i), i
				}
			}
		}
	}
	return "", 0
}

