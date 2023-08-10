package config

type Infura struct {
	BaseURL   string `mapstructure:"API_INFURA_BASE_URL"`
	BaseURLWS string `mapstructure:"API_INFURA_BASE_URL_WS"`
	APIKey    string `mapstructure:"API_INFURA_KEY"`
}
