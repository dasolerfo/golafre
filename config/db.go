package config

import (
	"os"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"    
	_ "github.com/lib/pq"                 
	_ "github.com/mattn/go-sqlite3"       
)

//

type DB struct {
	ConnectionType string
	Host           string
	Port           string
	DB_Name        string
	UserName       string
	Password       string
	DSN			   string
	connection     *sql.DB
}

var db *DB

var lock sync.Mutex // Per fer segur el patró Singleton

// Singleton per inicialitzar la instància única
func GetInstance() *DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db = &DB{
				ConnectionType: os.Getenv("DB_CONNECTION"),
				Host:           os.Getenv("DB_HOST"),
				Port:           os.Getenv("DB_PORT"),
				DB_Name:        os.Getenv("DB_DATABASE"),
				UserName:       os.Getenv("DB_USERNAME"),
				Password:       os.Getenv("DB_PASSWORD"),
				DSN:            buildDSN(),
			}
			err := db.initConnection()
			if err != nil {
				log.Fatalf("Failed to initialize database: %v", err)
			}
		}
	}
	return db
}

// Inicialitza la connexió a la base de dades
func (d *DB) initConnection() error {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.UserName, d.Password, d.Host, d.Port, d.DB_Name)
	conn, err := sql.Open(d.ConnectionType, dsn)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	// Configuració de connexions
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(0)

	// Test de connexió
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	d.connection = conn
	log.Println("Database connection established successfully.")
	return nil
}

func buildDSN() string {
	connType := os.Getenv("DB_CONNECTION")
	switch connType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
	case "sqlite3":
		return os.Getenv("DB_DATABASE") // Per SQLite, el "DB_DATABASE" és el nom del fitxer
	default:
		log.Fatalf("Unsupported DB_CONNECTION type: %s", connType)
		return ""
	}
}

// Retorna la connexió SQL
func (d *DB) GetConnection() *sql.DB {
	return d.connection
}

// Allibera recursos tancant la connexió
func (d *DB) Close() {
	if d.connection != nil {
		if err := d.connection.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}