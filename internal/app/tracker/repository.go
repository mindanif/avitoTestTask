package tracker

import "avitoTestTask/internal/app/model"

type SegmentRepository interface {
	Create(*model.Segment) error
	Delete(segment *model.Segment, slug string) error
	FindBySlug(string) (*model.Segment, error)
}

type UserSegmentRepository interface {
	Create(segment *model.UserSegment) error
	FindByUserId(id int) ([]model.Segment, error)
	Delete(id int, slug string) error
	DeleteBySlug(slug string) error
}
