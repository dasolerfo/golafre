package main

import (
	"fmt"
	"log"
	"os"
	"golafre/config"
)


func main() {
	// Carrega el fitxer .env
	err := appSetup.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error carregant .env: %v", err)
	}

	// Llegeix variables d'entorn
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("App Port: %s\n", appPort)
	fmt.Printf("DB User: %s, Password: %s\n", dbUser, dbPass)
}