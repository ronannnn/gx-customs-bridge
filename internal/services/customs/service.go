package customs

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/ronannnn/gx-customs-bridge/internal"
	"go.uber.org/zap"
)

type CustomsService interface {
	ListenImpPath()
}

func ProvideCustomsService(
	log *zap.SugaredLogger,
	customsCfg *internal.CustomsCfg,
	sasService *SasService,
) CustomsService {
	return &CustomsServiceImpl{
		log:        log,
		customsCfg: customsCfg,
		sasService: sasService,
	}
}

type CustomsServiceImpl struct {
	log        *zap.SugaredLogger
	customsCfg *internal.CustomsCfg
	sasService *SasService
}

func (srv *CustomsServiceImpl) ListenImpPath() {
	go srv.sasService.HandleBoxes(srv.log, srv.customsCfg.ImpPath)
}

// CustomsMessageFileHandler, 每个海关报文类型都需要实现这个接口
type CustomsMessageHandler interface {
	// meta
	DirName() string
	// file handler
	GenOutBoxFile(
		model any, // 传入的model
		uploadType string, // 报文类型，比如INV101, SAS121
		declareFlag string, // 是否申报
	) (string, error)
	ParseSentBoxFile(filename string) error
	ParseFailBoxFile(filename string) error
	ParseInBoxFile(filename string) error
	// dir handler
	HandleBoxes(log *zap.SugaredLogger, impPath string)
	HandleSentBox(log *zap.SugaredLogger, impPath string) error
	HandleFailBox(log *zap.SugaredLogger, impPath string) error
	HandleInBox(log *zap.SugaredLogger, impPath string) error
}

type CustomsMessage struct {
	CustomsMessageHandler
}

// TODO: 添加重试机制
func (cm CustomsMessage) HandleBoxes(log *zap.SugaredLogger, impPath string) {
	go cm.HandleSentBox(log, impPath)
	go cm.HandleFailBox(log, impPath)
	go cm.HandleInBox(log, impPath)
}

func (cm CustomsMessage) HandleSentBox(log *zap.SugaredLogger, impPath string) (err error) {
	return handleBox(log, impPath, cm.DirName(), "SentBox", cm.ParseSentBoxFile)
}

func (cm CustomsMessage) HandleFailBox(log *zap.SugaredLogger, impPath string) (err error) {
	return handleBox(log, impPath, cm.DirName(), "FailBox", cm.ParseFailBoxFile)
}

func (cm CustomsMessage) HandleInBox(log *zap.SugaredLogger, impPath string) (err error) {
	return handleBox(log, impPath, cm.DirName(), "InBox", cm.ParseInBoxFile)
}

func handleBox(
	log *zap.SugaredLogger,
	impPath string,
	dirName string,
	boxName string,
	handleBoxFileFn func(filename string) error,
) (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					if err = handleBoxFileFn(event.Name); err != nil {
						log.Errorf("handle outbox file error: %+v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Errorf("error: %+v", err)
			}
		}
	}()

	path := filepath.Join(impPath, dirName, boxName)
	if err = watcher.Add(path); err != nil {
		return
	}
	<-done
	return
}
