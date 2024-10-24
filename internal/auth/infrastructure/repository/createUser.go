package infraSqlxRepository

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (a *ImplSqlx) CreateUserAccount(payload *authDto.RegisterDb, db *sqlx.DB) error {
	query := `INSERT INTO "public".users 
	(id, first_name, last_name, password, email, phone_number, tax_id, wallet_address, identity_document_url, 
	is_uiff, is_exposed, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`

	_, err := db.Exec(query, payload.ID, payload.FirstName, payload.LastName, payload.Password, payload.Email, payload.PhoneNumber, payload.TaxID,
		payload.WalletAddress, payload.IdentityDocument, payload.IsUIFF, payload.IsExposed, payload.Role, payload.CreatedAt, payload.UpdatedAt) // Se incluye UpdatedAt
	if err != nil {
		return err
	}

	return nil
}
