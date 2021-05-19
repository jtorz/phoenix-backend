package mail

type ServiceConfig struct {
	Host           string `mapstructure:"MAIL_HOST"`
	Port           string `mapstructure:"MAIL_PORT"`
	Username       string `mapstructure:"MAIL_USERNAME"`
	Password       string `mapstructure:"MAIL_PASSWORD"`
	Encryption     string `mapstructure:"MAIL_ENCRYPTION"`
	ConnectTimeout int    `mapstructure:"MAIL_CONNECT_TIMEOUT"`
	SendTimeout    int    `mapstructure:"MAIL_SEND_TIMEOUT"`
}

func (svc ServiceConfig) SetDefaults(setDefault func(key string, v interface{})) error {
	setDefault("MAIL_CONNECT_TIMEOUT", 10)
	setDefault("MAIL_SEND_TIMEOUT", 10)
	return nil
}
