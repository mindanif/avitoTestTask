package sqlTracker

import (
	"avitoTestTask/internal/app/model"
	"avitoTestTask/internal/app/tracker"
	"database/sql"
)

type SegmentRepository struct {
	tracker *Tracker
}

//Здесь пишем взаимодействие с базой данных

// Create ...
func (r *SegmentRepository) Create(s *model.Segment) error {
	slug := s.Slug
	return r.tracker.db.QueryRow(
		"INSERT INTO segments (slug) VALUES ($1) RETURNING id, created_at",
		slug,
	).Scan(&s.Id, &s.CreatedAt)
}
func (r *SegmentRepository) Delete(segment *model.Segment, slug string) error {
	return r.tracker.db.QueryRow(
		"DELETE FROM segments WHERE slug = $1 RETURNING id, slug, created_at",
		slug,
	).Scan(&segment.Id, &segment.Slug, &segment.CreatedAt)
}
func (r *SegmentRepository) FindBySlug(slug string) (*model.Segment, error) {
	s := &model.Segment{}
	if err := r.tracker.db.QueryRow(
		"SELECT id, slug, created_at FROM segments WHERE slug = $1",
		slug,
	).Scan(
		&s.Id,
		&s.Slug,
		&s.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, tracker.ErrorRecordNotFound
		}

		return nil, err
	}

	return s, nil
}
