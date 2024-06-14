package resolver

import (
	"github.com/99designs/gqlgen/codegen/config"
	"pet-pamper/handler"
	//_ "pet-pamper/cmd/types"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	config      config.Config
	authHandler handler.IAuthHandler
}

func (r *Resolver) FromHandlerDeps(deps int) {

}

func NewResolver() *Resolver {
	authHandler := handler.NewAuthHandler()
	return &Resolver{
		config:      config.Config{},
		authHandler: authHandler,
	}
}
