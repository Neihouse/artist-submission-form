package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ArtistSubmission struct {
	ID            uint   `gorm:"primary_key;auto_increment"`
	BandName      string `gorm:"size:255"`
	Photo         string `gorm:"size:255"`
	Type          string `gorm:"size:50"`
	PersonalName  string `gorm:"size:255"`
	ContactEmail  string `gorm:"size:255"`
	PhoneNumber   string `gorm:"size:20"`
	Website       string `gorm:"size:255"`
	Location      string `gorm:"size:255"`
	RecordLabel   string `gorm:"size:255"`
	OpeningFor    string `gorm:"size:255"`
	Configuration string `gorm:"size:100"`
	Genres        string `gorm:"size:255"`
	Bio           string `gorm:"type:text"`
	MusicLinks    string `gorm:"type:json"`
	SocialLinks   string `gorm:"type:json"`
	Message       string `gorm:"type:text"`
}

func main() {
	r := gin.Default()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&ArtistSubmission{})

	r.POST("/submit", func(c *gin.Context) {
		var submission ArtistSubmission
		if err := c.ShouldBindJSON(&submission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&submission)
		c.JSON(http.StatusOK, gin.H{"message": "Submission received"})
	})

	r.Run()
}
