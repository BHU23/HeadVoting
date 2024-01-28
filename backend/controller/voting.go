package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"net/http"

	"github.com/BHU23/HeadVoting/entity"
	"github.com/gin-gonic/gin"
)

type votingPayload struct {
	HashVote   string
	Signeture  string
	StudenID   string
	VoterID    *uint
	CandidatID *uint
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

	Datahash := data.StudenID + string(candidat.NameCandidat)

	// Hash the concatenated data using SHA-512
	hashDigestByte := sha512.Sum512([]byte(Datahash))

	inPutSignature := data.Signeture

	// Insert line breaks and headers into the public key
	formattedPublicKey := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", voter.PublicKey)

	publicKey := []byte(formattedPublicKey)
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

	signatureBytes, err := base64.StdEncoding.DecodeString(inPutSignature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting to RSA public key"})
		return
	}

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA512, hashDigestByte[:], signatureBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signature verification failed:"})
		return
	}

	fmt.Println("Signature verified successfully!")

	var votings []entity.Voting
	var hashVotesData = ""
	db.Preload("Vote").Preload("Candidat").Find(&votings)
	if len(votings) == 0 {
		hashVotesData = data.StudenID + candidat.NameCandidat + data.Signeture
	} else {
		for i := 0; i < len(votings); i++ {
			hashVotesData += votings[i].StudenID + string(votings[i].Candidat.NameCandidat) + string(votings[i].Signeture)
		}
	}
	hashVotesData = hashVotesData + data.StudenID + candidat.NameCandidat + data.Signeture
	// Hash the concatenated data using SHA-256
	hashedVotesData := sha512.Sum512([]byte(hashVotesData))
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

// 1 code go PrivateKeyDecrypt 
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

// 2 code go publicDecrypt but not work
// func RsaDecrypt(publicKey []byte, ciphertext []byte) ([]byte, error) {
// 	block, _ := pem.Decode(publicKey)
// 	if block == nil {
// 		return nil, errors.New("public key error!")
// 	}
// 	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rsaPub, ok := pub.(*rsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("invalid public key type")
// 	}

// 	// testDecData := RSA_public_decrypt(rsaPub, ciphertext)
// 	testDecData, err := PublicDecrypt(rsaPub, ciphertext)
// 	if err != nil {
// 		return nil, errors.New("public decrypt failed")
// 	}
// 	err = nil
// 	return testDecData, err
// }

// func RSA_public_decrypt(pubKey *rsa.PublicKey, data []byte) []byte {
// 	c := new(big.Int)
// 	m := new(big.Int)
// 	m.SetBytes(data)
// 	e := big.NewInt(int64(pubKey.E))
// 	c.Exp(m, e, pubKey.N)
// 	out := c.Bytes()
// 	skip := 0
// 	for i := 2; i < len(out); i++ {
// 		if i+1 >= len(out) {
// 			break
// 		}
// 		if out[i] == 0xff && out[i+1] == 0 {
// 			skip = i + 2
// 			break
// 		}
// 	}

// 	return out[skip:]
// }
// // 3 code go publicDecrypt but not work
// var hashPrefixes = map[crypto.Hash][]byte{
// 	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
// 	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
// 	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
// 	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
// 	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
// 	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
// 	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
// 	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
// }

// // copy from crypt/rsa/pkcs1v5.go
// func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
// 	e := big.NewInt(int64(pub.E))
// 	c.Exp(m, e, pub.N)
// 	return c
// }

// // copy from crypt/rsa/pkcs1v5.go
// func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
// 	// Special case: crypto.Hash(0) is used to indicate that the data is
// 	// signed directly.
// 	if hash == 0 {
// 		return inLen, nil, nil
// 	}

// 	hashLen = hash.Size()
// 	if inLen != hashLen {
// 		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
// 	}
// 	prefix, ok := hashPrefixes[hash]
// 	if !ok {
// 		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
// 	}
// 	return
// }

// // copy from crypt/rsa/pkcs1v5.go
// func leftPad(input []byte, size int) (out []byte) {
// 	n := len(input)
// 	if n > size {
// 		n = size
// 	}
// 	out = make([]byte, size)
// 	copy(out[len(out)-n:], input)
// 	return
// }
// func unLeftPad(input []byte) (out []byte) {
// 	n := len(input)
// 	t := 2
// 	for i := 2; i < n; i++ {
// 		if input[i] == 0xff {
// 			t = t + 1
// 		} else {
// 			if input[i] == input[0] {
// 				t = t + int(input[1])
// 			}
// 			break
// 		}
// 	}
// 	out = make([]byte, n-t)
// 	copy(out, input[t:])
// 	return
// }

// // copy&modified from crypt/rsa/pkcs1v5.go
// func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
// 	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
// 	if err != nil {
// 		return nil, err
// 	}

// 	tLen := len(prefix) + hashLen
// 	k := (pub.N.BitLen() + 7) / 8
// 	if k < tLen+11 {
// 		return nil, fmt.Errorf("length illegal")
// 	}

// 	c := new(big.Int).SetBytes(sig)
// 	m := encrypt(new(big.Int), pub, c)
// 	em := leftPad(m.Bytes(), k)
// 	out = unLeftPad(em)

// 	err = nil
// 	return
// }

// func PrivateEncrypt(privt *rsa.PrivateKey, data []byte) ([]byte, error) {
// 	signData, err := rsa.SignPKCS1v15(nil, privt, crypto.Hash(0), data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return signData, nil
// }
// func PublicDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
// 	decData, err := publicDecrypt(pub, crypto.Hash(0), nil, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return decData, nil
// }

// func ImportSPKIPublicKeyPEM(spkiPEM string) *rsa.PublicKey {
// 	body, _ := pem.Decode([]byte(spkiPEM))
// 	publicKey, _ := x509.ParsePKIXPublicKey(body.Bytes)
// 	if publicKey, ok := publicKey.(*rsa.PublicKey); ok {
// 		return publicKey
// 	} else {
// 		return nil
// 	}
// }

// func decryptChunk(ciphertextBytesChunk []byte, writer *bytes.Buffer, pubKey *rsa.PublicKey) {
// 	// Decrypt each signature chunk
// 	ciphertextInt := new(big.Int)
// 	ciphertextInt.SetBytes(ciphertextBytesChunk)
// 	decryptedPaddedInt := decrypt(new(big.Int), pubKey, ciphertextInt)
// 	// Remove padding
// 	decryptedPaddedBytes := make([]byte, pubKey.Size())
// 	decryptedPaddedInt.FillBytes(decryptedPaddedBytes)
// 	start := bytes.Index(decryptedPaddedBytes[1:], []byte{0}) + 1 // // 0001FF...FF00<data>: Find index after 2nd 0x00
// 	decryptedBytes := decryptedPaddedBytes[start:]
// 	// Write decrypted signature chunk
// 	writer.Write(decryptedBytes)
// }

// func decrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
// 	// Textbook RSA
// 	e := big.NewInt(int64(pub.E))
// 	c.Exp(m, e, pub.N)
// 	return c
// }
