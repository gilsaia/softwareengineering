package common

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SelfClaims struct {
	UserId int64
	jwt.StandardClaims
}

func GetToken(userId int64) (string, error) {
	claims := SelfClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: 7200,
			Issuer:    "VehicleManager",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString([]byte(tokenKey))
	return ss, err
}

func AuthToken(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if token != "abc" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token:%s", token)
	}
	return ctx, nil
}
