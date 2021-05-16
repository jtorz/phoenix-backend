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
		return "Status Droppped"
	case StatusCaptured:
		return "Status Captured"
	case StatusActive:
		return "Status Active"
	case StatusInactive:
		return "Status Inactive"
	}
	return "Status Unknown"
}
