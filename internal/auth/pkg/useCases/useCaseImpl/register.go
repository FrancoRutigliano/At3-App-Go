package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"at3-back/pkg/validator"
	"net/http"

	"github.com/google/uuid"
)

func (a *Auth) Register(payload authDto.RegisterUser) httpresponse.ApiResponse {
	exists, err := a.Repository.Impl.FindByEmail(payload.Email, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	if exists {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email already exist", nil)
	}

	uuid := uuid.New().String()

	hashed, err := hash.HashPassword(payload.Password)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}

	// parsear fechas
	createdUnix := validator.DateToUnix()
	updatedUnix := validator.DateToUnix()

	user := &authDto.RegisterDb{
		ID:               uuid,
		FirstName:        payload.FirstName,
		LastName:         payload.LastName,
		Email:            payload.Email,
		Password:         hashed,
		PhoneNumber:      payload.PhoneNumber,
		TaxID:            payload.TaxID,
		WalletAddress:    payload.WalletAddress,
		IdentityDocument: payload.IdentityDocument,
		IsUIFF:           payload.IsUIFF,
		IsExposed:        payload.IsExposed,
		Role:             1,
		CreatedAt:        createdUnix,
		UpdatedAt:        updatedUnix,
	}

	err = a.Repository.Impl.CreateUserAccount(user, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	token, err := a.JwtService.GenerateTokenRegister(uuid)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	err = a.EmailService.SendRegisterEmail(payload.Email, token)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusServiceUnavailable, "Service error: Unavailable sending email", nil)
	}

	return *httpresponse.NewApiError(http.StatusCreated, "Registration successfully, please check your email to confirm your account. ", token)
}
