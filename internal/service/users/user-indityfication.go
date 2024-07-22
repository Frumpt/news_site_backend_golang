package users

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const (
	authorizationHeader = "Authorization"
)

func UserIndifity(c fiber.Ctx) error {

	HeadersReq := c.GetReqHeaders()

	Tokens := HeadersReq[authorizationHeader]

	if len(Tokens) == 0 {
		return c.SendStatus(401)
	}

	TokenAuth := Tokens[0]

	PartsToken := strings.Split(TokenAuth, " ")

	if len(PartsToken) != 2 {
		fmt.Println(PartsToken)
		return c.SendStatus(401)
	}

	userId, err := parseToken(PartsToken[1])
	if err != nil {
		return c.SendStatus(401)
	}
	c.Locals("userId", userId)

	return nil
}
