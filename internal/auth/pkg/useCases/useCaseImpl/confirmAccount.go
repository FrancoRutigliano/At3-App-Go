package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	httpresponse "at3-back/pkg/httpResponse"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func (a *Auth) Confirm(tokenS, entity string) httpresponse.ApiResponse {
	var user authDto.RegisterDb
	var company authDto.RegisterCompanyDB
	token, err := a.JwtService.ValidateTokenFromQuery(tokenS)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	id, err := a.JwtService.GetUUIdFromToken(token)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	switch entity {
	case "user":
		userJson, err := a.Redis.Get(a.Ctx, "uuid:"+id).Result()
		if err == redis.Nil {
			return *httpresponse.NewApiError(http.StatusNotFound, "User not found or expired", nil)
		} else if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "can't get user", nil)
		}

		err = json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
		}

		err = a.Repository.Impl.CreateUserAccount(&user, a.Db)
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "can't create user on db", nil)
		}

		err = a.Redis.Del(a.Ctx, "uuid:"+id).Err()
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
		}
	case "company":

		companyJson, err := a.Redis.Get(a.Ctx, "companyuuid:"+id).Result()
		if err == redis.Nil {
			return *httpresponse.NewApiError(http.StatusNotFound, "User not found or expired", nil)
		} else if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "can't get user", nil)
		}

		err = json.Unmarshal([]byte(companyJson), &company)
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
		}

		err = a.Repository.Impl.CreateCompanyAccount(&company, a.Db)
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "can't create user on db", nil)
		}

		err = a.Redis.Del(a.Ctx, "companyuuid:"+id).Err()
		if err != nil {
			return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
		}
	}
	return *httpresponse.NewApiError(http.StatusOK, "user validated", nil)
}
