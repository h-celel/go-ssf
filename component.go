package go_ssf

import (
	"context"
)

type ComponentType string

type Component interface {
	Status(ctx context.Context) error
}
