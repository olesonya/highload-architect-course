package database

import (
	"errors"
	"fmt"

	spew "github.com/davecgh/go-spew/spew"
	logger "github.com/sirupsen/logrus"

	gormDriver "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	migrate "github.com/golang-migrate/migrate/v4"
	psqlMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config interface {
	GetDBHost() string
	GetDBPort() uint
	GetDBName() string
	GetDBUser() string
	GetDBPassword() string
	GetMigrationsPath() string
}

type Database struct {
	config Config
	gormDB *gorm.DB
}

func NewDatabase(config Config) *Database {
	return &Database{
		config: config,
	}
}

func getVersion(m *migrate.Migrate) uint {
	version, dirty, err := m.Version()

	if dirty {
		logger.Fatalf("dirty active version: %d", version)
	}

	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		logger.Fatalf("failed to check active version: %v", err)
	}

	return version
}

func (db *Database) getPath() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		db.config.GetDBHost(),
		db.config.GetDBPort(),
		db.config.GetDBUser(),
		db.config.GetDBPassword(),
		db.config.GetDBName(),
	)
}

func (db *Database) Connect() {
	options := []gorm.Option{
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		},
	}

	gormDB, err := gorm.Open(gormDriver.Open(db.getPath()), options...)
	if err != nil {
		logger.Fatalf("gorm.Open(...): %v", err)
	}

	logger.WithFields(logger.Fields{
		"host":   spew.Sprintf("%#v", db.config.GetDBHost()),
		"port":   spew.Sprintf("%#v", db.config.GetDBPort()),
		"dbname": spew.Sprintf("%#v", db.config.GetDBName()),
	}).Infof("Connected to the database")

	db.gormDB = gormDB
}

func (db *Database) GetGormDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) Disconnect() {
	conn, err := db.gormDB.DB()
	if err != nil {
		logger.Fatalf("db.gormDB.DB(): %v", err)
	}

	if err := conn.Close(); err != nil {
		logger.Fatalf("conn.Close(): %v", err)
	}

	logger.Infoln("Gracefull disconnect from the database")
}

func (db *Database) MigrateDB() {
	gormDB, err := gorm.Open(gormDriver.Open(db.getPath()))
	if err != nil {
		logger.Fatalf("gorm.Open(...): %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		logger.Fatalf("gormDB.DB(): %v", err)
	}

	driver, err := psqlMigrate.WithInstance(sqlDB, &psqlMigrate.Config{})
	if err != nil {
		logger.Fatalf("psqlMigrate.WithInstance(...): %v", err)
	}

	defer driver.Close()

	migrationURL := fmt.Sprintf("file://%s", db.config.GetMigrationsPath())

	migration, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		logger.Fatalf("migrate.NewWithDatabaseInstance(...): %v", err)
	}

	defer migration.Close()

	currentVersion := getVersion(migration)
	logger.Infof("Start migrate from current version: %d", currentVersion)

	err = migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("No migrations are required")

			return
		}

		logger.Fatalf("migrator.Up(): %v", err)
	}

	newVersion := getVersion(migration)
	logger.Infof("Current active version: %d", newVersion)
}
