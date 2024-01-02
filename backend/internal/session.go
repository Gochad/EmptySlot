package internal

import (
	"context"
	"gorm.io/gorm"
)

func Database(ctx context.Context) *gorm.DB {
	v, _ := ctx.Value("DB").(*gorm.DB)
	return v
}
