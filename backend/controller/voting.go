package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/BHU23/HeadVoting/entity"
)

// POST /voting
func CreateVoting(c *gin.Context) {
	var voting entity.Voting
	var candidat entity.Candidat
	var voter entity.Voter

	// bind เข้าตัวแปร voting
	if err := c.ShouldBindJSON(&voting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา candidat ด้วย id
	db.First(&candidat, voting.CandidatID)
	if candidat.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "candidat not found"})
		return
	}

	// ค้นหา voter ด้วย id

	if tx := entity.DB().Where("student_id = ?", voting.StudenID).First(&voter); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voter not found"})
		return
	}


	// สร้าง Voting
	u := entity.Voting{
		StudenID: voting.StudenID,
		// HashVote: voting.HashVote,
		Signeture: voting.Signeture, 
		VoterID: voter.ID,
		Voter: voter,
		CandidatID: candidat.ID,
		Candidat: candidat,
	}

	// บันทึก
	if err := db.Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": u})
}

// GET /voting/:id
func GetVoting(c *gin.Context) {
	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var voting entity.Voting
	id := c.Param("id")
	db.Preload("Vote").Preload("Candidat").First(&voting, id)
	if voting.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "voting not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": voting})
}

// GET /Voting
func ListVotings(c *gin.Context) {

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var votings []entity.Voting
	db.Preload("Vote").Preload("Candidat").Find(&votings)
	c.JSON(http.StatusOK, votings)
}

// DELETE /votings/:id
func DeleteUser(c *gin.Context) {

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	var voting entity.Voting
	db.First(&voting, id)
	if voting.ID != 0 {
		db.Delete(&voting)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted success"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
	}

}

// PATCH /votings
func UpdateVoting(c *gin.Context) {
	var voting entity.Voting
	var result entity.Voting

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&voting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา user ด้วย id
	if tx := db.Where("id = ?", voting.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := db.Save(&voting).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": voting})
}

// GET /users
func ListCandidat(c *gin.Context) {

	var candidate []entity.Candidat
	if err := entity.DB().Raw("SELECT * FROM candidats").Scan(&candidate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": candidate})
}