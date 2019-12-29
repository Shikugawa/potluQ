package service

import (
	"io/ioutil"

	"os"

	"context"

	"time"

	"crypto/rsa"

	"github.com/Shikugawa/potraq/ent"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Client *ent.Client
}

var (
	privateKeyPath = os.Getenv("RSA_PRIVATE_VERIFIER_PATH")
)

func (service *JwtService) CreateRequesterJwt(ctx context.Context, user *ent.User) (string, error) {
	privKey, err := service.prepareRsaPrivateKey(ctx, user)
	if err != nil {
		return *new(string), err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.StandardClaims)
	claims.Id = user.UserID
	claims.Subject = "Token is used for register music only"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) CreateJukeboxJwt(ctx context.Context, user *ent.User) (string, error) {
	privKey, err := service.prepareRsaPrivateKey(ctx, user)
	if err != nil {
		return *new(string), err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.StandardClaims)
	claims.Id = user.UserID
	claims.Subject = "Token is used for jukebox that allows to get music in addition to requester service"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) prepareRsaPrivateKey(ctx context.Context, user *ent.User) (*rsa.PrivateKey, error) {
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
