package main

import (
	"filestore-server/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/file/upload", handler.UploadHandler)
	var addr = ":8089"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("Failed to start server, err : %s", err.Error())
	} else {
		log.Println("start serve %s!", addr)
	}

}
