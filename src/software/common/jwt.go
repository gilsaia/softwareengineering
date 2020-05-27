package common

import "github.com/dgrijalva/jwt-go"

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
	ss, err := token.SignedString(tokenKey)
	return ss, err
}
