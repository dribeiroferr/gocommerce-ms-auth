package databases

import "database/sql"

// DBManager defines the interface for managing database connections
type DBManager interface { 
	Connect() (*sql.DB, error)
}