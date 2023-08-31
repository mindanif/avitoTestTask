package model

import "time"

type UserSegment struct {
	UserId      int       `json:"user_id"`
	SegmentSlug string    `json:"segment_slug"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"expired_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
