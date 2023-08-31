package sqlTracker

import (
	"avitoTestTask/internal/app/tracker"
	"database/sql"
	_ "github.com/lib/pq"
)

// Tracker ...
type Tracker struct {
	db                    *sql.DB
	segmentRepository     *SegmentRepository
	userSegmentRepository *UserSegmentRepository
}

func New(db *sql.DB) *Tracker {
	return &Tracker{
		db: db,
	}
}

func (t *Tracker) Segment() tracker.SegmentRepository {
	if t.segmentRepository != nil {
		return t.segmentRepository
	}
	t.segmentRepository = &SegmentRepository{
		tracker: t,
	}
	return t.segmentRepository
}

func (t *Tracker) UserSegment() tracker.UserSegmentRepository {
	if t.segmentRepository != nil {
		return t.userSegmentRepository
	}
	t.userSegmentRepository = &UserSegmentRepository{
		tracker: t,
	}
	return t.userSegmentRepository
}
