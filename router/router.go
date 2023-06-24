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
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/p/", views.HTML.Detail) //文章详情页
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost) //用于搜索
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
