package apis

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/infra/models/response"
)

func (hs *HttpServer) GenSasXml(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload common.MessageRequestPayload
	if err = render.DefaultDecoder(r, &payload); err != nil {
		response.FailWithErr(w, r, err)
		return
	}
	var id string
	if id, err = hs.customsSasService.GenOutBoxFile(payload.Data, payload.UploadType, payload.DeclareFlag); err != nil {
		response.FailWithErr(w, r, err)
		return
	}
	response.OkWithData(w, r, id)
}
