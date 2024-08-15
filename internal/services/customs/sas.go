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
	"github.com/ronannnn/infra"
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
	rmqClient               *infra.RabbitmqClient
	customsCommonXmlService common.CustomsCommonXmlService
	sasXmlService           sas.SasXmlService
}

func ProvideSasService(
	customsCfg *internal.CustomsCfg,
	log *zap.SugaredLogger,
	rmqClient *infra.RabbitmqClient,
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

func (srv *SasService) GenOutBoxFile(model any, uploadType string, declareFlag string) (err error) {
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
		if xmlBytes, err = srv.sasXmlService.GenInv101Xml(inv101, declareFlag); err != nil {
			return
		}
		// write xml bytes to file
		filename := fmt.Sprintf("INV101_%s.xml", *inv101.Head.ImpexpMarkcd)
		zipFlePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), OutBoxDirName, fmt.Sprintf("INV101_%s.zip", *inv101.Head.ImpexpMarkcd))
		var zipFileBytes []byte
		if zipFileBytes, err = internal.ZipFile(filename, xmlBytes); err != nil {
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
		// generate xml bytes
		var xmlBytes []byte
		if xmlBytes, err = srv.sasXmlService.GenSas121Xml(sas121, declareFlag); err != nil {
			return
		}
		// write xml bytes to file
		filename := fmt.Sprintf("SAS121_%s.xml", *sas121.Head.IoTypecd)
		zipFlePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), OutBoxDirName, fmt.Sprintf("SAS121_%s.zip", *sas121.Head.IoTypecd))
		var zipFileBytes []byte
		if zipFileBytes, err = internal.ZipFile(filename, xmlBytes); err != nil {
			return
		}
		if err = os.WriteFile(zipFlePath, zipFileBytes, 0644); err != nil {
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
	if strings.HasSuffix(filename, ".tmp") {
		srv.log.Infof("skip tmp file")
		return
	}
	filenameWithoutParentDir := internal.GetFilenameFromPath(filename)
	filePath := filepath.Join(
		srv.customsCfg.ImpPath,
		srv.DirName(),
		InBoxDirName,
		filenameWithoutParentDir,
	)
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
	handledFilesPath := filepath.Join(
		srv.customsCfg.ImpPath,
		srv.DirName(),
		HandledFilesDirName,
		InBoxDirName,
		time.Now().Format("2006-01-02"),
		fmt.Sprintf("handled_%s_%s", time.Now().Format("2006-01-02-15-04-05"), filenameWithoutParentDir),
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
	splitFilenamePrefixStrList := strings.Split(filenamePrefix, "_")
	// 1. Successed/Failed
	// 2. INV101/SAS121
	// 3. impexpMarkcd
	// 4. id
	// 5. 海关客户端打上的时间戳
	if len(splitFilenamePrefixStrList) != 4 {
		err = fmt.Errorf("filename prefix is invalid: %s", filenamePrefix)
		return
	}
	uploadType := splitFilenamePrefixStrList[1]
	impexpMarkcd := splitFilenamePrefixStrList[2]

	// get xml bytes
	filePath := filepath.Join(srv.customsCfg.ImpPath, srv.DirName(), InBoxDirName, filename)
	var xmlBytes []byte
	if xmlBytes, err = os.ReadFile(filePath); err != nil {
		return
	}

	// parse xml bytes
	var crm commonmodels.CommonResponeMessage
	if crm, err = srv.customsCommonXmlService.ParseCommonResponseMessageXml(xmlBytes); err != nil {
		return
	}

	// convert data to json bytes
	mrr := commonmodels.MessageResponseResult{
		ImpexpMarkcd:         impexpMarkcd,
		UploadType:           uploadType,
		CommonResponeMessage: crm,
	}
	// convert data to json bytes
	var jsonbytes []byte
	if jsonbytes, err = json.Marshal(commonmodels.ReceiptResult{
		ReceiptType: uploadType,
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
