package apis

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"
	"github.com/ronannnn/infra/models/response"
)

func (hs *HttpServer) GenSasXml(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload commonmodels.MessageRequestPayload
	if err = render.DefaultDecoder(r, &payload); err != nil {
		response.FailWithErr(w, r, err)
		return
	}
	if err = hs.customsSasService.GenOutBoxFile(payload.Data, payload.UploadType, payload.DeclareFlag); err != nil {
		response.FailWithErr(w, r, err)
		return
	}
	response.Ok(w, r)
}
