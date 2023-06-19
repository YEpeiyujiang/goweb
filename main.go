package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Base-Type", "application/json")
	var indexData IndexData
	indexData.Title = "叶培宇的个人博客"
	indexData.Desc = "两天完成"
	jsonStrc, _ := json.Marshal(indexData)
	w.Write(jsonStrc)
}

func main() {
	//程序入口 一个项目 只能有一个入口

	//web程序 http协议 ip port

	server := http.Server{ //绑定服务端口信息
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
