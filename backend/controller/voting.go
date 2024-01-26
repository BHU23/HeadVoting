package controller

import (
	"fmt"
	"net/http"
	"crypto/sha256"
	"crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    // "encoding/base64"
    "encoding/pem"
    "log"

	"github.com/BHU23/HeadVoting/entity"
	"github.com/gin-gonic/gin"
)

type votingPayload struct {
	HashVote   string
	Signeture  string
	StudenID   string
	VoterID    *uint
	CandidatID *uint
	HashAuthen string
}

func RsaDecrypt(privateKeyBytes, cipherText []byte) ([]byte, error) {
    block, _ := pem.Decode(privateKeyBytes)
    if block == nil {
        return nil, fmt.Errorf("failed to parse PEM block containing the private key")
    }

    privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %v", err)
    }

    decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
    if err != nil {
        return nil, fmt.Errorf("failed to decrypt: %v", err)
    }
	log.Println(string(decrypted))
    return decrypted, nil
}


// POST /voting
func CreateVoting(c *gin.Context) {
	var data votingPayload
	var candidat entity.Candidat
	var voter entity.Voter

	// bind เข้าตัวแปร voting
	if err := c.ShouldBindJSON(&data); err != nil {
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
	db.First(&candidat, data.CandidatID)
	if candidat.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "candidat not found"})
		return
	}

	// ค้นหา voter ด้วย id

	if tx := entity.DB().Where("student_id = ?", data.StudenID).First(&voter); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voter not found"})
		return
	}

voter.PublishKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDlOJu6TyygqxfWT7eLtGDwajtNFOb9I5XRb6khyfD1Yt3YiCgQ
WMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76xFxdU6jE0NQ+Z+zEdhUTooNR
aY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4gwQco1KRMDSmXSMkDwIDAQAB
AoGAfY9LpnuWK5Bs50UVep5c93SJdUi82u7yMx4iHFMc/Z2hfenfYEzu+57fI4fv
xTQ//5DbzRR/XKb8ulNv6+CHyPF31xk7YOBfkGI8qjLoq06V+FyBfDSwL8KbLyeH
m7KUZnLNQbk8yGLzB3iYKkRHlmUanQGaNMIJziWOkN+N9dECQQD0ONYRNZeuM8zd
8XJTSdcIX4a3gy3GGCJxOzv16XHxD03GW6UNLmfPwenKu+cdrQeaqEixrCejXdAF
z/7+BSMpAkEA8EaSOeP5Xr3ZrbiKzi6TGMwHMvC7HdJxaBJbVRfApFrE0/mPwmP5
rN7QwjrMY+0+AbXcm8mRQyQ1+IGEembsdwJBAN6az8Rv7QnD/YBvi52POIlRSSIM
V7SwWvSK4WSMnGb1ZBbhgdg57DXaspcwHsFV7hByQ5BvMtIduHcT14ECfcECQATe
aTgjFnqE/lQ22Rk0eGaYO80cc643BXVGafNfd9fcvwBMnk0iGX0XRsOozVt5Azil
psLBYuApa66NcVHJpCECQQDTjI2AQhFc1yRnCU/YgDnSpJVm1nASoRUnU8Jfm3Oz
uku7JUXcVpt08DFSceCEX9unCuMcT72rAQlLpdZir876
-----END RSA PRIVATE KEY-----
`
	publishKey := []byte(voter.PublishKey)
	log.Println(string(publishKey))
	cipherText := (data.Signeture)
	log.Println(string(data.Signeture))
	log.Println(string(cipherText))
	SignedData, _ := RsaDecrypt(publishKey, []byte(cipherText))
	SignedDigest := fmt.Sprintf("%x", SignedData)
	log.Println(string(SignedDigest))
	log.Println(string("========="))
	// Concatenate the required values for hashing
	hashData := data.StudenID + string(candidat.NameCandidat)
	// Hash the concatenated data using SHA-256
	hashedData := sha256.Sum256([]byte(hashData))
	hashDigest := fmt.Sprintf("%x", hashedData)
	log.Println(string(hashDigest))
	log.Println(string(data.HashAuthen))

	
	if (SignedDigest != hashDigest){
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not authention, Please check you private key !"})
		return
	}
	fmt.Print(hashDigest)
	fmt.Print(SignedDigest)
	fmt.Print("==========")
	if hashDigest == SignedDigest {
		fmt.Print("1")
	} else {
		fmt.Print("00000000000")
	}
	// สร้าง Voting
	u := entity.Voting{
		StudenID: data.StudenID,
		HashVote: data.HashAuthen,
		// HashVote:   data.HashVote,
		Signeture:  data.Signeture,
		VoterID:    voter.ID,
		Voter:      voter,
		CandidatID: candidat.ID,
		Candidat:   candidat,
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
