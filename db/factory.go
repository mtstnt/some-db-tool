package db

import (
	"database/sql"
	"fmt"

	"github.com/mtstnt/mog/config"
)

func DBFactory(driverName string, tableName string, conf *config.MogConfig) (Migrator, error) {
	var migrator Migrator
	switch driverName {
	case "mysql":
		dsn := MySQLGenerateDSN(conf)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("error opening mysql db: %w", err)
		}
		migrator = NewMySQLMigrator(tableName, db)
	default:
		return nil, fmt.Errorf("error db factory, driver not found")
	}
	return migrator, nil
}
