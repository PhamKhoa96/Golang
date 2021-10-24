//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage) // NewEvent, NewGreeter, NewMessage are initializers or providers
	return Event{}, nil
}
