package note

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	mapset "github.com/deckarep/golang-set"
	lorem "github.com/drhodes/golorem"
)

type Note struct {
	Id           int        `json:"id"`
	AuthorId     int        `json:"author_id"`
	CreateTime   time.Time  `json:"create_datetime"`
	LastEditTime time.Time  `json:"lastedit_datetime"`
	Title        string     `json:"title"`
	Text         string     `json:"text"`
	Tags         mapset.Set `json:"tags"`
}

func CreateMock() Note {
	id := gofakeit.Int32()
	author := gofakeit.Int32()
	createTime := gofakeit.Date()
	lastEditTime := gofakeit.Date()
	title := gofakeit.HackerPhrase()
	text := lorem.Sentence(10, 100)
	tags := mapset.NewSet()
	for i := 0; i < 5; i++ {
		tags.Add(gofakeit.BuzzWord())
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

func FilterNotesByTags(tags mapset.Set, notes []Note) []Note {
	if len(tags.ToSlice()) == 1 {
		return notes
	}
	result := make([]Note, 0)
	fmt.Println(tags)
	for _, item := range notes {
		fmt.Println(item.Tags)
		if tags.IsSubset(item.Tags) {
			result = append(result, item)
		}
	}
	return result
}
