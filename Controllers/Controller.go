package Controllers

import (
	"CytusSongsSelector/Models"
	"CytusSongsSelector/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSongsController(c *gin.Context) {
	var songs []Models.Song
	db := Services.DbConnect()
	db.Find(&songs)
	c.JSON(http.StatusOK, gin.H{"data": songs})
}

func AddSongController(c *gin.Context) {
	var request Models.Song
	db := Services.DbConnect()
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song := Models.Song{Chapter: request.Chapter, Sequence: request.Sequence, Name: request.Name, Difficulty: request.Difficulty}
	db.Create(&song)
	c.JSON(http.StatusOK, gin.H{"data": song})
}

//func EditSongController(c *gin.Context) {
//	var request Models.Song
//	db := Services.DbConnect()
//	if err := c.ShouldBindJSON(&request); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//}

func GetSongsByDifficultyController(c *gin.Context) {
	var song Models.Song
	db := Services.DbConnect()
	if err := db.Where("Difficulty = ?", c.Param("difficulty")).First(&song).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": song})
}
