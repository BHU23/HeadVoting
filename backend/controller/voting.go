package controller

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"log"

	// "io/ioutil"

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

//	func RsaDecrypt(privateKey []byte, ciphertext []byte) ([]byte, error) {
//		block, _ := pem.Decode(privateKey)
//		if block == nil {
//			return nil, errors.New("private key error!")
//		}
//		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//		if err != nil {
//			return nil, err
//		}
//		return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
//	}
func RsaDecrypt(publicKey []byte, ciphertext []byte) ([]byte, error) {
    block, _ := pem.Decode(publicKey)
    if block == nil {
        return nil, errors.New("public key error!")
    }
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        return nil, errors.New("invalid public key type")
    }
	log.Println(pub)
	log.Println(rsaPub)
	log.Println("=====")
    // Use RSA encryption with public key for decryption
    return rsa.EncryptPKCS1v15(rand.Reader, rsaPub, ciphertext)
    // return rsa.DecryptPKCS1v15(rand.Reader, rsaPub, ciphertext)
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

	// 	voter.PublishKey = `
	// -----BEGIN RSA PRIVATE KEY-----
	// MIICXQIBAAKBgQDlOJu6TyygqxfWT7eLtGDwajtNFOb9I5XRb6khyfD1Yt3YiCgQ
	// WMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76xFxdU6jE0NQ+Z+zEdhUTooNR
	// aY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4gwQco1KRMDSmXSMkDwIDAQAB
	// AoGAfY9LpnuWK5Bs50UVep5c93SJdUi82u7yMx4iHFMc/Z2hfenfYEzu+57fI4fv
	// xTQ//5DbzRR/XKb8ulNv6+CHyPF31xk7YOBfkGI8qjLoq06V+FyBfDSwL8KbLyeH
	// m7KUZnLNQbk8yGLzB3iYKkRHlmUanQGaNMIJziWOkN+N9dECQQD0ONYRNZeuM8zd
	// 8XJTSdcIX4a3gy3GGCJxOzv16XHxD03GW6UNLmfPwenKu+cdrQeaqEixrCejXdAF
	// z/7+BSMpAkEA8EaSOeP5Xr3ZrbiKzi6TGMwHMvC7HdJxaBJbVRfApFrE0/mPwmP5
	// rN7QwjrMY+0+AbXcm8mRQyQ1+IGEembsdwJBAN6az8Rv7QnD/YBvi52POIlRSSIM
	// V7SwWvSK4WSMnGb1ZBbhgdg57DXaspcwHsFV7hByQ5BvMtIduHcT14ECfcECQATe
	// aTgjFnqE/lQ22Rk0eGaYO80cc643BXVGafNfd9fcvwBMnk0iGX0XRsOozVt5Azil
	// psLBYuApa66NcVHJpCECQQDTjI2AQhFc1yRnCU/YgDnSpJVm1nASoRUnU8Jfm3Oz
	// uku7JUXcVpt08DFSceCEX9unCuMcT72rAQlLpdZir876
	// -----END RSA PRIVATE KEY-----
	// `
	voter.PublishKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDlOJu6TyygqxfWT7eLtGDwajtN
FOb9I5XRb6khyfD1Yt3YiCgQWMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76
xFxdU6jE0NQ+Z+zEdhUTooNRaY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4
gwQco1KRMDSmXSMkDwIDAQAB
-----END PUBLIC KEY-----
`
	// encrypted := data.Signeture

	// // Special thanks to: https://stackoverflow.com/a/53077471
	// privateKey := []byte(voter.PublishKey)
	// cipherText, _ := base64.StdEncoding.DecodeString(encrypted)
	// SignedDigest, _ := RsaDecrypt(privateKey, []byte(cipherText))
	// log.Println(string(SignedDigest))
	// log.Println(string("========="))

	// // Concatenate the required values for hashing
	// hashData := data.StudenID + string(candidat.NameCandidat)
	// // Hash the concatenated data using SHA-256
	// hashedData := sha256.Sum256([]byte(hashData))
	// hashDigest := fmt.Sprintf("%x", hashedData)
	// log.Println(string(hashDigest))

	// if string(SignedDigest) != hashDigest {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Can not authention, Please check you private key !"})
	// 	return
	// }
	encrypted := data.Signeture

	// Decode base64-encoded ciphertext
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		// Handle decoding error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Base64 decoding error"})
		return
	}

	// Decrypt using RSA publicKey key
	publicKey := []byte(voter.PublishKey)
	decryptedDigest, err := RsaDecrypt(publicKey, cipherText)
	if err != nil {
		// Handle decryption error
		c.JSON(http.StatusBadRequest, gin.H{"error": "RSA decryption error"})
		return
	}
	log.Println(string(decryptedDigest))
	log.Println(string("========="))

	// Concatenate the required values for hashing
	hashData := data.StudenID + string(candidat.NameCandidat)

	// Hash the concatenated data using SHA-256
	hashedData := sha256.Sum256([]byte(hashData))
	hashDigest := fmt.Sprintf("%x", hashedData)
	log.Println(string(hashDigest))
	// Compare the decrypted digest with the hash
	if !bytes.Equal(decryptedDigest, []byte(hashDigest)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed. Check your private key!"})
		return
	}
	var votings []entity.Voting
	var hashVotesData = ""
	db.Preload("Vote").Preload("Candidat").Find(&votings)
	if len(votings) == 0 {
		hashVotesData = data.StudenID + candidat.NameCandidat
	} else {
		for i := 0; i < len(votings); i++ {
			hashVotesData += votings[i].StudenID + string(votings[i].Candidat.NameCandidat)
		}
	}

	// Hash the concatenated data using SHA-256
	hashedVotesData := sha256.Sum256([]byte(hashData))
	HashVotes := fmt.Sprintf("%x", hashedVotesData)
	// สร้าง Voting
	u := entity.Voting{
		StudenID:   data.StudenID,
		HashVote:   string(HashVotes),
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
