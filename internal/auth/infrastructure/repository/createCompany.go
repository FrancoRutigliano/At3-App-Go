package infraSqlxRepository

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (a *ImplSqlx) CreateCompanyAccount(payload *authDto.RegisterCompanyDB, db *sqlx.DB) error {
	query := `INSERT INTO "public".companies 
	(id, business_name, email, legal_representative_name, legal_representative_id, password, 
	phone_number, tax_id, address, company_certificate_url, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);`

	_, err := db.Exec(query,
		payload.ID,
		payload.BusinessName,
		payload.Email,
		payload.LegalRepresentativeName,
		payload.LegalRepresentativeID,
		payload.Password,
		payload.PhoneNumber,
		payload.TaxID,
		payload.Address,
		payload.CompanyCertificateURL,
		payload.Role,
		payload.CreatedAt,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
