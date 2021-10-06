package output

import (
	"encoding/xml"
	"lab4/pkg/document"
	"os"
)

type XmlOutput struct{}

func (o XmlOutput) SaveFile(doc document.Document, filename string) {
	type XmlDocument struct {
		XMLName xml.Name `xml:"document"`
		Title   string   `xml:"title"`
		Text    string   `xml:"text"`
		Authors []string `xml:"author"`
		Hash    [16]byte `xml:"hash"`
	}

	xmlDoc := XmlDocument{Title: doc.Title, Text: doc.Text, Authors: doc.Authors, Hash: doc.Hash}
	xmlOut, _ := xml.MarshalIndent(xmlDoc, " ", "  ")
	xmlString := string(xmlOut)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	file.WriteString(xmlString)
	file.Close()
}
