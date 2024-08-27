package db

import (
	"time"

	infraDb "github.com/ronannnn/infra/db"
	"github.com/ronannnn/infra/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var tables = []any{}

func ProvideService(
	cfg *infraDb.Cfg,
	logger *zap.SugaredLogger,
) (db *gorm.DB, err error) {
	startTime := time.Now()
	defer log.PrintModuleLaunchedDuration(logger, "db", startTime)
	return infraDb.New(cfg, false, tables)
}
