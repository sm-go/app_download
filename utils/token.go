package utils

import (
	"app-download/model"
	"crypto/rsa"
	"errors"
	"log"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
)

type UserAccessTokenClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

type UserRefreshTokenClaims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateUserAccessToken(user *model.User, key *rsa.PrivateKey) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 3600*24

	claims := UserAccessTokenClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Println("Failed to sign id token string")
		return "", err
	}
	return ss, nil
}

func GenerateUserRefreshToken(user *model.User, key string) (*model.RefreshTokenData, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(time.Duration(60*60*24*3) * time.Second)
	tokenID, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate refresh token ID")
		return nil, err
	}

	claims := UserRefreshTokenClaims{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("Failed to sign refresh token string")
		return nil, err
	}

	return &model.RefreshTokenData{
		SS:        ss,
		ID:        tokenID,
		ExpiresIn: tokenExp.Sub(currentTime),
	}, nil
}

func ValidateUserAccessToken(tokenString string, key *rsa.PublicKey) (*UserAccessTokenClaims, error) {
	claims := &UserAccessTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("access token is invalid")
	}

	claims, ok := token.Claims.(*UserAccessTokenClaims)

	if !ok {
		return nil, errors.New("access token valid but couldn't parse claims")
	}

	return claims, nil
}

func ValidateUserRefreshToken(tokenString string, key string) (*UserRefreshTokenClaims, error) {
	claims := &UserRefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("refresh token is invalid")
	}

	claims, ok := token.Claims.(*UserRefreshTokenClaims)
	if !ok {
		return nil, errors.New("refresh token valid but couldn't parse claims")
	}

	return claims, nil
}
