package repository

import (
	"github.com/triardn/golang-base-project/internal/app/commons"
	"gopkg.in/gorp.v2"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
	DB *gorp.DbMap
}

// Repository all repo object injected here
type Repository struct {
	Cache  ICacheRepository
	Person IPersonRepository
}
