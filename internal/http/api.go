package api

import (
	"amazonOpenApi/internal/http/gen"
	"amazonOpenApi/internal/http/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	amazon *usecase.Amazon
}

func NewApi(db *gorm.DB) *Api {
	return &Api{amazon: usecase.NewAmazon(db)}
}

var _ gen.ServerInterface = (*Api)(nil)

func (p *Api) AddAmazon(ctx echo.Context) error {
	return p.amazon.AddAmazon(ctx)
}

func (p *Api) FindAmazonById(ctx echo.Context, asin string) error {
	return p.amazon.FindAmazonById(ctx, asin)
}

func (p *Api) UpdateAmazon(ctx echo.Context, asin string) error {
	return p.amazon.UpdateAmazon(ctx, asin)
}

func (p *Api) PatchAmazon(ctx echo.Context, asin string) error {
	return p.amazon.PatchAmazon(ctx, asin)
}

func (p *Api) DeleteAmazon(ctx echo.Context, asin string) error {
	return p.amazon.DeleteAmazon(ctx, asin)
}

func (p *Api) UndeleteAmazon(ctx echo.Context, asin string) error {
	return p.amazon.UndeleteAmazon(ctx, asin)
}
