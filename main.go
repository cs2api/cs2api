package main

import (
	"log"
	"os"

	"github.com/jexlor/cs2api/db"
	"github.com/jexlor/cs2api/dev"

	"github.com/jexlor/cs2api/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
		log.Printf("No PORT specified, defaulting to %s", port)
	}

	router := setupRouter()

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	apiGroup := router.Group("/cs2api")
	{
		apiGroup.GET("/", api.LandingPage)
		apiGroup.GET("/skins", api.GetAllSkins)
		apiGroup.GET("/skins/search", api.GetSkinById)
		apiGroup.GET("/skins/search/n", api.GetSkinByName)
		apiGroup.GET("/collections", api.GetCollections)
		apiGroup.GET("/collections/search/n", api.GetCollectionByName)
		apiGroup.POST("/skins", dev.AddSkin)
	}
	return router
}
