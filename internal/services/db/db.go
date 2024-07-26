package db

import (
	"time"

	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var tables = []any{}

func ProvideService(
	cfg *cfg.Db,
	log *zap.SugaredLogger,
) (db *gorm.DB, err error) {
	startTime := time.Now()
	defer infra.PrintModuleLaunchedDuration(log, "db", startTime)
	return infra.NewDb(cfg, false, tables)
}
