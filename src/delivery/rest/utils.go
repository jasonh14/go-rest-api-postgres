package rest

import (
	"app/src/model"
	"errors"
	"net/http"
	"strings"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("Invalid authorization")
	}

	return model.UserSession{
		JWTToken: splitString[1],
	}, nil
}
