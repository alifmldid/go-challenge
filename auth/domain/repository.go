package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthRepository interface{
	FindUserById(c context.Context, authPayload AuthPayload) (res User, err error)
	GenerateToken(c context.Context, user User) (signedToken string, err error)
}

type authRepository struct{
	Conn *gorm.DB
}

func NewAuthRepository(Conn *gorm.DB) AuthRepository {
	return &authRepository{Conn}
}

func (repo *authRepository) FindUserById(c context.Context, authPayload AuthPayload) (res User, err error){
	user := repo.Conn.Where("username = ?", authPayload.Username).First(&res)

	return res, user.Error
}

func (repo *authRepository) GenerateToken(c context.Context, user User) (signedToken string, err error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err = token.SignedString([]byte("SECRET_NUMBER"))

	return signedToken, nil
}
