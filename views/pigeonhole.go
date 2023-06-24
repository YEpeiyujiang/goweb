package views

import (
	"goweb/common"
	"goweb/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPosPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
