package rmq

import (
	"time"

	"github.com/ronannnn/infra/log"
	"github.com/ronannnn/infra/mq/rabbitmq"
	"go.uber.org/zap"
)

func ProvideService(
	logger *zap.SugaredLogger,
	cfg *rabbitmq.Cfg,
) (rmqClient *rabbitmq.Client) {
	startTime := time.Now()
	defer log.PrintModuleLaunchedDuration(logger, "rmq", startTime)
	return rabbitmq.New(logger, cfg)
}
