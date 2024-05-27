package sample_data

type account struct {
	IBAN    string `json:"iban"`
	BIC     string `json:"bic"`
	Comment string `json:"comment"`
}

type client struct {
	Name     string    `json:"name"`
	Address  [5]string `json:"address"`
	Accounts []account `json:"accounts"`
}

type balance struct {
	Product        string  `json:"product"`
	OpeningBalance float64 `json:"opening_balance"`
	Income         float64 `json:"income"`
	Outcome        float64 `json:"outcome"`
	Balance        float64 `json:"balance"`
}

type description struct {
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Account string `json:"account"`
}

type transaction struct {
	Date        string      `json:"date"`
	Description description `json:"description"`
	Withdraw    float64     `json:"withdraw"`
	Deposit     float64     `json:"deposit"`
	Balance     float64     `json:"balance"`
}

type BankAccount struct {
	Client            client        `json:"client"`
	Balance           []balance     `json:"balance"`
	Transactions      []transaction `json:"transactions"`
	TransactionsRange string        `json:"transactions_range"`
	GeneratedDate     string        `json:"generated_date"`
}

func GetBankAccountData() BankAccount {
	data := BankAccount{
		Client: client{
			Name:    "SANDRA SAULGRIEZE",
			Address: [5]string{"14 The Dale", "Whitefield hall", "Bettystown", "Meath", "A92N27C"},
			Accounts: []account{
				{
					IBAN:    "IE30REVO99036022547749",
					BIC:     "REVOIE23",
					Comment: "",
				},
				{
					IBAN:    "LT093250041069208595",
					BIC:     "REVOLT21",
					Comment: "(You cannot use this IBAN for bank transfers. Please use the IBAN found in the app)",
				},
				{
					IBAN:    "LT087070024346246713",
					BIC:     "RETBLT21",
					Comment: "(You cannot use this IBAN for bank transfers. Please use the IBAN found in the app)",
				},
			},
		},
		Balance: []balance{
			{
				Product:        "Account (Current Account)",
				OpeningBalance: 2.52,
				Income:         1944.09,
				Outcome:        1978.00,
				Balance:        36.43,
			},
			{
				Product:        "Total",
				OpeningBalance: 2.52,
				Income:         1944.09,
				Outcome:        1978.00,
				Balance:        36.43,
			},
		},
		Transactions: []transaction{
			{
				Date: "3 Feb 2023",
				Description: description{
					Title:   "Apple Pay Top-Up by *5453",
					From:    "",
					To:      "",
					Account: "",
				},
				Withdraw: 0.00,
				Deposit:  50.00,
				Balance:  52.52,
			},
			{
				Date: "3 Feb 2023",
				Description: description{
					Title:   "Apple Pay Top-Up by *5453",
					From:    "",
					To:      "",
					Account: "",
				},
				Withdraw: 0.00,
				Deposit:  100.00,
				Balance:  152.52,
			},
			{
				Date: "3 Feb 2023",
				Description: description{
					Title:   "To LINA MILLER SAULGRIEZE",
					From:    "",
					To:      "LINA MILLER SAULGRIEZE",
					Account: "Card: 416598******8124",
				},
				Withdraw: 100.00,
				Deposit:  0.00,
				Balance:  52.52,
			},
			{
				Date: "7 Feb 2023",
				Description: description{
					Title:   "To LINA MILLER SAULGRIEZE",
					From:    "",
					To:      "LINA MILLER SAULGRIEZE",
					Account: "",
				},
				Withdraw: 10.00,
				Deposit:  00.00,
				Balance:  42.52,
			},
		},
		TransactionsRange: "1 February 2023 to 29 March 2023",
		GeneratedDate:     "27 May 2024",
	}

	return data
}
