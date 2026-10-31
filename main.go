package main

import (
	"log"
	"net/http"
)

func main() {
	// 把本地 ./static 目录作为静态资源根目录
	// http.Dir("./static") 表示目录，FileServer 会自动处理 index.html 等
	fs := http.FileServer(http.Dir("."))

	// 所有路径都交给文件服务器处理
	http.Handle("/", fs)

	log.Println("Server starting on :8080, open http://0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
