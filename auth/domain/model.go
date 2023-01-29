package domain

import "github.com/golang-jwt/jwt"

type AuthPayload struct{
	Username string
	Password string
}

type User struct{
	Id int
	Username string
	Password string
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Id int `json:"id"`
}