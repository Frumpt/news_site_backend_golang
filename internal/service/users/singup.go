package users

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	singingKey = "sajfafadssjfklas"
	tokenTTL   = 12 * time.Hour
)

type tokenClimes struct {
	jwt.StandardClaims
	UserId uint
}

func generateToken(password string, UserId uint) (string, error) {

	_, data, err := findUser(uint(UserId), password)

	if err == nil || data == nil {
		return "", fmt.Errorf("user not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&tokenClimes{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			UserId,
		})

	return token.SignedString([]byte(singingKey))
}

func parseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClimes{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid, singing method")
		}

		return []byte(singingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClimes)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClimes")
	}

	return claims.UserId, nil
}

func findUser(id uint, password string) (int64, []byte, error) {

	var user models.Users

	passwordHash := PasswordHasher(password)

	res := db.DataBase.Select("Id", "Name", "Password").Find(&user, "id = ?", id)

	if user.Password == &passwordHash {

		data, err := json.Marshal(user)

		if err != nil {
			return 0, []byte{}, err
		}

		return res.RowsAffected, data, res.Error
	}

	return 0, []byte{}, fmt.Errorf("login or password not true")
}

type JsonDataSing struct {
	Id       uint   `json:"Id"`
	Password string `json:"Password"`
}

func SingUp(body []byte) ([]byte, error) {
	var JsonData JsonDataSing
	err := json.Unmarshal(body, &JsonData)
	if err != nil {
		return nil, err
	}

	token, errGener := generateToken(JsonData.Password, JsonData.Id)

	if errGener != nil {
		return nil, errGener
	}

	data, errJSON := json.Marshal(token)

	if errJSON != nil {
		return nil, errJSON
	}

	return data, nil
}
