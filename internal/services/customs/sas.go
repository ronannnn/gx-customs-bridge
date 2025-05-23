package customs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/sasmodels"
	"github.com/ronannnn/gx-customs-bridge/pkg/utils"
	"github.com/ronannnn/infra/mq/rabbitmq"
	infraUtils "github.com/ronannnn/infra/utils"
	"go.uber.org/zap"
)

type SasUploadType string

const (
	SasInv101 SasUploadType = "INV101"
	SasSas121 SasUploadType = "SAS121"
	SasIcp101 SasUploadType = "ICP101"
)

type SasReceiptType string

const (
	// inv101 related
	SasInv201 SasReceiptType = "INV201" // 核注清单审批回执
	SasInv202 SasReceiptType = "INV202" // 核注清单自动生成报关单同一编号报文
	SasInv211 SasReceiptType = "INV211" // 核注清单记账回执

	// sas121 related
	SasSas221 SasReceiptType = "SAS221" // 核放单审核回执
	SasSas223 SasReceiptType = "SAS223" // 核放单过卡回执
	SasSas224 SasReceiptType = "SAS224" // 核放单查验处置回执

	// icp101 related
	SasSas251 SasReceiptType = "SAS251" // 两步申报核放单审核回执/回执补发送
	SasSas253 SasReceiptType = "SAS253" // 两步申报核放单过卡回执
)

type SasService struct {
	CustomsMessage
	log                     *zap.SugaredLogger
	filepathHandler         *common.FilepathHandler
	rmqClient               *rabbitmq.Client
	customsCommonXmlService common.CustomsCommonXmlService
	sasXmlService           sas.SasXmlService
}

func ProvideSasService(
	customsCfg *internal.CustomsCfg,
	log *zap.SugaredLogger,
	rmqClient *rabbitmq.Client,
	customsCommonXmlService common.CustomsCommonXmlService,
	sasXmlService sas.SasXmlService,
) *SasService {
	srv := &SasService{
		CustomsMessage: CustomsMessage{
			customsCfg: customsCfg,
		},
		log:                     log,
		rmqClient:               rmqClient,
		customsCommonXmlService: customsCommonXmlService,
		sasXmlService:           sasXmlService,
	}
	srv.CustomsMessage.CustomsMessageHandler = srv
	srv.filepathHandler = common.NewFilepathHandler(customsCfg.IcCardMap, srv.DirName())
	return srv
}

func (srv *SasService) DirName() string {
	return "Sas"
}

func (srv *SasService) GenOutBoxFile(model any, uploadType string, declareFlag string, companyType string) (err error) {
	switch SasUploadType(uploadType) {
	case SasInv101:
		var inv101 sasmodels.Inv101
		if modelMap, ok := model.(map[string]any); ok {
			// 先转成json，再转成struct
			var tmpBytes []byte
			if tmpBytes, err = json.Marshal(modelMap); err != nil {
				return
			}
			if err = json.Unmarshal(tmpBytes, &inv101); err != nil {
				return
			}
		} else if inv101, ok = model.(sasmodels.Inv101); !ok {
			err = commonmodels.ErrParseInv101
			return
		}
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenInv101Xml(inv101, declareFlag, companyType); err != nil {
			return
		}
		// write xml bytes to file
		var sasFilenameParts sas.FilenameParts
		if sasFilenameParts, err = sas.NewSasFilenameParts(sas.UploadTypeInv101, inv101.Head.ImpexpMarkcd, inv101.Head.EtpsInnerInvtNo); err != nil {
			return
		}
		zipFlePath := srv.filepathHandler.GenOutBoxPath(companyType, sasFilenameParts.GenOutBoxFilename("zip"))
		var zipFileBytes []byte
		if zipFileBytes, err = internal.ZipFile(sasFilenameParts.GenOutBoxFilename("xml"), xmlBytes); err != nil {
			return
		}
		if err = os.WriteFile(zipFlePath, zipFileBytes, 0644); err != nil {
			return
		}
	case SasSas121:
		var sas121 sasmodels.Sas121
		if modelMap, ok := model.(map[string]any); ok {
			// 先转成json，再转成struct
			var tmpBytes []byte
			if tmpBytes, err = json.Marshal(modelMap); err != nil {
				return
			}
			if err = json.Unmarshal(tmpBytes, &sas121); err != nil {
				return
			}
		} else if sas121, ok = model.(sasmodels.Sas121); !ok {
			err = commonmodels.ErrParseInv101
			return
		}
		// validate
		if sas121.Head.IoTypecd == nil {
			err = fmt.Errorf("IoTypecd is required")
			return
		}
		if sas121.Head.EtpsPreentNo == nil {
			err = fmt.Errorf("EtpsPreentNo is required")
			return
		}
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenSas121Xml(sas121, declareFlag, companyType); err != nil {
			return
		}
		var sasFilenameParts sas.FilenameParts
		if sasFilenameParts, err = sas.NewSasFilenameParts(sas.UploadTypeSas121, sas121.Head.IoTypecd, sas121.Head.EtpsPreentNo); err != nil {
			return
		}
		// write xml bytes to file
		zipFlePath := srv.filepathHandler.GenOutBoxPath(companyType, sasFilenameParts.GenOutBoxFilename("zip"))
		var zipFileBytes []byte
		if zipFileBytes, err = internal.ZipFile(sasFilenameParts.GenOutBoxFilename("xml"), xmlBytes); err != nil {
			return
		}
		if err = os.WriteFile(zipFlePath, zipFileBytes, 0644); err != nil {
			return
		}
	case SasIcp101:
		var icp101 sasmodels.Icp101
		if modelMap, ok := model.(map[string]any); ok {
			// 先转成json，再转成struct
			var tmpBytes []byte
			if tmpBytes, err = json.Marshal(modelMap); err != nil {
				return
			}
			if err = json.Unmarshal(tmpBytes, &icp101); err != nil {
				return
			}
		} else if icp101, ok = model.(sasmodels.Icp101); !ok {
			err = commonmodels.ErrParseIcp101
			return
		}
		// validate
		if icp101.Head.EtpsPreentNo == nil {
			err = fmt.Errorf("EtpsPreentNo is required")
			return
		}
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenIcp101Xml(icp101, declareFlag, companyType); err != nil {
			return
		}
		var sasFilenameParts sas.FilenameParts
		emptyIoType := "N"
		if sasFilenameParts, err = sas.NewSasFilenameParts(sas.UploadTypeIcp101, &emptyIoType, icp101.Head.EtpsPreentNo); err != nil {
			return
		}
		// write xml bytes to file
		zipFlePath := srv.filepathHandler.GenOutBoxPath(companyType, sasFilenameParts.GenOutBoxFilename("zip"))
		var zipFileBytes []byte
		if zipFileBytes, err = internal.ZipFile(sasFilenameParts.GenOutBoxFilename("xml"), xmlBytes); err != nil {
			return
		}
		if err = os.WriteFile(zipFlePath, zipFileBytes, 0644); err != nil {
			return
		}
	default:
		err = fmt.Errorf("unsupported upload type: %s, only support INV101, SAS121 and ICP101", uploadType)
	}
	return
}

func (srv *SasService) HandleSentBoxFile(filename string, companyType string) (err error) {
	srv.log.Infof("SasService HandleSentBoxFile, %s, %s", filename, companyType)
	return
}

func (srv *SasService) HandleFailBoxFile(filename string, companyType string) (err error) {
	srv.log.Infof("SasService HandleFailBoxFile, %s", filename)
	var sasFilenameParts sas.FilenameParts
	if sasFilenameParts, err = sas.ParseSasFilename(filename); err != nil {
		return
	}
	if sasFilenameParts.RetryTimes >= 3 {
		srv.log.Errorf("retry times >= 3, move %s to FilesCannotUpload", filename)
		today := time.Now().Format("2006-01-02")
		cannotParsePath := srv.filepathHandler.GenHandledCannotParsePath(companyType, today)
		if err = infraUtils.CreateDirsIfNotExist(cannotParsePath); err != nil {
			return
		}
		cannotParseFilename := srv.filepathHandler.GenHandledCannotParsePath(companyType, today, filepath.Base(filename))
		if err = os.Rename(filename, cannotParseFilename); err != nil {
			return
		}
		return
	}
	// move back to OutBox
	sasFilenameParts.RetryTimes++
	outBoxPath := srv.filepathHandler.GenOutBoxPath(companyType, sasFilenameParts.GenOutBoxFilename("zip"))
	if err = os.Rename(filename, outBoxPath); err != nil {
		return
	}
	return
}

func (srv *SasService) HandleInBoxFile(filename string, companyType string) (err error) {
	srv.log.Infof("SasService HandleInBoxFile, %s", filename)
	if strings.HasSuffix(filename, ".tmp") {
		srv.log.Infof("skip tmp file")
		return
	}
	filenameWithoutParentDir := filepath.Base(filename)
	filePath := srv.filepathHandler.GenInBoxPath(companyType, filenameWithoutParentDir)
	if strings.HasPrefix(filenameWithoutParentDir, "Successed_") || strings.HasPrefix(filenameWithoutParentDir, "Failed_") {
		if err = srv.tryToHandleInBoxMessageResponseFile(filenameWithoutParentDir, companyType); err != nil {
			return
		}
	} else {
		// get xml bytes
		var xmlBytes []byte
		if xmlBytes, err = os.ReadFile(filePath); err != nil {
			return
		}

		// 解析MessageType
		var rmh commonmodels.ReceiptMessageHeader
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
		case SasSas251:
			data, err = srv.sasXmlService.ParseSas251Xml(xmlBytes)
		case SasSas253:
			data, err = srv.sasXmlService.ParseSas253Xml(xmlBytes)
		default:
			err = fmt.Errorf("unsupported receipt type: %s", rmh.EnvelopInfo.MessageType)
		}
		if err != nil {
			return
		}

		// convert data to json bytes
		var jsonbytes []byte
		if jsonbytes, err = json.Marshal(commonmodels.ReceiptResult{
			ReceiptType: rmh.EnvelopInfo.MessageType,
			Data:        data,
		}); err != nil {
			return
		}

		// publish message to rmq
		if err = srv.rmqClient.Push(jsonbytes); err != nil {
			return
		}
	}

	// 把这个文件移动到HandledFilesDirName下
	today := time.Now().Format("2006-01-02")
	handledFilesParentDirPath := srv.filepathHandler.GenHandledInBoxPath(companyType, today)
	if err = infraUtils.CreateDirsIfNotExist(handledFilesParentDirPath); err != nil {
		return
	}
	handledFilesPath := srv.filepathHandler.GenHandledInBoxPath(
		companyType,
		today,
		fmt.Sprintf("handled_%s_%s", time.Now().Format("2006-01-02-15-04-05"), filenameWithoutParentDir),
	)
	if err = os.Rename(filePath, handledFilesPath); err != nil {
		return
	}

	return
}

// 解析InBox里面Successed_或Failed_开头的文件
func (srv *SasService) tryToHandleInBoxMessageResponseFile(filename string, companyType string) (err error) {
	var sasFilenameParts sas.FilenameParts
	if sasFilenameParts, err = sas.ParseSasFilename(filename); err != nil {
		return
	}

	// get xml bytes
	filePath := srv.filepathHandler.GenInBoxPath(companyType, filepath.Base(filename))
	var xmlBytes []byte
	if xmlBytes, err = os.ReadFile(filePath); err != nil {
		return
	}

	// parse xml bytes
	var crm commonmodels.CommonResponeMessage
	if crm, err = srv.customsCommonXmlService.ParseCommonResponseMessageXml(xmlBytes); err != nil {
		var renameErr error
		// 如果解析失败，就把文件移动到FailedFilesDirName下
		today := time.Now().Format("2006-01-02")
		failedFilesParentDirPath := srv.filepathHandler.GenHandledCannotParsePath(companyType, today)
		if renameErr = infraUtils.CreateDirsIfNotExist(failedFilesParentDirPath); renameErr != nil {
			return renameErr
		}
		failedFilesPath := srv.filepathHandler.GenHandledCannotParsePath(
			companyType,
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
	mrr := commonmodels.MessageResponseResult{
		ImpexpMarkcd:         string(sasFilenameParts.ImpexpMarkcd),
		UploadType:           string(sasFilenameParts.UploadType),
		BizTime:              utils.ParseClientGeneratedTime(sasFilenameParts.Timestamp),
		CommonResponeMessage: crm,
	}
	// convert data to json bytes
	var jsonbytes []byte
	if jsonbytes, err = json.Marshal(commonmodels.ReceiptResult{
		ReceiptType: string(sasFilenameParts.UploadType),
		Data:        mrr,
	}); err != nil {
		return
	}

	// publish message to rmq
	if err = srv.rmqClient.Push(jsonbytes); err != nil {
		return
	}

	return
}
