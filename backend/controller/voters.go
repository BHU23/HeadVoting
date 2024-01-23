package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/BHU23/HeadVoting/entity"
)

// GET /voter
func ListVoters(c *gin.Context) {

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var voters []entity.Voter
	db.Find(&voters)
	c.JSON(http.StatusOK, voters)
}

// GET /voting // นำไปใช้ในการแสดง จำนวนผู้ใช้สิทธ์
func ListVoting(c *gin.Context) {

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var voting []entity.Voting
	db.Find(&voting)
	c.JSON(http.StatusOK, voting)
}