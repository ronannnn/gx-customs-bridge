package customs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/ronannnn/gx-customs-bridge/internal/services/rmq"
	"go.uber.org/zap"
)

type SasUploadType string

const (
	SasInv101 SasUploadType = "INV101"
	SasSas121 SasUploadType = "SAS121"
)

type SasReceiptType string

const (
	SasInv201 SasReceiptType = "INV201"
	SasInv202 SasReceiptType = "INV202"
	SasInv211 SasReceiptType = "INV211"

	SasSas221 SasReceiptType = "SAS221"
	SasSas223 SasReceiptType = "SAS223"
	SasSas224 SasReceiptType = "SAS224"
)

type SasService struct {
	CustomsMessage
	customsCfg              *internal.CustomsCfg
	log                     *zap.SugaredLogger
	rmqClient               *rmq.Client
	customsCommonXmlService common.CustomsCommonXmlService
	sasXmlService           sas.SasXmlService
}

func ProvideSasService(
	customsCfg *internal.CustomsCfg,
	log *zap.SugaredLogger,
	rmqClient *rmq.Client,
	customsCommonXmlService common.CustomsCommonXmlService,
	sasXmlService sas.SasXmlService,
) *SasService {
	srv := &SasService{
		customsCfg:              customsCfg,
		log:                     log,
		rmqClient:               rmqClient,
		customsCommonXmlService: customsCommonXmlService,
		sasXmlService:           sasXmlService,
	}
	srv.CustomsMessage.CustomsMessageHandler = srv
	return srv
}

func (srv *SasService) DirName() string {
	return "sas"
}

func (srv *SasService) GenOutBoxFile(model any, uploadType string, declareFlag string) (id string, err error) {
	id = uuid.New().String()
	switch SasUploadType(uploadType) {
	case SasInv101:
		// cast model to sas.Inv101
		inv101, ok := model.(sas.Inv101)
		if !ok {
			err = common.ErrParseInv101
			return
		}
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenInv101Xml(inv101, declareFlag); err != nil {
			return
		}
		// write xml bytes to file
		filePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), fmt.Sprintf("INV101_%s.xml", id))
		if err = os.WriteFile(filePath, xmlBytes, 0644); err != nil {
			return
		}
	case SasSas121:
		// cast model to sas.Sas121
		sas121, ok := model.(sas.Sas121)
		if !ok {
			err = common.ErrParseSas121
			return
		}
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenSas121Xml(sas121, declareFlag); err != nil {
			return
		}
		// write xml bytes to file
		filePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), fmt.Sprintf("SAS121_%s.xml", id))
		if err = os.WriteFile(filePath, xmlBytes, 0644); err != nil {
			return
		}
	default:
		err = fmt.Errorf("unsupported upload type: %s, only support INV101 and SAS121", uploadType)
	}
	return
}

func (srv *SasService) HandleSentBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleSentBoxFile, %s", filename)
	return
}

func (srv *SasService) HandleFailBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleFailBoxFile, %s", filename)
	return
}

func (srv *SasService) HandleInBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleInBoxFile, %s", filename)
	if strings.HasPrefix(filename, "Successed_") || strings.HasPrefix(filename, "Failed_") {
		err = srv.tryToHandleInBoxMessageResponseFile(filename)
		return
	}

	filenameWithoutParentDir := internal.GetFilenameFromPath(filename)

	// get xml bytes
	filePath := filepath.Join(
		srv.customsCfg.ImpPath,
		srv.DirName(),
		InBoxDirName,
		filenameWithoutParentDir,
	)
	var xmlBytes []byte
	if xmlBytes, err = os.ReadFile(filePath); err != nil {
		return
	}

	// 解析MessageType
	var rmh common.ReceiptMessageHeader
	if rmh, err = srv.customsCommonXmlService.ParseReceiptMessageHeader(xmlBytes); err != nil {
		return
	} else if rmh.EnvelopInfo.MessageType == "" {
		err = fmt.Errorf("该报文没有MessageType: %s", filename)
		return
	}

	var data any
	switch SasReceiptType(rmh.EnvelopInfo.MessageType) {
	case SasInv201:
		data, err = srv.sasXmlService.ParseInv201Xml(xmlBytes)
	case SasInv202:
		data, err = srv.sasXmlService.ParseInv202Xml(xmlBytes)
	case SasInv211:
		data, err = srv.sasXmlService.ParseInv211Xml(xmlBytes)
	case SasSas221:
		data, err = srv.sasXmlService.ParseSas221Xml(xmlBytes)
	case SasSas223:
		data, err = srv.sasXmlService.ParseSas223Xml(xmlBytes)
	case SasSas224:
		data, err = srv.sasXmlService.ParseSas224Xml(xmlBytes)
	default:
		err = fmt.Errorf("unsupported receipt type: %s", rmh.EnvelopInfo.MessageType)
	}
	if err != nil {
		return
	}

	// convert data to json bytes
	var jsonbytes []byte
	if jsonbytes, err = json.Marshal(common.ReceiptResult{
		ReceiptType: string(rmh.EnvelopInfo.MessageType),
		Data:        data,
	}); err != nil {
		return
	}

	// publish message to rmq
	if err = srv.rmqClient.Push(jsonbytes); err != nil {
		return
	}

	// 把这个文件移动到HandledFilesDirName下
	handledFilesPath := filepath.Join(
		srv.customsCfg.ImpPath,
		srv.DirName(),
		HandledFilesDirName,
		InBoxDirName,
		fmt.Sprintf("handled_%s_%s", time.Now().Format("2006-01-02-15:04:05"), filenameWithoutParentDir),
	)
	if err = os.Rename(filePath, handledFilesPath); err != nil {
		return
	}

	return
}

// 解析InBox里面Successed_或Failed_开头的文件
func (srv *SasService) tryToHandleInBoxMessageResponseFile(filename string) (err error) {
	// get id from filename
	filenamePrefix := internal.GetFilenamePrefix(filename)
	id := strings.Split(filenamePrefix, "_")[1]

	// get xml bytes
	filePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), filename)
	var xmlBytes []byte
	if xmlBytes, err = os.ReadFile(filePath); err != nil {
		return
	}

	// parse xml bytes
	var crm common.CommonResponeMessage
	if crm, err = srv.customsCommonXmlService.ParseCommonResponseMessageXml(xmlBytes); err != nil {
		return
	}

	// convert rmq message to json bytes
	mrr := common.MessageResponseResult{
		Id:                   id,
		CommonResponeMessage: crm,
	}
	var mrrJsonBytes []byte
	if mrrJsonBytes, err = json.Marshal(mrr); err != nil {
		return
	}

	// publish message to rmq
	if err = srv.rmqClient.Push(mrrJsonBytes); err != nil {
		return
	}

	return
}
