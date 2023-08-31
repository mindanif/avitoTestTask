package tracker

type Tracker interface {
	Segment() SegmentRepository
	UserSegment() UserSegmentRepository
}
