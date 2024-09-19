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
	for _, card := range srv.customsCfg.IcCards {
		go srv.sasService.HandleBoxes(srv.log, card.ImpPath)
		go srv.decService.HandleBoxes(srv.log, card.ImpPath)
	}
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
		companyType string, // 公司类型，比如gxhg, gxwl, gxgyl
	) error
	HandleSentBoxFile(filename string, companyType string) error
	HandleFailBoxFile(filename string, companyType string) error
	HandleInBoxFile(filename string, companyType string) error
	// dir handler
	HandleBoxes(log *zap.SugaredLogger, companyType string)
	HandleSentBox(log *zap.SugaredLogger, companyType string) error
	HandleFailBox(log *zap.SugaredLogger, companyType string) error
	HandleInBox(log *zap.SugaredLogger, companyType string) error
}

type CustomsMessage struct {
	CustomsMessageHandler
	customsCfg *internal.CustomsCfg
}

// TODO: 添加重试机制
func (cm CustomsMessage) HandleBoxes(log *zap.SugaredLogger, companyType string) {
	go cm.HandleInBox(log, companyType)
}

func (cm CustomsMessage) HandleSentBox(log *zap.SugaredLogger, companyType string) (err error) {
	filepathHandler := common.NewFilepathHandler(cm.customsCfg.IcCardMap, cm.DirName())
	return handleBox(
		log,
		filepathHandler.GenSentBoxPath(companyType),
		func(filename string) error {
			return cm.HandleSentBoxFile(filename, companyType)
		},
	)
}

func (cm CustomsMessage) HandleFailBox(log *zap.SugaredLogger, companyType string) (err error) {
	filepathHandler := common.NewFilepathHandler(cm.customsCfg.IcCardMap, cm.DirName())
	return handleBox(
		log,
		filepathHandler.GenFailBoxPath(companyType),
		func(filename string) error {
			return cm.HandleFailBoxFile(filename, companyType)
		},
	)
}

func (cm CustomsMessage) HandleInBox(log *zap.SugaredLogger, companyType string) (err error) {
	filepathHandler := common.NewFilepathHandler(cm.customsCfg.IcCardMap, cm.DirName())
	inBoxPath := filepathHandler.GenInBoxPath(companyType)
	// 处理当前InBox文件夹下面所有的文件(fsnotify代码块的逻辑不会处理之前已经存在的文件)
	// 创建存放处理后文件的文件夹
	if err = utils.CreateDirsIfNotExist(filepathHandler.GenHandledInBoxPath(companyType)); err != nil {
		return
	}
	if err = filepath.WalkDir(inBoxPath, func(filename string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if err = cm.HandleInBoxFile(filename, companyType); err != nil {
			log.Errorf("handle in box file error: %+v", err)
		}
		return nil
	}); err != nil {
		return
	}
	return handleBox(
		log,
		inBoxPath,
		func(filename string) error {
			return cm.HandleInBoxFile(filename, companyType)
		},
	)
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
