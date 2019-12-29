package service

import (
	"io/ioutil"

	"os"

	"context"

	"errors"

	"time"

	"crypto/rsa"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/message"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	client *ent.Client
}

var (
	privateKeyPath = os.Getenv("RSA_PRIVATE_VERIFIER_PATH")
)

func (service *JwtService) CreateRequesterJwt(ctx context.Context, credential *message.Credential) (string, error) {
	privKey, err := service.prepareRsaPrivateKey(ctx, credential)
	if err != nil {
		return *new(string), err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.StandardClaims)
	claims.Subject = "Token is used for register music only"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) CreateJukeboxJwt(ctx context.Context, credential *message.Credential) (string, error) {
	privKey, err := service.prepareRsaPrivateKey(ctx, credential)
	if err != nil {
		return *new(string), err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.StandardClaims)
	claims.Subject = "Token is used for jukebox that allows to get music in addition to requester service"
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString(privKey)
	if err != nil {
		return *new(string), err
	}
	return tokenstr, nil
}

func (service *JwtService) prepareRsaPrivateKey(ctx context.Context, credential *message.Credential) (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	userserv := UserService{
		client: service.client,
	}
	result := userserv.VerityPassword(ctx, credential)
	if !result {
		return nil, errors.New("can't create vaild token") // TODO: enrich message
	}

	return signKey, nil
}
