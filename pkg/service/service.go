package service

import (
	"log"

	"github.com/borud/tt/pkg/auth"
	"github.com/borud/tt/pkg/store"
	ttv1 "github.com/borud/tt/pkg/tt/v1"
)

// Service implements the gRPC service.
type Service interface {
	ttv1.TTServiceServer
}

// Config for service.
type Config struct {
	DB  store.Store
	JWT *auth.JWT
}

type service struct {
	config Config
}

// New service.
func New(c Config) Service {
	if c.DB == nil {
		log.Fatalf("db was nil")
	}

	if c.JWT == nil {
		log.Fatal("JWT instance was nil")
	}

	return &service{
		config: c,
	}
}
