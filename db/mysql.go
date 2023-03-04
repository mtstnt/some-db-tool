package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mtstnt/mog/config"
)

var _ Migrator = (*MySQLMigrator)(nil)

type MySQLMigrator struct {
	table string
	db    *sql.DB
}

func MySQLGenerateDSN(conf *config.MogConfig) string {
	return fmt.Sprintf("%s:%s(%s:%d)/%s", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Name)
}

func NewMySQLMigrator(migrationTableName string, db *sql.DB) *MySQLMigrator {
	return &MySQLMigrator{
		table: migrationTableName,
		db:    db,
	}
}

func (m *MySQLMigrator) GetCurrentMigrationInfo() (MigrationInfo, error) {
	return MigrationInfo{}, nil
}

func (m *MySQLMigrator) CreateTable(tableName, config TableCreationConfig) error {
	return nil
}

func (m *MySQLMigrator) UpdateTable(tableName, config TableModificationConfig) error {
	return nil
}

func (m *MySQLMigrator) DeleteTable(tableName string) error {
	return nil
}
