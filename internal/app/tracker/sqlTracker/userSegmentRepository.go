package sqlTracker

import (
	"avitoTestTask/internal/app/model"
	"avitoTestTask/internal/app/tracker"
	"database/sql"
)

// Здесь пишем взаимодействие с базой данных
type UserSegmentRepository struct {
	tracker *Tracker
}

func (r *UserSegmentRepository) Create(u *model.UserSegment) error {
	return r.tracker.db.QueryRow(
		"INSERT INTO users_segments (user_id, segment_slug) VALUES ($1, $2) RETURNING created_at",
		u.UserId,
		u.SegmentSlug,
	).Scan(&u.CreatedAt)
}

func (r *UserSegmentRepository) Delete(id int, slug string) error {
	return r.tracker.db.QueryRow(
		"DELETE FROM users_segments WHERE user_id = $1 AND segment_slug = $2",
		id,
		slug,
	).Scan()
}

func (r *UserSegmentRepository) DeleteBySlug(slug string) error {
	return r.tracker.db.QueryRow(
		"DELETE FROM users_segments WHERE segment_slug = $1",
		slug,
	).Scan()
}

func (r *UserSegmentRepository) FindByUserId(id int) ([]model.Segment, error) {
	segments := make([]model.Segment, 0)
	rows, err := r.tracker.db.Query(
		"SELECT segment_slug FROM users_segments WHERE user_id = $1",
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, tracker.ErrorRecordNotFound
		}

		return nil, err
	}
	for rows.Next() {
		var slug string
		err := rows.Scan(&slug)
		if err != nil {
			return nil, err
		}
		segment, err := r.tracker.Segment().FindBySlug(slug)
		if err != nil {
			return nil, err
		}
		segments = append(segments, *segment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}
