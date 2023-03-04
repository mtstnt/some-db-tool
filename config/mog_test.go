package config

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestLoadMogConfig(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    MogConfig
		wantErr bool
	}{
		{
			name: "test mog config loads database",
			args: args{
				reader: strings.NewReader(`
database:
    driver: mysql
    host: localhost
    port: 3306
    name: testdb
    user: root
    password: 123456 # Use ENV_[YOUR ENVIRONMENT VARIABLE NAME] to load from environment variables.`,
				),
			},
			want: MogConfig{
				Database: struct {
					Driver   string
					Host     string
					Port     int
					Name     string
					User     string
					Password string
				}{
					Driver:   "mysql",
					Host:     "localhost",
					Port:     3306,
					Name:     "testdb",
					User:     "root",
					Password: "123456",
				},
			},
		},
		{
			name: "test mog config loads project section",
			args: args{
				reader: strings.NewReader(`
project:
    migration_dir: ./migrations
`),
			},
			want: MogConfig{
				Project: struct {
					MigrationDir string "yaml:\"migration_dir\""
				}{
					MigrationDir: "./migrations",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadMogConfig(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMogConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadMogConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
