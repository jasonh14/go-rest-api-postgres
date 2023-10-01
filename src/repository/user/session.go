package user

import (
	"app/src/model"
	"app/src/tracing"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	ctx, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()

	accessClaims := Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ur.accessExp).Unix(),
			Subject:   userID,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)
	return accessToken.SignedString(ur.signKey)

}

func (ur *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	accessToken, err := ur.generateAccessToken(ctx, userID)
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")

	defer span.End()
	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (ur *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return &ur.signKey.PublicKey, nil
	})

	if err != nil {
		return "", err
	}

	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}

	return "", errors.New("unauthorized")
}
