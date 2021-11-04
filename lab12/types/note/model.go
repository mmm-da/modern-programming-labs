package note

import (
	"time"

	"github.com/brianvoe/gofakeit"
	lorem "github.com/drhodes/golorem"
)

type Note struct {
	Id           int       `json:"id"`
	AuthorId     int       `json:"author_id"`
	CreateTime   time.Time `json:"create_datetime"`
	LastEditTime time.Time `json:"lastedit_datetime"`
	Title        string    `json:"title"`
	Text         string    `json:"text"`
	Tags         []string  `json:"tags"`
}

func CreateMock() Note {
	id := gofakeit.Int32()
	author := gofakeit.Int32()
	createTime := gofakeit.Date()
	lastEditTime := gofakeit.Date()
	title := gofakeit.HackerPhrase()
	text := lorem.Sentence(10, 100)
	tags := make([]string, 0)
	for i := 0; i < 5; i++ {
		tags = append(tags, gofakeit.BuzzWord())
	}
	result := Note{Id: int(id), AuthorId: int(author), CreateTime: createTime, LastEditTime: lastEditTime, Title: title, Text: text, Tags: tags}
	return result
}

func CreateMockArray(n int) []Note {
	result := make([]Note, 0)
	for i := 0; i < n; i++ {
		result = append(result, CreateMock())
	}
	return result
}
