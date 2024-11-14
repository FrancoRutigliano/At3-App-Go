package authDto

type RegisterUser struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number"`
	TaxID            string `json:"tax_id"`
	WalletAddress    string `json:"wallet_address"`
	IdentityDocument string `json:"identity_document_url"`
	Country          string `json:"country"`
	PostalCode       string `json:"postal_code"`
	Address          string `json:"address"`
	AddressNumber    int    `json:"address_number"`
	IsUIFF           bool   `json:"is_uiff"`
	IsExposed        bool   `json:"is_exposed"`
}

type RegisterDb struct {
	ID               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number"`
	TaxID            string `json:"tax_id"`
	WalletAddress    string `json:"wallet_address"`
	IdentityDocument string `json:"identity_document_url"`
	Country          string `json:"country"`
	PostalCode       string `json:"postal_code"`
	Address          string `json:"address"`
	AddressNumber    int    `json:"address_number"`
	IsUIFF           bool   `json:"is_uiff"`
	IsExposed        bool   `json:"is_exposed"`
	Role             int    `json:"role"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}
