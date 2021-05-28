package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	ASPNETSessionIDName = "ASP.NET_SessionId"
	MRHSessionName      = "MRHSession"
)

type Config struct {
	Action          string `validate:"required"`
	AspNetSessionId string `validate:"required"`
	MRHSession      string `validate:"required"`
	EmployeeID      string `validate:"required"`
	PositionID      string `validate:"required"`
	LogLevel        string `validate:"required"`
}

func DefaultConfig() Config {
	return Config{
		LogLevel: "info",
	}
}

func (cfg *Config) Verify() error {
	log.Trace("Validating config")
	validate := validator.New()
	return validate.Struct(cfg)
}
