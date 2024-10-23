package authDto

import (
	"time"
)

type RegisterUser struct {
	ID               string    `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Password         string    `json:"-"`
	PhoneNumber      string    `json:"phone_number"`
	TaxID            string    `json:"tax_id"`
	WalletAddress    string    `json:"wallet_address"`
	IdentityDocument string    `json:"identity_document_url"`
	IsUIFF           bool      `json:"is_uiff"`
	IsExposed        bool      `json:"is_exposed"`
	Role             int       `json:"role"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
