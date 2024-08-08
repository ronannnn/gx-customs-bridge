package customs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
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
	log           *zap.SugaredLogger
	customsCfg    *internal.CustomsCfg
	sasXmlService sas.SasXmlService
}

func ProvideSasService(
	log *zap.SugaredLogger,
	customsCfg *internal.CustomsCfg,
	sasXmlService sas.SasXmlService,
) *SasService {
	srv := &SasService{
		log:           log,
		customsCfg:    customsCfg,
		sasXmlService: sasXmlService,
	}
	srv.CustomsMessage.CustomsMessageHandler = srv
	return srv
}

func (srv *SasService) DirName() string {
	return "sas"
}

func (srv *SasService) GenOutBoxFile(model any, uploadType string, declareFlag string) (id string, err error) {
	id = uuid.New().String()
	switch uploadType {
	case string(SasInv101):
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
	case string(SasSas121):
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

func (srv *SasService) ParseSentBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleSentBoxFile, %s", filename)
	return
}

func (srv *SasService) ParseFailBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleFailBoxFile, %s", filename)
	return
}

func (srv *SasService) ParseInBoxFile(filename string) (err error) {
	srv.log.Infof("SasService HandleInBoxFile, %s", filename)
	return
}
