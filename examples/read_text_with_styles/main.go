package main

import (
	"fmt"

	"github.com/aruncs31s/pdf"
)

func main() {
	f, r, err := pdf.Open("./dms.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sentences, err := r.GetStyledTexts()
	if err != nil {
		panic(err)
	}

	// Print all sentences
	for _, sentence := range sentences {
		bold := ""
		if sentence.IsBold {
			bold = " [BOLD]"
		}
		fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s , Is Bold: %s \n",
			sentence.Font,
			sentence.FontSize,
			sentence.X,
			sentence.Y,
			sentence.S,
			func() string {
				if bold == "" {
					return "false"
				}
				return "true"
			}())
	}
}
