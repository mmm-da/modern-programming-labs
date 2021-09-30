package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type TxtInput struct {
	metaStringList []string
	text           string
}

func (i TxtInput) GetText() string {
	return i.text
}

func (i TxtInput) GetTitle() string {
	return i.metaStringList[0]
}

func (i TxtInput) GetAuthors() []string {
	authorListString := i.metaStringList[1]
	return strings.Split(authorListString, ";")
}

func (i TxtInput) GetHash() string {
	return i.metaStringList[2]
}

func (i *TxtInput) ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	metaString, err := reader.ReadString(byte('\n'))
	if err != nil {
		log.Fatal(err)
	}

	i.metaStringList = strings.Split(metaString, "/")

	text, err := reader.ReadString(byte(0))
	i.text = text

	file.Close()
}
