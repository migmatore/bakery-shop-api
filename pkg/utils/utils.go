package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

func ParsePgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.Is(err, pgErr) {
		pgErr = err.(*pgconn.PgError)
		return fmt.Errorf("database error. message: %s, details: %s, where: %s, sqlstate: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState())
	}

	return err
}

func FiberError(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})
}
