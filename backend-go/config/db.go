package config

import (
	"errors"
	"fmt"
)

type DB struct {
	Name     string `mapstructure:"API_DATABASE_NAME"`
	Host     string `mapstructure:"API_DATABASE_HOST"`
	Port     int    `mapstructure:"API_DATABASE_PORT"`
	User     string `mapstructure:"API_DATABASE_USER"`
	Password string `mapstructure:"API_DATABASE_PASSWORD"`
	SSL      string `mapstructure:"API_DATABASE_SSL"`
}

func (d DB) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

func (d *DB) GetConnectionString() (string, error) {
	err := d.validate()
	if err != nil {
		return "", err
	}

	return d.Info(), err
}

func (d DB) validate() error {
	if d.Name == "" {
		return errors.New("database name can't be empty")
	}

	if d.Port == 0 {
		return errors.New("database port can't be empty")
	}

	if d.User == "" {
		return errors.New("database user can't be empty")
	}

	if d.Password == "" {
		return errors.New("database password can't be empty")
	}

	return nil
}
