package Models

import "Go_OOP/Technical_Service/Entity/EntityStruct"

type Fileupload struct {
	*EntityStruct.Fileupload
}

type ResponseUploadFile struct {
	Message  string `json:"msg"`
	Code     string `json:"code"`
	Filepath string `json:"filepath"`
	Fileid   int    `json:"fileid"`
}

func (b *Fileupload) TableName() string {
	return "fileupload"
}
