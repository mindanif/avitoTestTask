package model

import "time"

type Segment struct {
	Id        int        `json:"id"`
	Slug      string     `json:"slug"`
	CreatedAt *time.Time `json:"created_at"`
}
