package config

type Conf struct {
	HTTP     `mapstructure:",squash"`
	Contract `mapstructure:",squash"`
	Log      `mapstructure:",squash"`
	Infura   `mapstructure:",squash"`
	Webhook  `mapstructure:",squash"`
	DB       `mapstructure:",squash"`
}
