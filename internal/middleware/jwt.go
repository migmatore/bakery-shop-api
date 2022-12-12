package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey:   []byte("secret"),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

type TokenWithClaims struct {
	Token    string
	Expires  int64
	Id       int
	Customer bool
	StoreId  int
	Admin    bool
}

// GenerateNewAccessToken TODO move from middleware package
// TODO create access levels
func GenerateNewAccessToken(id int, customer bool, storeId int, admin bool) (*TokenWithClaims, error) {
	expires := time.Now().Add(time.Hour * 48).Unix()

	claims := jwt.MapClaims{
		"id":       id,
		"customer": customer,
		"store_id": storeId,
		"admin":    admin,
		"exp":      expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &TokenWithClaims{
		Token:    t,
		Expires:  expires,
		Id:       id,
		Customer: customer,
		StoreId:  storeId,
		Admin:    admin,
	}, nil
}
