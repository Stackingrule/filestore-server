package handler

import (
	"filestore-server/meta"
	"filestore-server/util"
	"io"
	"io/ioutil"
	"log"
	_ "net"
	"net/http"
	"os"
	"time"
)

// UploadHandler ： 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			log.Printf("Failed to get data, err: %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-06 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("ile to create file, err: %s\n", err.Error())
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			log.Printf("Failed to save data into file, err: %s\n", err.Error())
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler : 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

// GetFileMetaHandler : 获取文件元信息
func GetFileMetaHandler() {

}

// FileQueryHandler : 查询批量的文件元信息

// DownloadHandler : 文件下载接口

// FileMetaUpdateHandler ： 更新元信息接口(重命名)

// FileDeleteHandler : 删除文件及元信息

// TryFastUploadHandler : 尝试秒传接口

// DownloadURLHandler : 生成文件的下载地址
