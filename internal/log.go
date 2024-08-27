package internal

import (
	"github.com/ronannnn/infra/log"
	"go.uber.org/zap"
)

func ProvideLog(logCfg *log.Cfg) (*zap.SugaredLogger, error) {
	return log.New(logCfg)
}
