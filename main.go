package main

import (
	"filestore-server/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("file/download", handler.DownloadHandler)

	var addr = ":8089"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Failed to start server, err : %s\n", err.Error())
	} else {
		log.Printf("Start serve %s!\n", addr)
	}

}
