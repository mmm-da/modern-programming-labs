package main

import (
	"lab4/pkg/document"
	"lab4/pkg/input"
	"lab4/pkg/output"
)

func main() {
	fileInput := "text.txt"
	fileOutput := "text.xml"

	textInput := input.TxtInput{}
	textOutput := output.XmlOutput{}

	docBuilder := document.Builder{}

	textInput.ReadFile(fileInput)
	authors := textInput.GetAuthors()
	for _, author := range authors {
		docBuilder.AddAuthor(author)
	}
	docBuilder.SetTitle(textInput.GetTitle())
	docBuilder.SetText(textInput.GetText())
	docBuilder.SetHash(textInput.GetHash())
	doc := docBuilder.GetDocument()

	textOutput.SaveFile(doc, fileOutput)
}
