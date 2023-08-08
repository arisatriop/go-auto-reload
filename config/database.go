package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	gormSQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var sqlDB *sql.DB
var gormDB *gorm.DB

type DB struct {
	SqlDB  *sql.DB
	GormDB *gorm.DB
	SqlTx  *sql.Tx
	GormTx *gorm.DB
}

func CreateDBConnection() {
	sqlDB = SqlConnection()
	gormDB = GormConnection(sqlDB)
	fmt.Println("Connected to database")
}

func SqlConnection() *sql.DB {
	dbAddress := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 dbAddress,
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	sqlCon, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	sqlCon.SetMaxIdleConns(10)
	sqlCon.SetMaxOpenConns(100)
	sqlCon.SetConnMaxLifetime(5 * time.Minute)
	sqlCon.SetConnMaxIdleTime(60 * time.Minute)

	pingErr := sqlCon.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return sqlCon
}

func GormConnection(sqlDB *sql.DB) *gorm.DB {
	gormCon, err := gorm.Open(gormSQL.New(gormSQL.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	return gormCon
}

func GetDBConnection() *DB {
	return &DB{
		SqlDB:  sqlDB,
		GormDB: gormDB,
		SqlTx:  nil,
		GormTx: nil,
	}
}
