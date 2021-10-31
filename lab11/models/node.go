package models

import "time"

type Note struct {
	Id               int       `json:"id"`
	AuthorId         int       `json:"author_id"`
	CreateDatetime   time.Time `json:"create_datetime"`
	LastEditDatetime time.Time `json:"lastedit_datetime"`
	Text             string    `json:"text"`
	Tags             []string  `json:"tags"`
}
