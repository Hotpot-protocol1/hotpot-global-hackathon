package config

import "time"

type HTTP struct {
	Port           string        `mapstructure:"API_PORT"`
	Host           string        `mapstructure:"API_HOST"`
	SSL            bool          `mapstructure:"API_SSL"`
	ServerCertPath string        `mapstructure:"API_CERT_PATH"`
	ServerKeyPath  string        `mapstructure:"API_CERT_KEY"`
	ReadTimeout    time.Duration `mapstructure:"API_READ_TIMEOUT"`
	WriteTimeout   time.Duration `mapstructure:"API_WRITE_TIMEOUT"`
	IdleTimeout    time.Duration `mapstructure:"API_IDLE_TIMEOUT"`
}
