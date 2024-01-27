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
			StudentID:   "B6400001",
			StudentName: "s1",
			PublishKey: `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCRl3fhbfVHn8+45qhNaVV3ApaL
K6iQZ+mpwY1Cme7iSXYntbQCvhpXxILDnYEGGJu1Yndtgn+HQMnJH3lAeb1vDuOl
VJF34wkQcwqAVuxE0QS92mJvTIVhU89TO+mFMzYV5Vklj2Smg6M/fQDIvLBIUJGR
QTRT7IK5bYSAo+RqXwIDAQAB`,
		},
		{
			StudentID:   "B6400002",
			StudentName: "s2",
			PublishKey: `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGWLIXVKNsaqTTmx1w9re3TuzlIb
MR5VM/icmZcl6Wjd+RIEko5QvzKp5Hv5pJEknjEDR9SQAHFvWIUGf1GCk596yCic
Rxgxrkmxm21as4bBdnUx40IjWsuxnVNRvzuoGQG7h5I3qJONiXnUqe7bmk/2sKoc
xbM7pxscFN4tC1dlAgMBAAE=
`,
		},
		{
			StudentID:   "B6400003",
			StudentName: "s3",
			PublishKey:  `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGPm6Ipihgcpi9HiKJ54z2LNgEOQ
5sk27UUhEAP+6rWAohqmZ+TTHSWgAygrXx89oLpZ9mGwy+it6fELUm8JVtlk1l/J
AnuDjzPscYh0P+T2zfb1P837YGEdEaI4NwFVo4nD6u0860nf9fExy32VHb4fuPUq
37r4n+NnQsyPpKwzAgMBAAE=
`,
		},
		{
			StudentID:   "B6400004",
			StudentName: "s4",
			PublishKey:  `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGPm6Ipihgcpi9HiKJ54z2LNgEOQ
5sk27UUhEAP+6rWAohqmZ+TTHSWgAygrXx89oLpZ9mGwy+it6fELUm8JVtlk1l/J
AnuDjzPscYh0P+T2zfb1P837YGEdEaI4NwFVo4nD6u0860nf9fExy32VHb4fuPUq
37r4n+NnQsyPpKwzAgMBAAE=
`,
		},
		{
			StudentID:   "B6400005",
			StudentName: "s5",
			PublishKey:  `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGPm6Ipihgcpi9HiKJ54z2LNgEOQ
5sk27UUhEAP+6rWAohqmZ+TTHSWgAygrXx89oLpZ9mGwy+it6fELUm8JVtlk1l/J
AnuDjzPscYh0P+T2zfb1P837YGEdEaI4NwFVo4nD6u0860nf9fExy32VHb4fuPUq
37r4n+NnQsyPpKwzAgMBAAE=
`,
		},
		{
			StudentID:   "B6400006",
			StudentName: "s6",
			PublishKey:  `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDa7klMlKEa//gL1ly+jD2iV/mj
2gVbjzJsWLK2jQRhUUi0mVcZGbRoJmoZUsQEXIkXCUUzemsSaOSEIcFvl2ZN54Nm
gvV/JNKxohdoRscjqAQF9VCrVHJXlLzu/0S9udqq6JJnF02T6vGlbE3oI/qBIXS+
lLXGENakHECtfTQW/QIDAQAB
`,
		},
		{
			StudentID:   "B6400007",
			StudentName: "s7",
			PublishKey:  `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGq14iacCooDuS7fxApO2vBrTFAS
/C64Gnego/QbVo+Z8AbvCUBa1sncdazK/L55coANiATKhF/Zn0+5ywCsQhSY8ZAp
3S2nR1QBra/DltojAVja6bBOGA0ES33HSvSqFG6XZ8oVB5medkMKLoK8F/NC5Ccu
RyaqDj7eau1xJGLtAgMBAAE=
`,
		},
		{
			StudentID:   "B6400008",
			StudentName: "s8",
			PublishKey:  `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDA4sx4VWpY8DZWmGlwgTiA65q
h9D135jQG5K9rTOgzjK8BcT/+kdgbK4rEXH3L97Xf/oxwuAKzwvXSAdTAGEjFucD
rGMM+jBtJgo3mpDJkynB3w/9KGLrcMk7Cq3Hmu1MWTd8mGbUnHi0lgzwyPp0KLO9
5DQMrTFnLgMxNb2fZwIDAQAB
`,
		},
		{
			StudentID:   "B6400009",
			StudentName: "s9",
			PublishKey:  `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDxOdR9GtkPlAQP0DqcQCmJuIb6
lDROc4a1D5BQsmrqJEV1qc5WRFH7J4jSNV5+zEnoVvuUWK+WIy8Z1q9caNpeEYCT
aPOxKRr4mauoVkPnqbBOZTs9mumCG9MyzzbY185japHBhQK02EvdRXOSPgE8vcJl
n0X1UnKr9L446kx8vQIDAQAB
`,
		},
		{
			StudentID:   "B6400010",
			StudentName: "s10",
			PublishKey:  `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHP95VxYmLggJlh2oWpCii6MMg
NC9pajZ73aszTf7srmFJvF8tx4l/FFripT/Rq08GLzhPObx5Ojzyy+i1DK2IgVHZ
58lOmfqT3x9lKEP0/AsAhlRvdttdf2d3GqIu6powYr5FtxWTxTKTnR/HGGJUapoz
fq7KRjqYR3vWz2x41QIDAQAB
`,
		},
	}

	for _, voter := range voter {
		db.Create(&voter)
	}
	return database, nil
}
