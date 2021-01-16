package meta

// FileMeta: 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetMeta: 通过sha1值获得文件元信息
func GetMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
