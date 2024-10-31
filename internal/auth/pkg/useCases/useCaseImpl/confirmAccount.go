package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	httpresponse "at3-back/pkg/httpResponse"
	"context"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func (a *Auth) Confirm(tokenS string) httpresponse.ApiResponse {
	var ctx = context.Background()
	token, err := a.JwtService.ValidateTokenFromQuery(tokenS)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	id, err := a.JwtService.GetUUIdFromToken(token)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	var user authDto.RegisterDb

	userJson, err := a.Redis.Get(ctx, "uuid:"+id).Result()
	if err == redis.Nil {
		return *httpresponse.NewApiError(http.StatusNotFound, "User not found or expired", nil)
	} else if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}

	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}

	err = a.Repository.Impl.CreateUserAccount(&user, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}

	err = a.Redis.Del(ctx, "uuid:"+id).Err()
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, "user validated", nil)
}
