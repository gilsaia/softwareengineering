package common

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type SelfClaims struct {
	UserId     int64
	Permission int
	jwt.StandardClaims
}

func GetToken(userId int64, permission int) (string, error) {
	claims := SelfClaims{
		userId,
		permission,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "VehicleManager",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString([]byte(tokenKey))
	return ss, err
}

func AuthToken(ctx context.Context) (context.Context, error) {
	tokenStr, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token:%s", token)
	}
	if claims, ok := token.Claims.(*SelfClaims); ok {
		ctx = context.WithValue(ctx, "userId", claims.UserId)
		ctx = context.WithValue(ctx, "permission", claims.Permission)
	}
	return ctx, nil
}

func AuthPermission(ctx context.Context, requiredPermission int) BgErr {
	permission, _ := ctx.Value("permission").(int)
	if permission >= requiredPermission {
		return Success
	}
	return BgErr{ErrNo: 40002,
		ErrMsg: string(permission)}
}
