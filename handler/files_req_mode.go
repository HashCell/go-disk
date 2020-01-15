package handler

type GetFileMetaReq struct {
	FileHash string `form:"file_hash" bind:"required"`
}

type DownloadFileReq struct {
	FileHash string `form:"file_hash" bind:"required"`
}

type UpdateFileMetaReq struct {
	FileHash string `form:"file_hash" bind:"required"`
	Filename string `form:"filename"`
}

type DeleteFileReq struct {
	FileHash string `form:"file_hash" bind:"required"`
}