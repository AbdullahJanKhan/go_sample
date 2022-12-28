package service

import (
	"errors"
	"time"

	"github.com/abdullahjankhan/go_sample/models"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	CreateToken(email, code string) (string, error)
	VerifyToken(tokenStr string) (string, string, error)
}
type jwtService struct {
	configService GlobleConfigService
}

func NewJWTService(configService GlobleConfigService) JWTService {
	return &jwtService{
		configService: configService,
	}
}
func (j *jwtService) CreateToken(email, code string) (string, error) {

	claim := jwt.MapClaims{
		"email": email,
		"code":  code,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := j.configService.GetConfig().JwtSecret
	return token.SignedString([]byte(secret))
}

func (j *jwtService) VerifyToken(tokenStr string) (string, string, error) {

	secret := j.configService.GetConfig().JwtSecret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		stdErr := &models.StandardError{
			Code:        models.INVALID_TOKEN,
			ActualError: err,
			Line:        "VerifyToken():126",
			Message:     models.INVALID_TOKEN_MESSAGE,
		}
		return "", "", stdErr
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("cannot convert claim to MapClaims")
	}
	if !token.Valid {
		return "", "", errors.New("token is invalid")
	}

	emailVal, found := claim["email"]
	if !found {
		return "", "", errors.New("bad token")
	}
	codeVal, found := claim["code"]
	if !found {
		return "", "", errors.New("bad token")
	}
	code := codeVal.(string)
	email := emailVal.(string)

	return email, code, nil
}
