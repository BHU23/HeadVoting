package entity

import (
	"gorm.io/gorm"
)

type Voter struct {
	gorm.Model
	StudentID   string `gorm:"uniqueIndex"`
	StudentName string
	PublicKey  string

	Votings []Voting `gorm:"foreignKey:VoterID"`
}

type Candidat struct {
	gorm.Model
	NameCandidat string

	Votings []Voting `gorm:"foreignKey:CandidatID"`
}

type Voting struct {
	gorm.Model
	StudenID   string
	HashVote   string   `gorm:"type:longtext"`
	Signeture  string   `gorm:"type:longtext"`
	VoterID    uint   	`gorm:"uniqueIndex"` // ไม่ให้มีการโหวดซ้ำ
	Voter      Voter 	`gorm:"references:id"`
	CandidatID uint
	Candidat   Candidat `gorm:"references:id"`
}
