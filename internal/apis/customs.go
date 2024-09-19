package apis

import (
	"net/http"

	"github.com/ronannnn/gx-customs-bridge/internal/base/reason"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
	"github.com/ronannnn/infra/msg"
)

func (hs *HttpServer) GenSasXml(w http.ResponseWriter, r *http.Request) {
	var payload commonmodels.MessageRequestPayload
	if hs.h.BindAndCheck(w, r, &payload) {
		return
	}
	if err := hs.customsSasService.GenOutBoxFile(payload.Data, payload.UploadType, payload.DeclareFlag, string(payload.CompanyType)); err != nil {
		hs.h.Fail(w, r, err, nil)
	} else {
		hs.h.Success(w, r, msg.New(reason.SuccessToGenSasXml), nil)
	}
}

func (hs *HttpServer) GenDecXml(w http.ResponseWriter, r *http.Request) {
	var payload decmodels.MessageRequestPayload
	if hs.h.BindAndCheck(w, r, &payload) {
		return
	}
	if err := hs.customsDecService.GenOutBoxFile(payload.Data, "", string(payload.OperType), string(payload.CompanyType)); err != nil {
		hs.h.Fail(w, r, err, nil)
	} else {
		hs.h.Success(w, r, msg.New(reason.SuccessToGenSasXml), nil)
	}
}
