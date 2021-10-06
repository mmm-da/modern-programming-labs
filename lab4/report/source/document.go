package document

import (
	"crypto/md5"
	"fmt"
)

type Document struct {
	Title, Text string
	Authors     []string
	Hash        [16]byte
}

type builder interface {
	reset()
	SetBody(string)
	SetTitle(string)
	AddAuthor(string)
	SetHash([16]byte)
	CheckHash() bool
	GetDocument() Document
}

// 1 строка заголовок/авторы через запятую/хеш
// все остальное тело статьи
type Builder struct {
	result Document
}

func (b *Builder) SetText(text string) {
	b.result.Text = text
}

func (b *Builder) SetTitle(title string) {
	b.result.Title = title
}

func (b *Builder) AddAuthor(author string) {
	b.result.Authors = append(b.result.Authors, author)
}

func (b *Builder) SetHash(textHash string) {
	hash := []byte(textHash)
	copy(b.result.Hash[:], hash)
}

func (b *Builder) CheckHash() bool {
	hashSum := md5.Sum([]byte((b.result.Text)))
	return hashSum == b.result.Hash
}

func (b *Builder) Reset() {
	b.result = Document{}
}

func (b Builder) GetDocument() Document {
	if b.CheckHash() {
		fmt.Println("Document text is correct")
	} else {
		fmt.Println("Document text is incorrect")
	}
	return b.result
}
