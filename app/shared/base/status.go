package base

// Status logical status of the record on th DB.
type Status byte

const (
	// StatusDroppped logical status for records that shouldn't be considered in any process.
	StatusDroppped Status = iota
	// StatusCaptured recors that are being captured.
	StatusCaptured
	// StatusActive records that have been validated to be used in the process.
	StatusActive
	// StatusInactive record that were once active.
	StatusInactive
)

func (s Status) String() string {
	switch s {
	case StatusDroppped:
		return "Droppped"
	case StatusCaptured:
		return "Captured"
	case StatusActive:
		return "Active"
	case StatusInactive:
		return "Inactive"
	}
	return "Unknown"
}
