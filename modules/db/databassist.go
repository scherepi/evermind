package db
// our general database driver 

import (
	"database/sql"
	
	"https://github.com/mattn/go-sqlite3"
)

type DatabassistModule struct {

}

type EvermindDatabassist struct {
	Modules []DatabassistModule
}