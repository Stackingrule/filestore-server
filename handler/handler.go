package handler

import (
	"io"
	"io/ioutil"
	_ "net"
	"net/http"
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

	}
}

// UploadSucHandler : 上传已完成

// GetFileMetaHandler : 获取文件元信息

// FileQueryHandler : 查询批量的文件元信息

// DownloadHandler : 文件下载接口

// FileMetaUpdateHandler ： 更新元信息接口(重命名)

// FileDeleteHandler : 删除文件及元信息

// TryFastUploadHandler : 尝试秒传接口

// DownloadURLHandler : 生成文件的下载地址
