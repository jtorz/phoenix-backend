package base

// RecordActions acttion that can be performed with the record.
type RecordActions []string

// Simple actions that can be performed with a record.
func (a *RecordActions) SimpleActions(s Status) {
	*a = NewRecordActionsSimple(s)
}

// NewRecordActionsSimple simple actions that can be performed with a record.
func NewRecordActionsSimple(s Status) RecordActions {
	switch s {
	case StatusCaptured:
		return []string{"edit", "delete", "validate"}
	case StatusActive:
		return []string{"aproved", "edit", "invalidate"}
	case StatusInactive:
		return []string{"disaproved", "edit", "validate"}
	default:
		return []string{}
	}
}
