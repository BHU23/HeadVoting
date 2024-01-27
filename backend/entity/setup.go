package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectDB() (*gorm.DB, error) {
	var err error
	var database *gorm.DB
	database, err = gorm.Open(sqlite.Open("cs-66.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Voter{},
		&Candidat{},
		&Voting{},
	)
	db = database
	// SetUp Candidate
	database.Where(Candidat{NameCandidat: "A"}).FirstOrCreate(&Candidat{NameCandidat: "A"})
	database.Where(Candidat{NameCandidat: "B"}).FirstOrCreate(&Candidat{NameCandidat: "B"})
	database.Where(Candidat{NameCandidat: "C"}).FirstOrCreate(&Candidat{NameCandidat: "C"})

	// Voter Data (ex)
	voter := []Voter{
		{
			StudentID:  "B6400001",
			StudentName: "s1",
			PublishKey: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCRl3fhbfVHn8+45qhNaVV3ApaLK6iQZ+mpwY1Cme7iSXYntbQCvhpXxILDnYEGGJu1Yndtgn+HQMnJH3lAeb1vDuOlVJF34wkQcwqAVuxE0QS92mJvTIVhU89TO+mFMzYV5Vklj2Smg6M/fQDIvLBIUJGQTRT7IK5bYSAo+RqXwIDAQAB",
		},
		{
			StudentID:  "B6400002",
			StudentName: "s2",
			PublishKey: "",
		},
		{
			StudentID:  "B6400003",
			StudentName: "s3",
			PublishKey: "",
		},
		{
			StudentID:  "B6400004",
			StudentName: "s4",
			PublishKey: "",
		},
		{
			StudentID:  "B6400005",
			StudentName: "s5",
			PublishKey: "",
		},
		{
			StudentID:  "B6400006",
			StudentName: "s6",
			PublishKey: "",
		},
		{
			StudentID:  "B6400007",
			StudentName: "s7",
			PublishKey: "",
		},
		{
			StudentID:  "B6400008",
			StudentName: "s8",
			PublishKey: "",
		},
		{
			StudentID:  "B6400009",
			StudentName: "s9",
			PublishKey: "",
		},
		{
			StudentID:  "B6400010",
			StudentName: "s10",
			PublishKey: "",
		},
	}

	for _, voter := range voter {
		db.Create(&voter)
	}
	return database, nil
}
