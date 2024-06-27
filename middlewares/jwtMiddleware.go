package middlewares

import (
	"TugasMID2/constants"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
