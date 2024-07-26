package internal

import (
	"github.com/ronannnn/infra"
	"go.uber.org/zap"
)

func ProvideLog() *zap.SugaredLogger {
	return GlLog
}

var GlLog = func() *zap.SugaredLogger {
	log, err := infra.NewLog(&GlCfg.Log)
	if err != nil {
		panic(err)
	}
	return log
}()
