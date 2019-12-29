package middleware

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

var (
	publicKeyPath = os.Getenv("RSA_PUBLIC_VERIFIER_PATH")
)

type Authenticator struct {
}

func InitAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (auth *Authenticator) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubKey, err := auth.pubKey()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			_, err := token.Method.(*jwt.SigningMethodRSA)
			if !err {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return pubKey, nil
			}
		})

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	}
}

func (auth *Authenticator) pubKey() (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}
