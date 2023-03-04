package db

import "time"

type DataType int

const (
	Integer = iota + 1
	Double
	Varchar
	TextField
	DateTime
	Time
	Date
	Timestamp
)

type MigrationInfo struct {
	Batches struct {
		Timestamp     time.Time
		MigrationName string
	}
	Exists bool
}

type TableCreationConfig struct {
	Columns []struct {
		Name     string
		DataType DataType
	}
}

type TableModificationConfig struct {
	Columns []struct {
		Name     string
		DataType DataType
	}
}

type Migrator interface {
	GetCurrentMigrationInfo() (MigrationInfo, error)
	CreateTable(tableName, config TableCreationConfig) error
	UpdateTable(tableName, config TableModificationConfig) error
	DeleteTable(tableName string) error
}
