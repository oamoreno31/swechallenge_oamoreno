package main

import (
	"log"
	"os"
	"sweapi/config"
	"sweapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	}

	// Conectar a la base de datos
	config.ConnectDB()
	defer func() {
		if err := config.DB.Close(nil); err != nil {
			log.Printf("Error al cerrar la base de datos: %v", err)
		}
	}()

	// Configurar router
	router := gin.Default()

	// Configurar rutas
	routes.SetupRoutes(router)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en puerto %s", port)
	router.Run(":" + port)
}
