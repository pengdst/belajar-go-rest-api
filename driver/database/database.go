package database

import (
	"gorm.io/gorm"
)

const (
	SQL              = "sql"
	NoSQL            = "nosql"
	DriverPostgresql = "postgresql"
	DriverSqlite     = "sqlite"
	DriverMysql      = "mysql"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

type Database struct {
	*gorm.DB
	Connections map[string]*gorm.DB
}

func NewDatabase() *Database {
	return &Database{
		Connections: make(map[string]*gorm.DB),
	}
}

func (d Database) GetConnection(names ...string) *gorm.DB {
	name := ""
	if len(names) > 0 {
		name = names[0]
	}

	if name == "" {
		return d.DB
	}

	db, ok := d.Connections[name]
	if !ok {
		return nil
	}

	return db
}

func (d *Database) AddConnection(name string, config Config) error {
	var db *gorm.DB
	var err error
	if config.Driver == DriverPostgresql {
		db, err = ConfigurePostgresql(config)
	} else if config.Driver == DriverMysql {
		db, err = ConfigureMysql(config)
	} else if config.Driver == DriverSqlite {
		db, err = ConfigureSqlite()
	}

	if err != nil {
		return err
	}

	d.Connections[name] = db
	return nil
}
