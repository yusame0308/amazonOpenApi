package usecase

import (
	"amazonOpenApi/internal/http/gen"
	"amazonOpenApi/internal/repository"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Amazon struct {
	db *gorm.DB
}

func NewAmazon(db *gorm.DB) *Amazon {
	return &Amazon{
		db: db,
	}
}

func (p *Amazon) AddAmazon(ctx echo.Context) error {
	// リクエストを取得
	amazon := new(gen.Amazon)
	err := ctx.Bind(amazon)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format")
	}
	//（バリデーション）
	// Create
	now := time.Now()
	p.db.Create(&repository.AmazonData{
		Asin:        amazon.Asin,
		ProductName: amazon.ProductName,
		MakerName:   amazon.MakerName,
		Price:       amazon.Price,
		Reason:      amazon.Reason,
		Url:         amazon.Url,
		CreatedAt:   now,
		UpdatedAt:   now,
		Status:      true,
	})
	return ctx.JSON(http.StatusCreated, amazon)
}

func (p *Amazon) FindAmazonById(ctx echo.Context, asin string) error {
	// データを取得
	m := new(repository.AmazonData)
	if tx := p.db.First(m, "asin = ?", asin); tx.Error != nil {
		return sendError(ctx, http.StatusNotFound, tx.Error.Error())
	}
	amazon := &gen.Amazon{
		ProductName: m.ProductName,
		MakerName:   m.MakerName,
		Price:       m.Price,
		Reason:      m.Reason,
		Url:         m.Url,
		Asin:        m.Asin,
	}
	return ctx.JSON(http.StatusOK, amazon)
}

func (p *Amazon) UpdateAmazon(ctx echo.Context, asin string) error {
	// リクエストを取得
	amazon := new(gen.Amazon)
	err := ctx.Bind(amazon)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format")
	}
	// Update
	now := time.Now()
	p.db.Model(repository.AmazonData{}).
		Where("status = ?", true).
		Where("asin = ?", asin).
		Updates(repository.AmazonData{
			ProductName: amazon.ProductName,
			MakerName:   amazon.MakerName,
			Price:       amazon.Price,
			Reason:      amazon.Reason,
			Url:         amazon.Url,
			Asin:        amazon.Asin,
			UpdatedAt:   now,
		})
	return ctx.JSON(http.StatusOK, amazon)
}

func (p *Amazon) PatchAmazon(ctx echo.Context, asin string) error {
	// リクエストを取得
	amazonPatch := new(gen.AmazonPatch)
	err := ctx.Bind(amazonPatch)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format")
	}
	m := new(repository.AmazonData)
	if tx := p.db.Model(m).First(m, "asin = ?", asin); tx.Error != nil {
		return sendError(ctx, http.StatusBadRequest, tx.Error.Error())
	}

	if amazonPatch.MakerName != nil {
		m.MakerName = *amazonPatch.MakerName
	}
	if amazonPatch.ProductName != nil {
		m.ProductName = *amazonPatch.ProductName
	}
	if amazonPatch.Price != nil {
		m.Price = *amazonPatch.Price
	}
	if amazonPatch.Reason != nil {
		m.Reason = *amazonPatch.Reason
	}
	if amazonPatch.Url != nil {
		m.Url = *amazonPatch.Url
	}
	// Update
	m.UpdatedAt = time.Now()
	p.db.Model(m).
		Where("status = ?", true).
		Where("asin = ?", asin).
		Updates(m)
	return ctx.JSON(http.StatusOK, gen.Amazon{
		ProductName: m.ProductName,
		MakerName:   m.MakerName,
		Price:       m.Price,
		Reason:      m.Reason,
		Url:         m.Url,
		Asin:        m.Asin,
	})
}

func (p *Amazon) DeleteAmazon(ctx echo.Context, asin string) error {
	p.db.Model(repository.AmazonData{}).
		Where("asin = ?", asin).
		Update("status", false)
	return ctx.String(http.StatusNoContent, "")
}

func (p *Amazon) UndeleteAmazon(ctx echo.Context, asin string) error {
	p.db.Model(repository.AmazonData{}).
		Where("asin = ?", asin).
		Update("status", true)
	return ctx.String(http.StatusNoContent, "")
}
