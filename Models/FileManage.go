package Models

import (
	"mime/multipart"
)

func UploadFile(file *multipart.FileHeader, filepath string) (fileId int, err error) {
	var adapter *Adapter
	adapter = adapter.GetAdapterIntance()

	tsql := "INSERT INTO [dbo].[fileupload] (filename,filepath,createtime) OUTPUT INSERTED.id VALUES ('" + file.Filename + "','" + filepath + "',GETDATE())"
	result := globalAdapterInstance.Raw(tsql)
	result.Save(&fileId)
	return fileId, result.Error
}

func GetUploadLists(fileupload *[]Fileupload) error {
	if err := globalAdapterInstance.Find(fileupload).Error; err != nil {
		return err
	}
	return nil
}

func SearchById(FileId string) (res Fileupload, err error) {
	err = globalAdapterInstance.Where("id = ?", FileId).First(&res).Error
	return res, err
}
