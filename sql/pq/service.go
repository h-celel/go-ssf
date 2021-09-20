package pq

import (
	"database/sql"

	goSSF "github.com/h-celel/go-ssf"
	ssfSQL "github.com/h-celel/go-ssf/sql"

	_ "github.com/lib/pq"
)

func InitComponent(service goSSF.Service, options *Option) error {
	if service == nil {
		return nil
	}

	if options == nil {
		options = NewOptionFromEnvironment()
	}

	var db *sql.DB
	var err error
	for i := uint(0); i < options.ConnectRetries; i++ {
		db, err = sql.Open("postgres", options.connectionString())
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}

	c := &component{db: db}

	service.AddComponent(ssfSQL.ComponentType, c)

	return nil
}
