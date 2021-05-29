package base

// RecordActions acttion that can be performed with the record.
type RecordActions []string

// NewRecordActionsCommon common actions that can be performed with a record.
func NewRecordActionsCommon(s Status) RecordActions {
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
