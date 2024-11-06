package authDto

type RegisterCompanyRequest struct {
	BusinessName            string `json:"business_name"`
	Email                   string `json:"email"`
	LegalRepresentativeName string `json:"legal_representative_name"`
	LegalRepresentativeID   string `json:"legal_representative_id"`
	Password                string `json:"password"`
	PhoneNumber             string `json:"phone_number"`
	TaxID                   string `json:"tax_id"`
	Address                 string `json:"address"`
	CompanyCertificateURL   string `json:"company_certificate_url"`
}

type RegisterCompanyDB struct {
	ID                      string `json:"id"`
	BusinessName            string `json:"business_name"`
	Email                   string `json:"email"`
	LegalRepresentativeName string `json:"legal_representative_name"`
	LegalRepresentativeID   string `json:"legal_representative_id"`
	Password                string `json:"password"`
	PhoneNumber             string `json:"phone_number"`
	TaxID                   string `json:"tax_id"`
	Address                 string `json:"address"`
	CompanyCertificateURL   string `json:"company_certificate_url"`
	Role                    int    `json:"role"`
	CreatedAt               int64  `json:"created_at"`
	UpdatedAt               int64  `json:"updated_at"`
}
