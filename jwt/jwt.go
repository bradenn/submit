package jwt

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
)

var Auth *jwtauth.JWTAuth

func Init() {
	Auth = jwtauth.New("HS256", []byte("b753b0d5-47a1-454b-a93c-9cad11ab663f"), nil)
}

type UserToken struct {
	UserId string
	Token  string
}

func NewAuthToken(userToken UserToken) {
	_, tokenString, _ := Auth.Encode(map[string]interface{}{
		"userId": userToken.UserId,
		"token":  userToken.Token,
	})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}
