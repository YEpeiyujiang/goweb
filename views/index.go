package views

import (
	"errors"
	"goweb/common"
	"goweb/service"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义

	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错:", err)
		index.WriteError(w, errors.New("系统错误 请联系管理员"))
	}

	index.WriteData(w, hr)
}
