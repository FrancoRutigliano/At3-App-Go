package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"at3-back/pkg/validator"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (a *Auth) Register(payload authDto.RegisterUser) httpresponse.ApiResponse {

	exists, err := a.Repository.Impl.FindByEmail(strings.ToLower(payload.Email), "users", a.Db)
	if err != nil {
		log.Println("Failed to User findByEmail: Email:", payload.Email, ", Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}
	if exists {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email already exist", nil)
	}

	uuid := uuid.New().String()

	hashed, err := hash.HashPassword(payload.Password)
	if err != nil {
		log.Println("Failed to hash password: Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	// parsear fechas
	createdUnix := validator.DateToUnix()
	updatedUnix := validator.DateToUnix()

	user := &authDto.RegisterDb{
		ID:               uuid,
		FirstName:        strings.ToLower(payload.FirstName),
		LastName:         strings.ToLower(payload.LastName),
		Email:            strings.ToLower(payload.Email),
		Password:         hashed,
		PhoneNumber:      payload.PhoneNumber,
		TaxID:            payload.TaxID,
		WalletAddress:    payload.WalletAddress,
		IdentityDocument: payload.IdentityDocument,
		Country:          strings.ToLower(payload.Country),
		PostalCode:       payload.PostalCode,
		Address:          strings.ToLower(payload.Address),
		AddressNumber:    payload.AddressNumber,
		IsUIFF:           payload.IsUIFF,
		IsExposed:        payload.IsExposed,
		Role:             1,
		CreatedAt:        createdUnix,
		UpdatedAt:        updatedUnix,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println("Failed to marshal user JSON: Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	err = a.Redis.Set(a.Ctx, "uuid:"+user.ID, userJson, 24*time.Hour).Err()
	if err != nil {
		log.Println("Redis error Set user: ID", user.ID, "Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	token, err := a.JwtService.GenerateTokenRegister(map[string]interface{}{
		"uuid": uuid,
	})
	if err != nil {
		log.Println("Failed to generate token: Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	err = a.EmailService.SendRegisterEmail(payload.Email, token, "user")
	if err != nil {
		log.Println("Failed service email: Error:", err)
		return *httpresponse.NewApiError(http.StatusServiceUnavailable, "Service error: Unavailable sending email", nil)
	}

	return *httpresponse.NewApiError(http.StatusCreated, "Registration successfully, please check your email to confirm your account. ", token)
}
