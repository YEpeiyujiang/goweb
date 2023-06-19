package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Base-Type", "application/json")
	var indexData IndexData
	indexData.Title = "叶培宇的个人博客"
	indexData.Desc = "即将崛起的超级大佬"
	jsonStrc, _ := json.Marshal(indexData)
	w.Write(jsonStrc)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "叶培宇的个人博客"
	indexData.Desc = "即将崛起的超级大佬"
	t := template.New("index.html")
	//拿到当前的路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	//程序入口 一个项目 只能有一个入口

	//web程序 http协议 ip port

	server := http.Server{ //绑定服务端口信息
		Addr: "127.0.0.1:8080",
	}
	//http.HandleFunc("/", index)
	//if err := server.ListenAndServe(); err != nil {
	//	log.Println(err)
	//}

	http.HandleFunc("/index", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
