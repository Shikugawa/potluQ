package service

import (
	"io/ioutil"

	"os"

	"time"

	"crypto/rsa"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/role"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Client *ent.Client
}

var (
	privateKeyPath = os.Getenv("RSA_PRIVATE_VERIFIER_PATH")
)

func (service *JwtService) DefaultJwt(user *ent.User) (string, error) {
	privKey, err := service.prepareRsaPrivateKey()
	if err != nil {
		return *new(string), err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_name"] = user.Name
	claims["role"] = role.DefaultRole
	claims["expires_at"] = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) UpdateRole(token *jwt.Token, club *ent.Club, role role.Role) (string, error) {
	privKey, err := service.prepareRsaPrivateKey()
	if err != nil {
		return *new(string), err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["club_name"] = club.Name
	claims["role"] = role

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) prepareRsaPrivateKey() (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	return signKey, nil
}
