package main

import (
	"fiber-app/pkg/server"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando aplicación...")
	app := server.NewApp()
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
