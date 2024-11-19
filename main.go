package main

import (
	"fmt"
	"log"
	"golafre/config"
)


func main() {
	// Carrega el fitxer .env
	err := config.LoadEnv(".env")

	if err != nil {
		log.Fatalf("Error carregant .env: %v", err)
	}

	database := config.GetInstance()
	
	// Llegeix variables d'entorn
	/*
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")*/

	fmt.Printf("App Name: %s\n", database.DB_Name)
	fmt.Printf("App Port: %s\n", database.Port)
	fmt.Printf("DB User: %s, Password: %s\n", database.UserName, database.Password)
}