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
	Voter := []Voter{
		{
			StudentID:   "B6400001",
			StudentName: "s1",
			PublicKey: `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAox9f7etWRQWeYcaTkBbK
			BlE4hHG0ohPqQKTqQk7L3SdePmvbYMiW707KBjWYkuhX841faoXChkOrrmraqRVS
			n8W9Pap4Ukokfnp6KyQ6azW73DfGWiusVTmW893ZmmLR1U8Uve/YKyqkEmJ4pni3
			/Go5BbOwgiImp/0Y5PhXY4d3ic2EqtNyb0+t3VE9IoCmT32u10az9FPDdXn1QtnQ
			IBEhMbKvxuv9b8MQ4nNvkMRP56HwmSNbw76N1FAyiSQLrolMKJ/n+0ZwCClrwqAN
			fzvwNRseR5HZKkgmbxYYnqNfr5Ri7wmXdrgI0rxcjuJszUW19geHbCj8CKTofUdx
			PQIDAQAB`,
		},
		{
			StudentID:   "B6400002",
			StudentName: "s2",
			PublicKey: `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkgMsebUwlpppclvHQhab
			ujauI4OjHIFspE7Dh56KytO1F5vLgyTq+XQcgnpCspRrIeoeA2Krr5PFgA4RyhfT
			7SebRF2VOS789v6h34eAZAM2Twfd3gqTB9zsc9TzvqKQmJuONjxi6Ck/90V6m1k0
			uLtJgeHqwpOyH6juMdInKOx0rV1oc1gSaFvEQVHW1JtqrkuncoTeD86tJnPDb55h
			QxZNpCm9D3OC9dckdQbxMYY1B2DwqLFq9akNUaWloHzHA4KeLpkJqP1kx0pjP6DJ
			AIHDSJh4BaBffOfPJSz1P+NkmU/0i+8/kqENkDNjVYYNDAzUmC1njAhfz226xEAJ
			ZQIDAQAB`,
		},
		{
			StudentID:   "B6400003",
			StudentName: "s3",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm/t/H0hXXm8CaajKc0wm
			FIKkV9d8UY7ba6aLUTB/li0L71SWFxBEWYM6awcyUU+2QwvEl55bw/9kRhzwnC8/
			36Y+cQtFO8f5AR4BMgFnQkhM04lWDaP+oGQ+SLQyOK49jGNMjsJTYxlKCSUX8Ih4
			UkxmdzWB/C+QJMiBkZOsqGwGubkjcXaqSUeMDiQ90Xkt7K44/78Hg3FqxBK6ppGp
			jIKzgcczp/aYU9MluzxcO8gccCARnfbzcpeAVOts4KCmlZTBIwZApXiXP6bHZ9Ao
			QT5vbGPaPZvauBwvU3rhNsI+mpuB/GQOGf7hmIvi3axffq5N9iqPh+vG6AqMvT18
			LQIDAQAB`,
		},
		{
			StudentID:   "B6400004",
			StudentName: "s4",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqAPl9AZR6nH9qu3aZGwl
			UiqTEhC2OqSUAdKOrJAxGaFMu14i2CwaTgP1JcbpkRkUSUlKPSBzQkIu9smzdckU
			ci61e7deYPY7cKMQaxmYMixA2p+cWpdmzRLiv4NmLB7Zv8t6hM4G8FfsEfnbfRI/
			Hc4vCz2rZZX/v/rTfKihGwYXo40Hbfi8yozeZ7Di+WhtJwrFNBob9doHNrMe4KFd
			tmBRXsMN7h+38b0Bu5bDF5Gy0/WJcBLI3qYp9PrRdFecwDz/pZN2zbqDEcdDpi8s
			JsXyNZiu1mvMxFpY/PmpfxdKw8jpQQexSWAcylw5LLiorTNTi+uBJ3/cPNqS5U0n
			7QIDAQAB`,
		},
		{
			StudentID:   "B6400005",
			StudentName: "s5",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwAeiav/3LMuakuOjTd3G
			UinWIrEdW6JTb0/ansTA8Soq8E/zJFUySjvo9ba0OUHIkDCt6rIY2+0s1XZrNpe3
			z+1nl7zMEi3H4wREEsGNlhr2n/IlsHWauQpJchCkeWRWe4+HIpm3PQg7N9oV7/94
			7z4EQ4r642owE4R/3TgTz/xP0hb8sSn4hReQjI1jaJs9xKUQQDpP6wML5JKncc6u
			T6+tDrFnbHTtu4UsoGHPXyDYuGeNuB20hX8N0ZCEjh45AZsdc1AVNpiolqxV9HKO
			4MQA/4yUmITQ0zHNhr3eJA23gGksjLkxqY4zwVCISRyb46OxdDjVp0Wx4ceN7wRj
			TwIDAQAB`,
		},
		{
			StudentID:   "B6400006",
			StudentName: "s6",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAj3hahJloLZ5ZbIiv8Kxv
			b8wrIaOIumFCEr4oUhS22z3iJauP+3cdt74RI2EQ6ENv8IX3kWN6E9I86klPYaRb
			YVEoDFcWftu/JH2qbQK/LTRCZ5VRaAmlgVk+DIkC2UdhGVt2GgVUSjxORUb0H3B7
			fcRb2wV+awM2WtZfriRpz4z2kXeQAdfpJYP6b38dN8mUgFi3xL++Mm+/4b6xYWXK
			5vAvuF/K+NvSwIEOx9LtYT2lOg3285dhLzAWg64VziW9ol/Q8FDXUaG5patWm/0k
			SDehhdEQYMwL0EEU0qOGre7QOz3e0PK1H2PCbMHqsjZAw7QjyYqOxvWRWo2SohrA
			2QIDAQAB`,
		},
		{
			StudentID:   "B6400007",
			StudentName: "s7",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxTx8aPJOZVa/NpEj2sSK
			YrKcrS7UfuTvW+haK9Gcn0fJC3J6fJ9NXAK8wJD+VPL+mv1F5ElRrkDUjUPN2AvM
			NRyp6btXf8Nnmr8dD8GitdIt+8fLhRRZwq5xypIi2ngwBhOP7Sv8U6gwOCbgumFV
			wv4Go07K2Fnr4VectmoNhKmtItZRxov3FE7UmGlhFav5xDMj5q2eEc1HLXBBwSHw
			BOI7Wuz3eMtJHRXYws+rX0wHitGMWpgpLYIcoOE1zx2YWaLsEbDxAmBOs5tiKWTa
			iZElyjh9cCYCQdSpaZUU35xoR7ZH5tFZGXZVNiVfhidecJ55SJLj+5XLx+dk5Jd7
			pQIDAQAB`,
		},
		{
			StudentID:   "B6400008",
			StudentName: "s8",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApkd2RXy++Z4T3v6l41ne
			5AyXSIZW6zuim7j5d24m8BbW/Xc1OM8H6RoohlC5K0Eg/R9CTFycdR90p+ARbzK4
			1Wp6KvRzlDw/EeJr9M54aNPaQKB/IGUm7PdPXqZY8yLZ4ywEtSGSYgxTE/7+mFyh
			s3rKkO0r9cq/Rp2qgjliuXjI6OcMeExh380/qp28NJmNcgEpEcHYa7aafUr6WS8i
			zLztHlhvqVF+mSveiln/WWMxK5uGhgxQ3BtOsq0cwnhIy49MM5nwU1K9KNnqFA/m
			pffWIKJFwwQCfP8ZtsXxANAlOFenxtuAoWQ3QYo7CJNA9jjYadNOXYr4VPl53xIU
			2QIDAQAB`,
		},
		{
			StudentID:   "B6400009",
			StudentName: "s9",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt4bmJs7qPtPN6IJ0tAx6
			dr7kJxnDNmmRIs9yvtrQe1fZT/y1esVq9aPiW1c7bGZPExtpF0qyDE9P9iPb2nDv
			g2Ifdp93Vgjuo7GX21h3itnssHvyczPnNzSn+DS5iwYrMXwe03zuQHTMEKiHS25g
			hGUgbNReUtqXDjimmrj8wdhfDBtJiDcj4teocS6+2+HCTOzkmHg9CzY4Kr3nCYNZ
			bbj4uHo94xndUS/AvgaVZa3kYC1o/kOZSozD0eKtVdc3YgEs13IXv3kHO3473+CI
			eO2CHcqsILoHQqxT7iyHTlChk8nIQ/3aLV7nxFRiUJL6gZFUQ3ue04mJJbiNjvdB
			jwIDAQAB`,
		},
		{
			StudentID:   "B6400010",
			StudentName: "s10",
			PublicKey:  `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlAj3T/GiHFUqPSHokemc
			93VhFSC4brF2rDB29lTkVKPyoaRqIT+bx1uqgv/NqlFFcm7eVvQ9xYq+DhigbZ85
			RkJS4UKdq0aRPalzuSAiLi+N9E/y99TKy8ZWfWrbac0Ly2oGX6q2omg09fHA9tna
			7r2SeTnMqiaDCtefxMQe3FsdtFKMPET70lfdOrKoZMRSMeBq6jVTN9kQqw+aYPt2
			8oMbfrWUkBMrY8tyQJMYGX65DS7PkMBCT9YZu9q5Wmvt6LXyY80tivN9H7s2xHbP
			JmAUWNY5nXyGAKsZLWe0HCOUXM7+UKI75kLLFgx9uSCiCn34WKbmozrryvu+XflN
			TQIDAQAB`,
		},
	}

	for _, Voter := range Voter {
		db.Create(&Voter)
	}
	return database, nil
}
