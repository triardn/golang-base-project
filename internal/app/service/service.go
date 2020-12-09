package service

import (
	"github.com/triardn/golang-base-project/config"
	"github.com/triardn/golang-base-project/internal/app/commons"
	"github.com/triardn/golang-base-project/internal/app/repository"
)

// Option anything any service object needed
type Option struct {
	config.AppConfig
	commons.Options
	*repository.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheckService
}
