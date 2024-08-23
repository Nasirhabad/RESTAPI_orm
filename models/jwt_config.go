package models

type JWTOptions struct {
	SecretKey      string
	ExpireDuration int
}
