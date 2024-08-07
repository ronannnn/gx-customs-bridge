package customs

import (
	"github.com/ronannnn/gx-customs-bridge/internal"
	"go.uber.org/zap"
)

type SasService struct {
	CustomsMessage
	log        *zap.SugaredLogger
	customsCfg *internal.CustomsCfg
}

func ProvideSasService(
	log *zap.SugaredLogger,
	customsCfg *internal.CustomsCfg,
) *SasService {
	srv := &SasService{
		log:        log,
		customsCfg: customsCfg,
	}
	srv.CustomsMessage.CustomsMessageHandler = srv
	return srv
}

func (srv *SasService) DirName() string {
	return "sas"
}

func (srv *SasService) HandleOutBoxFile(filename string) (err error) {
	srv.log.Info("SasService HandleOutBoxFile, %s", filename)
	return
}

func (srv *SasService) HandleSentBoxFile(filename string) (err error) {
	srv.log.Info("SasService HandleSentBoxFile, %s", filename)
	return
}

func (srv *SasService) HandleFailBoxFile(filename string) (err error) {
	srv.log.Info("SasService HandleFailBoxFile, %s", filename)
	return
}

func (srv *SasService) HandleInBoxFile(filename string) (err error) {
	srv.log.Info("SasService HandleInBoxFile, %s", filename)
	return
}
