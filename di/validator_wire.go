//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bps-kota-bontang/fiber-starter/validator"
	"github.com/google/wire"
)

// Wire Set for Database
var ValidatorSet = wire.NewSet(validator.ProvideValidator)
