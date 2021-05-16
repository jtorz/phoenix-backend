package syslog

// Logger writes to the system log.
type Logger interface {
	Errorf(format string, a ...interface{}) error
	Warningf(format string, a ...interface{}) error
	Infof(format string, a ...interface{}) error
}
