package sql

import (
	"database/sql"
	goSSF "github.com/h-celel/go-ssf"
)

const (
	ComponentType goSSF.ComponentType = "SQLComponent"
)

type Component interface {
	goSSF.Component
	GetDB() *sql.DB
}
