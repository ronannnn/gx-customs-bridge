package customs

import (
	"io/fs"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/infra/utils"
	"go.uber.org/zap"
)

const HandledFilesDirName = "Gx"
const OutBoxDirName = "OutBox"
const SentBoxDirName = "SentBox"
const FailBoxDirName = "FailBox"
const InBoxDirName = "InBox"

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
	HandleSentBoxFile(filename string) error
	HandleFailBoxFile(filename string) error
	HandleInBoxFile(filename string) error
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
	return handleBox(log, impPath, cm.DirName(), SentBoxDirName, cm.HandleSentBoxFile)
}

func (cm CustomsMessage) HandleFailBox(log *zap.SugaredLogger, impPath string) (err error) {
	return handleBox(log, impPath, cm.DirName(), FailBoxDirName, cm.HandleFailBoxFile)
}

func (cm CustomsMessage) HandleInBox(log *zap.SugaredLogger, impPath string) (err error) {
	return handleBox(log, impPath, cm.DirName(), InBoxDirName, cm.HandleInBoxFile)
}

func handleBox(
	log *zap.SugaredLogger,
	impPath string,
	dirName string,
	boxName string,
	handleBoxFileFn func(filename string) error,
) (err error) {
	path := filepath.Join(impPath, dirName, boxName)
	// out box的额外逻辑
	if boxName == InBoxDirName {
		// 创建存放处理后文件的文件夹
		handledFilesPath := filepath.Join(impPath, dirName, HandledFilesDirName, boxName)
		if err = utils.CreateDirsIfNotExist(handledFilesPath); err != nil {
			return
		}
		// 处理当前OutBox文件夹下面所有的文件(fsnotify代码块的逻辑不会处理之前已经存在的文件)
		if err = filepath.WalkDir(path, func(filename string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if err = handleBoxFileFn(filename); err != nil {
				log.Errorf("handle outbox file error: %+v", err)
			}
			return nil
		}); err != nil {
			return
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		// 监听文件变化
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

	if err = watcher.Add(path); err != nil {
		return
	}
	<-done
	return
}
