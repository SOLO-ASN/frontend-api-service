package handler

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/response"
	"api-service/internal/retriever"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IImagesHandler interface {
	Create(c *gin.Context)
	Upload(c *gin.Context)
}

type imagesHandler struct {
	retriever retriever.ImagesRetriever
}

func NewImagesHandler() IImagesHandler {
	return &imagesHandler{
		retriever: retriever.NewImagesRetriever(
			model.GetDb(false),
			&cache.Cache{}),
	}
}

func (s imagesHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s imagesHandler) Upload(c *gin.Context) {
	//TODO implement me
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		log.Fatal(err)
	}
	// 获取表单
	form := c.Request.MultipartForm
	// 获取参数upload后面的多个文件名，存放到数组files里面
	files := form.File["file"]
	// 遍历数组，每取出一个file就拷贝一次
	for _, file := range files {
		fileHandle, err := file.Open()
		defer fileHandle.Close()
		if err != nil {
			log.Fatal(err)
		}
		u := uuid.New()
		// 为文件生成一个新的文件名，以避免冲突
		newFileName := fmt.Sprintf("%s_%s", u, file.Filename) // 这里可以根据需要生成不同的文件名格式

		// 指定文件保存的路径
		filePath := filepath.Join("/home/l6-809/go/src/github.com/Solo-Mission/uploadImages", newFileName)

		out, err := os.Create(filePath)
		defer out.Close()
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, fileHandle)
		if err != nil {
			log.Fatal(err)
		}
		backfilePath := "https://airdrop.aspark.space/static/" + newFileName
		response.OutPut(c, response.WithCodeMessage{
			Code:    62001,
			Message: "SUCCESSED",
		}, backfilePath)
	}

}
