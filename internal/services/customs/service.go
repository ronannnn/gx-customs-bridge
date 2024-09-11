package customs

import (
	"io/fs"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/infra/utils"
	"go.uber.org/zap"
)

type CustomsService interface {
	ListenImpPath()
}

func ProvideCustomsService(
	log *zap.SugaredLogger,
	customsCfg *internal.CustomsCfg,
	sasService *SasService,
	decService *DecService,
) CustomsService {
	return &CustomsServiceImpl{
		log:        log,
		customsCfg: customsCfg,
		sasService: sasService,
		decService: decService,
	}
}

type CustomsServiceImpl struct {
	log        *zap.SugaredLogger
	customsCfg *internal.CustomsCfg
	sasService *SasService
	decService *DecService
}

func (srv *CustomsServiceImpl) ListenImpPath() {
	go srv.sasService.HandleBoxes(srv.log, srv.customsCfg.ImpPath)
	go srv.decService.HandleBoxes(srv.log, srv.customsCfg.ImpPath)
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
	) error
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
	go cm.HandleInBox(log, impPath)
}

func (cm CustomsMessage) HandleSentBox(log *zap.SugaredLogger, impPath string) (err error) {
	filepathHandler := common.FilepathHandler{ImpPath: impPath, BizType: cm.DirName()}
	return handleBox(log, filepathHandler.GenSentBoxPath(), cm.HandleSentBoxFile)
}

func (cm CustomsMessage) HandleFailBox(log *zap.SugaredLogger, impPath string) (err error) {
	filepathHandler := common.FilepathHandler{ImpPath: impPath, BizType: cm.DirName()}
	return handleBox(log, filepathHandler.GenFailBoxPath(), cm.HandleFailBoxFile)
}

func (cm CustomsMessage) HandleInBox(log *zap.SugaredLogger, impPath string) (err error) {
	filepathHandler := common.FilepathHandler{ImpPath: impPath, BizType: cm.DirName()}
	inBoxPath := filepathHandler.GenInBoxPath()
	// 处理当前InBox文件夹下面所有的文件(fsnotify代码块的逻辑不会处理之前已经存在的文件)
	// 创建存放处理后文件的文件夹
	if err = utils.CreateDirsIfNotExist(filepathHandler.GenHandledInBoxPath()); err != nil {
		return
	}
	if err = filepath.WalkDir(inBoxPath, func(filename string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if err = cm.HandleInBoxFile(filename); err != nil {
			log.Errorf("handle in box file error: %+v", err)
		}
		return nil
	}); err != nil {
		return
	}
	return handleBox(log, inBoxPath, cm.HandleInBoxFile)
}

func handleBox(
	log *zap.SugaredLogger,
	path string,
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

		// 监听文件变化
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					if err = handleBoxFileFn(event.Name); err != nil {
						log.Errorf("handle file error: %+v", err)
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
