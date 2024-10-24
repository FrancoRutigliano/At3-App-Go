package authUseCaseImpl

import (
	httpresponse "at3-back/pkg/httpResponse"
	"net/http"
)

func (a *Auth) Confirm(tokenS string) httpresponse.ApiResponse {
	token, err := a.JwtService.ValidateTokenFromQuery(tokenS)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	id, err := a.JwtService.GetUUIdFromToken(token)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	exist, err := a.Repository.Impl.FindByIdUpdate(id, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}
	if !exist {
		return *httpresponse.NewApiError(http.StatusBadRequest, "id doesn't exist", nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, "user validated", nil)
}
