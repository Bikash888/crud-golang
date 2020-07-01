package Controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadSingleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	fmt.Println(file.Size / 1024)
	fmt.Println(filepath.Base(file.Filename))
	c.SaveUploadedFile(file, filepath.Base(file.Filename))
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//upload multiple file at once
func UploadMultipleFile(c *gin.Context,) {
	err := c.Request.ParseMultipartForm(20000) // grab the multipart form
 	if err != nil {
 		fmt.Println(err)
 		return
 	}
  form, _ := c.MultipartForm()

	files := form.File["upload"]
	fmt.Println(len(files))
	for _, file := range files {
		c.SaveUploadedFile(file,filepath.Base(file.Filename) )
	}
  c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  

}
