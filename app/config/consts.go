package config

type Mode string

const (
	// ModeDebug indicates app mode is debug.
	ModeDebug Mode = "debug"
	// ModeRelease indicates app mode is release.
	ModeRelease Mode = "release"
)

// IsModeDebug Checks if the mode is "debug".
func IsModeDebug(m Mode) bool {
	return m == ModeDebug || m == ""
}

// IsModeDebug Checks if the mode is "release".
func IsRelease(m Mode) bool {
	return m == ModeRelease
}

const (
	// SysName the name of the aplication that is shown to the users.
	//
	// Can be used in emails, pdf, etc.
	SysName = "Phoenix App"

	// SysKey code of the aplication that can be used by the system.
	//
	// Can be used in output directory paths, file names, cache sufixes, etc.
	SysKey = "phoenix-server"

	// SysPkgName go root package name.
	SysPkgName = "github.com/jtorz/phoenix-backend"
)

const (
	// EnvPrefix enviroment variables prefix.
	EnvPrefix string = "PHOENIX"
)

type LogginLvl uint8

const (
	// LogDebug indicates logging level is for debug and upper levels.
	LogDebug LogginLvl = 0
	// LogWarning indicates logging level is for warning an upper levels.
	LogWarning LogginLvl = 50
	// LogError indicates logging level is for errors only an upper levels.
	LogError LogginLvl = 100
	// LogMax max logging level.
	LogMax LogginLvl = 255
)
