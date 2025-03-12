//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

// Initialize App
func InitializeApp() (*AppContainer, error) {
	wire.Build(AppSet)
	return &AppContainer{}, nil
}
