package internal

import (
	"context"

	"gorm.io/gorm"
)

func Database(ctx context.Context) *gorm.DB {
	v, _ := ctx.Value(DbKey).(*gorm.DB)
	return v
}
