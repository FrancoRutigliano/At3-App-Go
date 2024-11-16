package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	httpresponse "at3-back/pkg/httpResponse"
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func (a *Auth) Confirm(tokenS, entity string) httpresponse.ApiResponse {
	var user authDto.RegisterDb
	var company authDto.RegisterCompanyDB

	// Validar el token
	token, err := a.JwtService.ValidateTokenFromQuery(tokenS)
	if err != nil {
		log.Println("Error validating token from query:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	// Obtener UUID del token
	id, err := a.JwtService.GetUUIdFromToken(token)
	if err != nil {
		log.Println("Error extracting UUID from token:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	// Manejo de la entidad
	switch entity {
	case "user":
		userJson, err := a.Redis.Get(a.Ctx, "uuid:"+id).Result()
		if err == redis.Nil {
			log.Println("User not found or expired: ID", id)
			return *httpresponse.NewApiError(http.StatusNotFound, "User not found or expired", nil)
		} else if err != nil {
			log.Println("Redis error retrieving user: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Deserializar JSON
		err = json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			log.Println("Failed to unmarshal user JSON: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Crear usuario en la DB
		err = a.Repository.Impl.CreateUserAccount(&user, a.Db)
		if err != nil {
			log.Println("DB error creating user account: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Eliminar clave en Redis
		err = a.Redis.Del(a.Ctx, "uuid:"+id).Err()
		if err != nil {
			log.Println("Redis error deleting user key: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

	case "company":
		companyJson, err := a.Redis.Get(a.Ctx, "companyuuid:"+id).Result()
		if err == redis.Nil {
			log.Println("Company not found or expired: ID", id)
			return *httpresponse.NewApiError(http.StatusNotFound, "Company not found or expired", nil)
		} else if err != nil {
			log.Println("Redis error retrieving company: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Deserializar JSON
		err = json.Unmarshal([]byte(companyJson), &company)
		if err != nil {
			log.Println("Failed to unmarshal company JSON: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Crear compañía en la DB
		err = a.Repository.Impl.CreateCompanyAccount(&company, a.Db)
		if err != nil {
			log.Println("DB error creating company account: ID", id, ", Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}

		// Eliminar clave en Redis
		err = a.Redis.Del(a.Ctx, "companyuuid:"+id).Err()
		if err != nil {
			log.Println("Redis error deleting company key: ID", id, "Error:", err)
			return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
		}
	}

	return *httpresponse.NewApiError(http.StatusOK, "User validated", nil)
}
