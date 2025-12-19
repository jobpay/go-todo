package valueobject

type Status int

const (
	// 未完了状態
	StatusPending Status = 0
	// 完了状態
	StatusCompleted Status = 1
)

func (s Status) Bool() bool {
	return s == StatusCompleted
}

func FromBool(completed bool) Status {
	if completed {
		return StatusCompleted
	}
	return StatusPending
}

func (s Status) String() string {
	switch s {
	case StatusCompleted:
		return "completed"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}
