package controller

import (
	"bytes"
	"crypto/rand"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	// "encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	// "log"
	"net/http"
	// "io"
	"math/big"

	// "encoding/hex"

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

// 1
//
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
//
// 2
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

	// testDecData := RSA_public_decrypt(rsaPub, ciphertext)
	testDecData, err := PublicDecrypt(rsaPub, ciphertext)
	if err != nil {
		return nil, errors.New("public decrypt failed")
	}
	err = nil
	return testDecData, err
}

func RSA_public_decrypt(pubKey *rsa.PublicKey, data []byte) []byte {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}

	return out[skip:]
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

	PrivateKey := `
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
	voter.PublishKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDlOJu6TyygqxfWT7eLtGDwajtN
FOb9I5XRb6khyfD1Yt3YiCgQWMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76
xFxdU6jE0NQ+Z+zEdhUTooNRaY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4
gwQco1KRMDSmXSMkDwIDAQAB
-----END PUBLIC KEY-----
`
	// Generate a new RSA key pair
	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	fmt.Println("Error generating RSA key pair:", err)
	// 	return
	// }
	privateKey := []byte(PrivateKey)
	block1, _ := pem.Decode(privateKey)
	if block1 == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding private key block"})
		return
	}
	pri, err := x509.ParsePKCS1PrivateKey(block1.Bytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing private key", "details": err.Error()})
		return
	}

	// Example data to be signed
	// daSrecret := "Hello, RSA!"
	// hasheddaSrecret := sha256.Sum256([]byte(daSrecret))

	hashData := data.StudenID + string(candidat.NameCandidat)
	// Hash the concatenated data using SHA-256
	hasheddaSrecret := sha256.Sum256([]byte(hashData))
	hashDigest := fmt.Sprintf("%x", hasheddaSrecret)
	fmt.Println("hashDigest:", hashDigest)


	// Sign the data using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hasheddaSrecret[:])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error signing data", "details": err.Error()})
		return
	}
	inPutSignature := data.Signeture
	signatureInput := []byte(inPutSignature)

	fmt.Println("signatureInput:", signatureInput)
	fmt.Println("Signature:", signature)
	fmt.Println("signatureInput:", inPutSignature)
	fmt.Println("Signature:", string(signature))

	// hashData := data.StudenID + string(candidat.NameCandidat)
	// // Hash the concatenated data using SHA-256
	// hashedData := sha256.Sum256([]byte(hashData))
	// // hashDigest := fmt.Sprintf("%x", hashedData)

	publicKey := []byte(voter.PublishKey)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding public key block"})
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing public key", "details": err.Error()})
		return
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting to RSA public key"})
		return
	}

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hasheddaSrecret[:], signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Signature verification failed:"})
		return
	}

	fmt.Println("Signature verified successfully!")

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
	hashedVotesData := sha256.Sum256([]byte(hashVotesData))
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

var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// copy from crypt/rsa/pkcs1v5.go
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// copy from crypt/rsa/pkcs1v5.go
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v5.go
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
func unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

// copy&modified from crypt/rsa/pkcs1v5.go
func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := encrypt(new(big.Int), pub, c)
	em := leftPad(m.Bytes(), k)
	out = unLeftPad(em)

	err = nil
	return
}

func PrivateEncrypt(privt *rsa.PrivateKey, data []byte) ([]byte, error) {
	signData, err := rsa.SignPKCS1v15(nil, privt, crypto.Hash(0), data)
	if err != nil {
		return nil, err
	}
	return signData, nil
}
func PublicDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	decData, err := publicDecrypt(pub, crypto.Hash(0), nil, data)
	if err != nil {
		return nil, err
	}
	return decData, nil
}

func ImportSPKIPublicKeyPEM(spkiPEM string) *rsa.PublicKey {
	body, _ := pem.Decode([]byte(spkiPEM))
	publicKey, _ := x509.ParsePKIXPublicKey(body.Bytes)
	if publicKey, ok := publicKey.(*rsa.PublicKey); ok {
		return publicKey
	} else {
		return nil
	}
}

func decryptChunk(ciphertextBytesChunk []byte, writer *bytes.Buffer, pubKey *rsa.PublicKey) {
	// Decrypt each signature chunk
	ciphertextInt := new(big.Int)
	ciphertextInt.SetBytes(ciphertextBytesChunk)
	decryptedPaddedInt := decrypt(new(big.Int), pubKey, ciphertextInt)
	// Remove padding
	decryptedPaddedBytes := make([]byte, pubKey.Size())
	decryptedPaddedInt.FillBytes(decryptedPaddedBytes)
	start := bytes.Index(decryptedPaddedBytes[1:], []byte{0}) + 1 // // 0001FF...FF00<data>: Find index after 2nd 0x00
	decryptedBytes := decryptedPaddedBytes[start:]
	// Write decrypted signature chunk
	writer.Write(decryptedBytes)
}

func decrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	// Textbook RSA
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// 	voter.PublishKey = `
// -----BEGIN PUBLIC KEY-----
// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoZ67dtUTLxoXnNEzRBFB
// mwukEJGC+y69cGgpNbtElQj3m4Aft/7cu9qYbTNguTSnCDt7uovZNb21u1vpZwKH
// yVgFEGO4SA8RNnjhJt2D7z8RDMWX3saody7jo9TKlrPABLZGo2o8vadW8Dly/v+I
// d0YDheCkVCoCEeUjQ8koXZhTwhYkGPu+vkdiqX5cUaiVTu1uzt591aO5Vw/hV4DI
// hFKnOTnYXnpXiwRwtPyYoGTa64yWfi2t0bv99qz0BgDjQjD0civCe8LRXGGhyB1U
// 1aHjDDGEnulTYJyEqCzNGwBpzEHUjqIOXElFjt55AFGpCHAuyuoXoP3gQvoSj6RC
// sQIDAQAB
// -----END PUBLIC KEY-----
// `

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

// //////////////////////////////////////
// 	  pubKeyPem := `-----BEGIN PUBLIC KEY-----
// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoZ67dtUTLxoXnNEzRBFB
// mwukEJGC+y69cGgpNbtElQj3m4Aft/7cu9qYbTNguTSnCDt7uovZNb21u1vpZwKH
// yVgFEGO4SA8RNnjhJt2D7z8RDMWX3saody7jo9TKlrPABLZGo2o8vadW8Dly/v+I
// d0YDheCkVCoCEeUjQ8koXZhTwhYkGPu+vkdiqX5cUaiVTu1uzt591aO5Vw/hV4DI
// hFKnOTnYXnpXiwRwtPyYoGTa64yWfi2t0bv99qz0BgDjQjD0civCe8LRXGGhyB1U
// 1aHjDDGEnulTYJyEqCzNGwBpzEHUjqIOXElFjt55AFGpCHAuyuoXoP3gQvoSj6RC
// sQIDAQAB
// -----END PUBLIC KEY-----`
// encrypted := data.Signeture
// // Import public key
// pubKey := ImportSPKIPublicKeyPEM(voter.PublishKey);
// // Base64 decode ciphertext
// // ciphertextBytes, _ := base64.StdEncoding.DecodeString("ajQbkszbZ97YZaPSRBab9vj0DDLm9tTrQwSZ+ucPj+cYSmw06KLCtRH3SPn3b2DqSd1revLXqxMtSzFmjRvZ5F8y3nzdP8NJaRplOigbPFhKZTv7xBVK5ATEmLukgtI7f+d3KdmGUG+cyTkfxIrMBvB3BIS5oTiMNmC9pqLaWcDVF9qpuxnwEMQJbeO9nTklpdv+F8BrchHmeUkKRrMJBoPbbcfq9Hi4bHiFyxPWhwB66d/AryCKsFRhaX6hSkTL+0NvuhVhv98wdo3juv2Il50XKOCbfc8kUG628TcSK6n31piLF9cntSVTB/L/pVfcAxEwx4hcUhLuqmk6EZIJvGo0G5LM22fe2GWj0kQWm/b49Awy5vbU60MEmfrnD4/nGEpsNOiiwrUR90j5929g6knda3ry16sTLUsxZo0b2eRfMt583T/DSWkaZTooGzxYSmU7+8QVSuQExJi7pILSO3/ndynZhlBvnMk5H8SKzAbwdwSEuaE4jDZgvaai2lnA1RfaqbsZ8BDECW3jvZ05JaXb/hfAa3IR5nlJCkazCQaD223H6vR4uGx4hcsT1ocAeunfwK8girBUYWl+oUpEy/tDb7oVYb/fMHaN47r9iJedFyjgm33PJFButvE3Eiup99aYixfXJ7UlUwfy/6VX3AMRMMeIXFIS7qppOhGSCbxqNBuSzNtn3thlo9JEFpv2+PQMMub21OtDBJn65w+P5xhKbDToosK1EfdI+fdvYOpJ3Wt68terEy1LMWaNG9nkXzLefN0/w0lpGmU6KBs8WEplO/vEFUrkBMSYu6SC0jt/53cp2YZQb5zJOR/EiswG8HcEhLmhOIw2YL2motpZwNUX2qm7GfAQxAlt472dOSWl2/4XwGtyEeZ5SQpGswkGg9ttx+r0eLhseIXLE9aHAHrp38CvIIqwVGFpfqFKRMv7Q2+6FWG/3zB2jeO6/YiXnRco4Jt9zyRQbrbxNxIrqffWmIsX1ye1JVMH8v+lV9wDETDHiFxSEu6qaToRkgm8")
// ciphertextBytes, _ := base64.StdEncoding.DecodeString(encrypted)
// // Split ciphertext into signature chunks a 2048/8 bytes and decrypt each chunk
// reader := bytes.NewReader(ciphertextBytes)
// var writer bytes.Buffer
// ciphertextBytesChunk := make([]byte, 2048/8)
// for {
//     n, _ := io.ReadFull(reader, ciphertextBytesChunk)
//     if (n == 0) {
//         break
//     }
//     decryptChunk(ciphertextBytesChunk, &writer, pubKey)
// }
// // Concatenate decrypted signature chunks
// decryptedData := writer.String()
// fmt.Println(decryptedData)
// hexDecryptedDigest := hex.EncodeToString(ciphertextBytesChunk)
// log.Println(hexDecryptedDigest)
// //////////////////////////////////////
// encrypted := data.Signeture

// // Decode base64-encoded ciphertext
// cipherText, err := base64.StdEncoding.DecodeString(encrypted)
// if err != nil {
// 	// Handle decoding error
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Base64 decoding error"})
// 	return
// }

// // Decrypt using RSA publicKey key
// publicKey := []byte(voter.PublishKey)
// decryptedDigest, err := RsaDecrypt(publicKey, cipherText)
// if err != nil {
// 	// Handle decryption error
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "RSA decryption error"})
// 	return
// }
// log.Println(decryptedDigest)
// hexDecryptedDigest := hex.EncodeToString(decryptedDigest)
// log.Println(hexDecryptedDigest)
// log.Println("=========")

//==========================================
// Concatenate the required values for hashing
// hashData := data.StudenID + string(candidat.NameCandidat)

// // Hash the concatenated data using SHA-256
// hashedData := sha256.Sum256([]byte(hashData))
// hashDigest := fmt.Sprintf("%x", hashedData)
// log.Println(string("hashDigest"))
// log.Println(string(hashDigest))
// Compare the decrypted digest with the hash
// if !bytes.Equal(ciphertextBytesChunk, []byte(hashDigest)) {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed. Check your private key!"})
// 	return
// }
