package databases

import (
	"database/sql"
	"fmt"
	"time"

	commons "github.com/dribeiroferr/gocommerce-ms-auth/common"
	"github.com/dribeiroferr/gocommerce-ms-auth/config"
	_ "github.com/lib/pq"
)

// PostgresManager manages PostgreSQL connections
type PostgresManager struct {
	config *config.Config
}

// NewPostgresManager creates a new PostgresManager
func NewPostgresManager(cfg *config.Config) *PostgresManager {
	return &PostgresManager{config: cfg}
}

// Connect establishes a connection to the PostgreSQL database
func (pm *PostgresManager) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pm.config.RelationalDBHost, pm.config.RelationalDBPort, pm.config.RelationalDBUser,
		pm.config.RelationalDBPassword, pm.config.RelationalDBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Set max configs for the connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	commons.Log.Info("Successfully connected to PostgreSQL")

	return db, nil
}
