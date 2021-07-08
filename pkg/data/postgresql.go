package data

import "github.com/upper/db/v4/adapter/postgresql"

// ConnectionURL defines the DSN attributes.
var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `locahost`,
	User:     `postgres`,
	Password: ``,
}
