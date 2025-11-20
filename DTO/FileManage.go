package DTO

type UploadFileRequest struct {
	Message  string `json:"msg"`
	Code     string `json:"code"`
	Filepath string `json:"filepath"`
	Fileid   int    `json:"fileid"`
}
