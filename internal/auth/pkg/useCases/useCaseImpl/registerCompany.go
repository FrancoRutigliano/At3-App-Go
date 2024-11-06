package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"at3-back/pkg/validator"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (a *Auth) RegisterCompany(payload authDto.RegisterCompanyRequest) httpresponse.ApiResponse {
	var ctx = context.Background()

	exist, err := a.Repository.Impl.FindByEmail(payload.Email, "companies", a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}
	if exist {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email already exist", nil)
	}

	id := uuid.New().String()

	hashedPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	createdAt := validator.DateToUnix()
	updatedAt := validator.DateToUnix()

	var companyDb = authDto.RegisterCompanyDB{
		ID:                      id,
		BusinessName:            payload.BusinessName,
		Email:                   payload.Email,
		LegalRepresentativeName: payload.LegalRepresentativeName,
		LegalRepresentativeID:   payload.LegalRepresentativeID,
		Password:                hashedPassword,
		PhoneNumber:             payload.PhoneNumber,
		TaxID:                   payload.TaxID,
		Address:                 payload.Address,
		CompanyCertificateURL:   payload.CompanyCertificateURL,
		Role:                    4,
		CreatedAt:               createdAt,
		UpdatedAt:               updatedAt,
	}

	companyJson, err := json.Marshal(companyDb)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	err = a.Redis.Set(ctx, "uuid"+id, companyJson, 24*time.Hour).Err()
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	token, err := a.JwtService.GenerateTokenRegister(map[string]interface{}{
		"uuid": id,
	})
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	err = a.EmailService.SendRegisterEmail(payload.Email, token, "company")
	if err != nil {
		return *httpresponse.NewApiError(http.StatusServiceUnavailable, "Service error: Unavailable sending email", nil)
	}

	return *httpresponse.NewApiError(http.StatusCreated, "Registration successfully, please check your email to confirm your account. ", token)
}
