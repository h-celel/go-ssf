package sql

import (
	"database/sql"
	goSSF "github.com/h-celel/go-ssf"
)

func GetDefaultComponent(service goSSF.Service) Component {
	if service == nil {
		return nil
	}

	cs := service.GetComponentsByType(ComponentType)

	for _, c := range cs {
		if c, ok := c.(Component); ok {
			return c
		}
	}

	return nil
}

func GetComponent(service goSSF.Service, i uint32) Component {
	if service == nil {
		return nil
	}

	c := service.GetComponent(ComponentType, i)

	if c, ok := c.(Component); ok {
		return c
	}

	return nil
}

func GetDefaultDB(service goSSF.Service) *sql.DB {
	if service == nil {
		return nil
	}

	cs := service.GetComponentsByType(ComponentType)

	for _, c := range cs {
		if c, ok := c.(Component); ok {
			return c.GetDB()
		}
	}

	return nil
}

func GetDB(service goSSF.Service, i uint32) *sql.DB {
	if service == nil {
		return nil
	}

	c := service.GetComponent(ComponentType, i)

	if c, ok := c.(Component); ok {
		return c.GetDB()
	}

	return nil
}
