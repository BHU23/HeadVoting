package main

import (
	"github.com/BHU23/HeadVoting/controller"
	"github.com/BHU23/HeadVoting/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.ConnectDB()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Voting Routes
	r.GET("/voting", controller.ListCandidat)
	r.GET("/voting/:id", controller.GetVoting)
	r.POST("/votings", controller.CreateVoting)
	r.PATCH("/voting", controller.UpdateVoting)
	r.DELETE("/voting/:id", controller.DeleteUser)

	// Candidats
	r.GET("/candidats", controller.ListCandidat)
	
	// }
	// Run the server
	r.GET("/voters", controller.ListVoters)
	r.GET("/votinglist", controller.ListVoting)


	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
