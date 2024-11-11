package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
	"pixel-haven/interval/config"
	"pixel-haven/interval/event"
)

func Upload(group *gin.RouterGroup) {
	group.POST("/", func(c *gin.Context) {
		// 获取文件
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Failed to get file: %s", err)
			errors := []ErrorDetail{
				{
					Code:    ErrorCodeMissingField,
					Message: "An unexpected error occurred. Please try again later.",
				},
			}
			ErrorResponseHandler(c, http.StatusInternalServerError, "Internal Server Error", errors)
			c.Abort()
			return
		}
		defer file.Close()

		// 上传到 MinIO
		objectName := uuid.New().String()
		_, err = config.Conf.MinioClient().PutObject(c, config.BuketName, objectName, file, -1, minio.PutObjectOptions{})
		if err != nil {
			log.Printf("Failed to upload file to MinIO: %s", err)
			errors := []ErrorDetail{
				{
					Code:    ErrorCodeMissingField,
					Message: "An unexpected error occurred. Please try again later.",
				},
			}
			ErrorResponseHandler(c, http.StatusInternalServerError, "Internal Server Error", errors)
			c.Abort()
			return
		}
		event.Producer("upload", "file", objectName)

		SuccessResponseHandler(c, http.StatusCreated, objectName)

	})
}
