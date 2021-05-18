package config

type Mode string

const (
	// ModeDebug indicates app mode is debug.
	ModeDebug Mode = "debug"
	// ModeRelease indicates app mode is release.
	ModeRelease Mode = "release"
)

func IsModeDebug(m Mode) bool {
	return m == ModeDebug || m == ""
}

func IsRelease(m Mode) bool {
	return m == ModeRelease
}

const (
	// SysName the name of the aplication that is shown to the users
	//
	// Can be used in emails, pdf, etc.
	SysName = "Phoenix App"

	// SysKey code of the aplication that can be used bys the system itself.
	//
	// Can be user in output directory paths, file names, cache sufixes, etc.
	SysKey = "phoenix-server"

	// SysPkgName go root package name
	SysPkgName = "github.com/jtorz/phoenix-backend"
)

type EnvName string

const (
	EnvPrefix string = "PHOENIX"
	// SysPath root directory of the application..
	EnvSysPath EnvName = "PATH"
)
