package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed drawio/src/main/webapp/**
var embeddedFiles embed.FS

func main() {
	// 创建子文件系统，移除路径前缀
	webappFS, err := fs.Sub(embeddedFiles, "drawio/src/main/webapp")
	if err != nil {
		log.Fatal(err)
	}

	// 创建文件服务器
	fs := http.FS(webappFS)
	fileServer := http.FileServer(fs)

	// 设置路由
	http.Handle("/", fileServer)

	// 启动服务器
	log.Println("Serving at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
