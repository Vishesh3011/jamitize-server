package config

import (
	"errors"
	"syscall"
)

type jwtConfig struct {
	secret string
}

func newJWTConfig() (*jwtConfig, error) {
	secret, found := syscall.Getenv("JWT_SECRET")
	if !found {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	return &jwtConfig{secret: secret}, nil
}

func (c *jwtConfig) Secret() string {
	return c.secret
}
