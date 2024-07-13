package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func addPlant(c *gin.Context) {
	var plant Plant
	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&plant)
	c.JSON(http.StatusOK, plant)
}

func getPlant(c *gin.Context) {
	var plant Plant
	id := c.Param("id")
	if err := db.First(&plant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
		return
	}
	c.JSON(http.StatusOK, plant)
}

func classifySeed(c *gin.Context) {
	var seed Seed
	if err := c.ShouldBindJSON(&seed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&seed)
}

func sow() {
	var err error
	db, err = gorm.Open(sqlite.Open("flower.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Plant{}, &Seed{}, &Species{})

	r := gin.Default()

	r.POST("/plants", addPlant)
	r.GET("/plants/:id", getPlant)
	r.POST("/seeds", classifySeed)

	r.Run(":8080")
}
