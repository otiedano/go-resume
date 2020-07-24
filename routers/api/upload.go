package api

import (
	"fmt"
	"net/http"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"sz_resume_202005/utils/upload"

	"github.com/gin-gonic/gin"
)

//UploadImage 上传图片
// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	fmt.Printf("UploadImage worked\n")
	g := g.G(c)
	file, image, err := c.Request.FormFile("file")

	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	if image == nil {
		zlog.Warn("没读取到图片")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName, err := upload.GetImageNameBysha256(image)

	if err != nil {
		zlog.Error("err:", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		zlog.Error("err:图片格式或尺寸错误")
		g.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	err = service.AddImage(savePath + imageName)
	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullURL(imageName),
		"image_save_url": savePath + imageName,
	})
}

//UploadImg 富文本上传图片
func UploadImg(c *gin.Context) {
	fmt.Printf("UploadImage worked\n")
	g := g.G(c)
	file, image, err := c.Request.FormFile("file")

	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	if image == nil {
		zlog.Warn("没读取到图片")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName, err := upload.GetImageNameBysha256(image)

	if err != nil {
		zlog.Error("err:", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		zlog.Error("err:图片格式或尺寸错误")
		g.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	err = service.AddImage(savePath + imageName)
	if err != nil {
		zlog.Warn(err)
		g.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		//"location": savePath + imageName,
		"location": upload.GetImageFullURL(imageName),
	})

}
