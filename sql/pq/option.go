package pq

import (
	"fmt"
	"net/url"

	"github.com/h-celel/mapenv"
)

type Option struct {
	Host           string `json:"host" mpe:"POSTGRESQL_HOST"`
	Port           uint32 `json:"port" mpe:"POSTGRESQL_PORT"`
	User           string `json:"user" mpe:"POSTGRESQL_USER"`
	Password       string `json:"password" mpe:"POSTGRESQL_PASSWORD"`
	DBName         string `json:"db_name" mpe:"POSTGRESQL_DB_NAME"`
	SSLMode        string `json:"ssl_mode" mpe:"POSTGRESQL_SSL_MODE"`
	ConnectTimeout uint   `json:"connect_timeout" mpe:"POSTGRESQL_CONNECT_TIMEOUT"`
	ConnectRetries uint   `json:"connect_retries" mpe:"POSTGRESQL_CONNECT_RETRIES"`
}

func NewOption() *Option {
	return &Option{
		Host:           "localhost",
		Port:           5432,
		User:           "postgres",
		Password:       "postgres",
		DBName:         "postgres",
		SSLMode:        "require",
		ConnectTimeout: 10,
		ConnectRetries: 1,
	}
}

func NewOptionFromEnvironment() *Option {
	o := NewOption()
	err := mapenv.Decode(o)
	if err != nil {
		panic(err)
	}
	return o
}

func (o Option) connectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s connect_timeout=%d",
		url.PathEscape(o.DBName),
		url.PathEscape(o.User),
		url.PathEscape(o.Password),
		url.PathEscape(o.Host),
		o.Port,
		url.PathEscape(o.SSLMode),
		o.ConnectTimeout)
}
