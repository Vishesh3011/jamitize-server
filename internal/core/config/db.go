package config

import (
	"fmt"
	"syscall"
)

type dbConfig struct {
	host     string `json:"host"`
	port     string `json:"port"`
	username string `json:"username"`
	password string `json:"password"`
	database string `json:"database"`
}

func newDBConfig() (*dbConfig, error) {
	host, found := syscall.Getenv("MONGODB_HOST")
	if !found {
		return nil, fmt.Errorf("environment variable MONGODB_HOST not set")
	}

	port, found := syscall.Getenv("MONGODB_PORT")
	if !found {
		return nil, fmt.Errorf("environment variable MONGODB_PORT not set")
	}

	username, found := syscall.Getenv("MONGODB_USERNAME")
	if !found {
		return nil, fmt.Errorf("environment variable MONGODB_USERNAME not set")
	}

	password, found := syscall.Getenv("MONGODB_PASSWORD")
	if !found {
		return nil, fmt.Errorf("environment variable MONGODB_PASSWORD not set")
	}

	database, found := syscall.Getenv("MONGODB_DATABASE")
	if !found {
		return nil, fmt.Errorf("environment variable MONGODB_DATABASE not set")
	}

	return &dbConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}, nil
}

func (d *dbConfig) Host() string {
	return d.host
}
func (d *dbConfig) Port() string {
	return d.port
}

func (d *dbConfig) Username() string {
	return d.username
}

func (d *dbConfig) Password() string {
	return d.password
}

func (d *dbConfig) Database() string {
	return d.database
}
