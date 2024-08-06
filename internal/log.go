package internal

import (
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"go.uber.org/zap"
)

func ProvideLog(logCfg *cfg.Log) (*zap.SugaredLogger, error) {
	return infra.NewLog(logCfg)
}
