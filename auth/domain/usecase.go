package domain

import (
	"context"
)

type AuthUsecase interface{
	Login(c context.Context, authPayload AuthPayload) (code int, message string, err error)
}

type authUsecase struct{
	authRepository AuthRepository
}

func NewAuthUsecase(authRepository AuthRepository) AuthUsecase{
	return &authUsecase{authRepository}
}

func (uc *authUsecase) Login(c context.Context, authPayload AuthPayload) (code int, message string, err error) {
	user, err := uc.authRepository.FindUserById(c, authPayload)

	if err != nil {
		return 500, "", err
	}

	if user.Username == "" {
		return 404, "username tidak ditemukan", err 
	}

	if user.Password != authPayload.Password {
		return 400, "password tidak sesuai", nil
	}
	
	token, err := uc.authRepository.GenerateToken(c, user)

	return 200, token, nil
}
