package config

type Mode string

const (
	// ModeDebug indicates app mode is debug.
	ModeDebug Mode = "debug"
	// ModeRelease indicates app mode is release.
	ModeRelease Mode = "release"
)

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

	// ConfFile path to the configuration file.
	//
	// The path is relative to SysPath()
	ConfFile = "assets/config/config.json"

	// ConfSchemaFile json schema file used to validate the configuration file.
	//
	// The path is relative to SysPath()
	ConfSchemaFile = "assets/config/config-schema.json"
)

type EnvName string

const (
	EnvPrefix string = "PHOENIX"
	// SysPath root directory of the application..
	EnvSysPath EnvName = "PATH"
	// AppMode defines how the aplication is deployed.
	EnvAppMode EnvName = "MODE"
	// JWTKey key to validate the jwt.
	EnvJWTKey EnvName = "JWT_KEY"
	// CryptKey encryption key
	EnvCryptKey EnvName = "CRYPT_KEY"
)
