package main

import (
	"fmt"
	"goweb/config"
	"goweb/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
	home := path + "/template/template/home.html"
	header := path + "/template/template/layout/header.html"
	footer := path + "/template/template/layout/footer.html"
	personal := path + "/template/template/layout/personal.html"
	post := path + "/template/template/layout/post-list.html"
	pagination := path + "/template/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		fmt.Println("出错：", err)
	}
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "叶培宇",
			ViewCount:    123,
			CreateAt:     "2023-6-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}

func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
