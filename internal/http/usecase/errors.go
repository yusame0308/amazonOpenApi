package usecase

import (
	"amazonOpenApi/internal/http/gen"

	"github.com/labstack/echo/v4"
)

func sendError(ctx echo.Context, code int, message string) error {
	sendedErr := gen.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, sendedErr)
	return err
}
