package middleware

import (
	"encoding/json"
	"os"
	"petshop-backend/models"
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
)

func EncodeWithRoleHours(role, email string, hours int64) (string, error) {
	privatekey := os.Getenv("PRIVATEKEY")
	token := paseto.NewToken()
	// Set metadata: waktu pembuatan, masa berlaku, dll
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(time.Duration(hours) * time.Hour))
	token.SetString("email", email)
	token.SetString("role", role)
	key, err := paseto.NewV4AsymmetricSecretKeyFromHex(privatekey)
	return token.V4Sign(key, nil), err
}

func Decoder(tokenstr string) (payload models.Payload, err error) {
	publickey := os.Getenv("PUBLICKEY")

	// Remove "Bearer " prefix if present
	if strings.HasPrefix(tokenstr, "Bearer ") {
		tokenstr = strings.TrimPrefix(tokenstr, "Bearer ")
	}

	pubKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publickey)
	if err != nil {
		return payload, err
	}

	parser := paseto.NewParser()
	token, err := parser.ParseV4Public(pubKey, tokenstr, nil)
	if err != nil {
		return payload, err
	}

	json.Unmarshal(token.ClaimsJSON(), &payload)

	return payload, err
}
