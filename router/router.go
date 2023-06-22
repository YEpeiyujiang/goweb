package router

import (
	"goweb/api"
	"goweb/views"
	"net/http"
	_ "time"
)

func Router() {
	//1、返回一个页面view 2、返回数据json 3、返回静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
