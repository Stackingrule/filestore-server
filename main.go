package main

import (
	"filestore-server/handler"
	"log"
	"net/http"
)

func main() {

	//设置访问的路由
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.UpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	err := http.ListenAndServe(":8089", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
