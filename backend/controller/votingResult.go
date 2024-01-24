package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/BHU23/HeadVoting/entity"
)

func GetVotingByCandidateID_is_1(c *gin.Context) {
    var voting []entity.Voting

    if err := entity.DB().Preload("Voter").Preload("Candidat").Where("candidat_id = 1", ).Find(&voting).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": voting})
}

func GetVotingByCandidateID_is_2(c *gin.Context) {
    var voting []entity.Voting

    if err := entity.DB().Preload("Voter").Preload("Candidat").Where("candidat_id = 2", ).Find(&voting).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": voting})
}

func GetVotingByCandidateID_is_3(c *gin.Context) {
    var voting []entity.Voting

    if err := entity.DB().Preload("Voter").Preload("Candidat").Where("candidat_id = 3", ).Find(&voting).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": voting})
}