package customs

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/dec"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
	"github.com/ronannnn/infra/mq/rabbitmq"
	"github.com/ronannnn/infra/utils"
	"go.uber.org/zap"
)

type DecService struct {
	CustomsMessage
	customsCfg              *internal.CustomsCfg
	log                     *zap.SugaredLogger
	filepathHandler         *common.FilepathHandler
	rmqClient               *rabbitmq.Client
	customsCommonXmlService common.CustomsCommonXmlService
	decXmlService           dec.DecXmlService
}

func ProvideDecService(
	customsCfg *internal.CustomsCfg,
	log *zap.SugaredLogger,
	rmqClient *rabbitmq.Client,
	customsCommonXmlService common.CustomsCommonXmlService,
	decXmlService dec.DecXmlService,
) *DecService {
	srv := &DecService{
		customsCfg:              customsCfg,
		log:                     log,
		rmqClient:               rmqClient,
		customsCommonXmlService: customsCommonXmlService,
		decXmlService:           decXmlService,
	}
	srv.CustomsMessage.CustomsMessageHandler = srv
	srv.filepathHandler = common.NewFilepathHandler(customsCfg.ImpPath, srv.DirName())
	return srv
}

func (srv *DecService) DirName() string {
	return "Deccus001"
}

func (srv *DecService) GenOutBoxFile(model any, uploadType string, declareFlag string) (err error) {
	var decModel decmodels.Dec
	if modelMap, ok := model.(map[string]any); ok {
		// 先转成json，再转成struct
		var tmpBytes []byte
		if tmpBytes, err = json.Marshal(modelMap); err != nil {
			return
		}
		if err = json.Unmarshal(tmpBytes, &decModel); err != nil {
			return
		}
	} else if decModel, ok = model.(decmodels.Dec); !ok {
		err = commonmodels.ErrParseDecTmp
		return
	}
	// generate xml bytes
	var xmlBytes []byte
	if xmlBytes, err = srv.decXmlService.GenDecTmpXml(decModel, declareFlag); err != nil {
		return
	}
	// write xml bytes to file
	var decFilenameParts dec.FilenameParts
	if decFilenameParts, err = dec.NewDecFilenameParts(decModel.DecHead.IEFlag, decModel.DecSign.ClientSeqNo); err != nil {
		return
	}
	zipFlePath := srv.filepathHandler.GenOutBoxPath(decFilenameParts.GenOutBoxFilename("zip"))
	var zipFileBytes []byte
	if zipFileBytes, err = internal.ZipFile(decFilenameParts.GenOutBoxFilename("xml"), xmlBytes); err != nil {
		return
	}
	if err = os.WriteFile(zipFlePath, zipFileBytes, 0644); err != nil {
		return
	}
	return
}

func (srv *DecService) HandleSentBoxFile(filename string) (err error) {
	srv.log.Infof("DecService HandleSentBoxFile, %s", filename)
	return
}

func (srv *DecService) HandleFailBoxFile(filename string) (err error) {
	srv.log.Infof("DecService HandleFailBoxFile, %s", filename)
	var decFilenameParts dec.FilenameParts
	if decFilenameParts, err = dec.ParseDecFilename(filename); err != nil {
		return
	}
	if decFilenameParts.RetryTimes >= 3 {
		srv.log.Errorf("retry times >= 3, move %s to FilesCannotUpload", filename)
		today := time.Now().Format("2006-01-02")
		cannotParsePath := srv.filepathHandler.GenHandledCannotParsePath(today)
		if err = utils.CreateDirsIfNotExist(cannotParsePath); err != nil {
			return
		}
		cannotParseFilename := srv.filepathHandler.GenHandledCannotParsePath(today, filepath.Base(filename))
		if err = os.Rename(filename, cannotParseFilename); err != nil {
			return
		}
		return
	}
	// move back to OutBox
	decFilenameParts.RetryTimes++
	outBoxPath := srv.filepathHandler.GenOutBoxPath(decFilenameParts.GenOutBoxFilename("zip"))
	if err = os.Rename(filename, outBoxPath); err != nil {
		return
	}
	return
}

func (srv *DecService) HandleInBoxFile(filename string) (err error) {
	srv.log.Infof("DecService HandleInBoxFile, %s", filename)
	if strings.HasSuffix(filename, ".tmp") {
		srv.log.Infof("skip tmp file")
		return
	}
	filenameWithoutParentDir := filepath.Base(filename)
	filePath := srv.filepathHandler.GenInBoxPath(filenameWithoutParentDir)
	if strings.HasPrefix(filenameWithoutParentDir, "Successed_") || strings.HasPrefix(filenameWithoutParentDir, "Failed_") {
		if err = srv.tryToHandleInBoxMessageResponseFile(filenameWithoutParentDir); err != nil {
			return
		}
	} else {
		// get xml bytes
		var xmlBytes []byte
		if xmlBytes, err = os.ReadFile(filePath); err != nil {
			return
		}

		// 解析MessageType
		var dr decmodels.DecResult
		if dr, err = srv.parseDecResultXml(xmlBytes); err != nil {
			return
		}

		// convert data to json bytes
		var jsonbytes []byte
		if jsonbytes, err = json.Marshal(dr); err != nil {
			return
		}

		// publish message to rmq
		if err = srv.rmqClient.Push(jsonbytes); err != nil {
			return
		}
	}

	// 把这个文件移动到HandledFilesDirName下
	today := time.Now().Format("2006-01-02")
	handledFilesParentDirPath := srv.filepathHandler.GenHandledInBoxPath(today)
	if err = utils.CreateDirsIfNotExist(handledFilesParentDirPath); err != nil {
		return
	}
	handledFilesPath := srv.filepathHandler.GenHandledInBoxPath(
		today,
		fmt.Sprintf("handled_%s_%s", time.Now().Format("2006-01-02-15-04-05"), filenameWithoutParentDir),
	)
	if err = os.Rename(filePath, handledFilesPath); err != nil {
		return
	}

	return
}

func (s *DecService) parseDecImportResponseXml(xmlBytes []byte) (crm decmodels.DecImportResponse, err error) {
	err = xml.Unmarshal(xmlBytes, &crm)
	return
}

func (s *DecService) parseDecResultXml(xmlBytes []byte) (crm decmodels.DecResult, err error) {
	err = xml.Unmarshal(xmlBytes, &crm)
	return
}

// 解析InBox里面Successed_或Failed_开头的文件
func (srv *DecService) tryToHandleInBoxMessageResponseFile(filename string) (err error) {
	var decFilenameParts dec.FilenameParts
	if decFilenameParts, err = dec.ParseDecFilename(filename); err != nil {
		return
	}

	// get xml bytes
	filePath := srv.filepathHandler.GenInBoxPath(filepath.Base(filename))
	var xmlBytes []byte
	if xmlBytes, err = os.ReadFile(filePath); err != nil {
		return
	}

	// parse xml bytes
	var decImpResp decmodels.DecImportResponse
	if decImpResp, err = srv.parseDecImportResponseXml(xmlBytes); err != nil {
		var renameErr error
		// 如果解析失败，就把文件移动到FailedFilesDirName下
		today := time.Now().Format("2006-01-02")
		failedFilesParentDirPath := srv.filepathHandler.GenHandledCannotParsePath(today)
		if renameErr = utils.CreateDirsIfNotExist(failedFilesParentDirPath); renameErr != nil {
			return renameErr
		}
		failedFilesPath := srv.filepathHandler.GenHandledCannotParsePath(
			today,
			fmt.Sprintf("cannot_parse_%s", filename),
		)
		if renameErr = os.Rename(filePath, failedFilesPath); renameErr != nil {
			return renameErr
		}
		srv.log.Warnf("parse common response message xml failed, move %s to %s", filename, failedFilesParentDirPath)
		return
	}

	// convert data to json bytes
	mrr := decmodels.MessageResponseResult{
		ImpexpMarkcd:      string(decFilenameParts.ImpexpMarkcd),
		DecImportResponse: decImpResp,
	}
	// convert data to json bytes
	var jsonbytes []byte
	if jsonbytes, err = json.Marshal(mrr); err != nil {
		return
	}

	// publish message to rmq
	if err = srv.rmqClient.Push(jsonbytes); err != nil {
		return
	}

	return
}
