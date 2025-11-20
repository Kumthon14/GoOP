package Controllers

import (
	"Go_OOP/Models"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileManageController struct{}

func (f *FileManageController) NewFileManageController() {}

// @Summary Get File Detail
// @Description Get File Detail
// @ID GetFileDetail
// @Tags Files
// @Success 200 {object} []Models.Fileupload "Success"
// @Failure 400 {string} string "Error"
// @response 401 {string} string "Unauthorized"
// @Router /upload-api/getUploadLists [GET]
// @security ApiKeyAuth
func GetUploadLists(c *gin.Context) {
	var fileupload []Models.Fileupload
	err := Models.GetUploadLists(&fileupload)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, fileupload)
	}
}

func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	res := []Models.ResponseUploadFile{}
	files := form.File["file"]

	if files == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if len(files) < 1 {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	for _, file := range files {
		if file != nil && file.Size > 0 {
			uploaded := Models.ResponseUploadFile{}
			uploaded.Filepath, uploaded.Fileid, err = SaveFile(file)

			if err != nil {
				res = append(res, Models.ResponseUploadFile{
					Message:  "Upload File [" + file.Filename + "]: [Internal Server Error] " + err.Error(),
					Code:     "500",
					Filepath: "Error",
					Fileid:   -1,
				})
				continue
			}

			res = append(res, Models.ResponseUploadFile{
				Message:  "Upload File [" + file.Filename + "]: Success",
				Code:     "200",
				Filepath: uploaded.Filepath,
				Fileid:   uploaded.Fileid,
			})

		} else {
			res = append(res, Models.ResponseUploadFile{
				Message:  "Upload File [" + file.Filename + "]: [Bad Request Error] File is null or zero",
				Code:     "400",
				Filepath: "Error",
				Fileid:   1,
			})
		}
	}

	c.JSON(http.StatusMultiStatus, res)
}

func SaveFile(file *multipart.FileHeader) (filePath string, fileId int, err error) {
	prefixPath := "./Uploads"

	if _, err := os.Stat(prefixPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(prefixPath, os.ModePerm)
		if err != nil {
			return filePath, fileId, errors.New("Error Process Create File Path, " + err.Error())
		}
		err = nil
	}

	filePath = fmt.Sprintf("%s/%s", prefixPath, file.Filename)

	src, err := file.Open()

	if src != nil {
		defer func(src multipart.File) {
			err := src.Close()
			if err != nil {
				fmt.Println("Src Close Error.")
			}
		}(src)
	}

	if err != nil {
		return filePath, fileId, errors.New("error process open")
	}

	dst, err := os.Create(filePath)
	if dst != nil {
		defer func(dst *os.File) {
			err := dst.Close()
			if err != nil {
				fmt.Println("Dst Close Error")
			}
		}(dst)
	}

	if err != nil {
		return filePath, fileId, errors.New("Error Process Create File, " + err.Error())
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		return filePath, fileId, errors.New("Error Process Copy File, " + err.Error())
	}

	fileId, err = Models.UploadFile(file, filePath)

	if err != nil {
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			_ = os.Remove(filePath)
		}

		return filePath, fileId, errors.New("Error Process Insert Upload File Record, " + err.Error())
	}

	return filePath, fileId, nil
}

func DownloadFileById(c *gin.Context) {
	fmt.Println("NotFound")
	FileId := c.Param("fileId")
	res, err := Models.SearchById(FileId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Not Found")
		return
	}

	c.File(res.Filepath)
}
