package rmq

import (
	"time"

	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"go.uber.org/zap"
)

func ProvideService(
	log *zap.SugaredLogger,
	cfg *cfg.Rabbitmq,
) (rmqClient *infra.RabbitmqClient) {
	startTime := time.Now()
	defer infra.PrintModuleLaunchedDuration(log, "rmq", startTime)
	return infra.NewRabbitMq(log, cfg)
}
