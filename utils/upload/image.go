package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"sz_resume_202005/utils"
	"sz_resume_202005/utils/file"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
)

// GetImageFullURL get the full access path
func GetImageFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

// GetImageNameBysha256 get image name
func GetImageNameBysha256(image *multipart.FileHeader) (name string, err error) {
	ext := path.Ext(image.Filename)
	f, err := image.Open()
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	name = utils.EncodeSHA256ByByte(b) + ext
	//err = f.Close()
	return
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {

	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	zlog.Debug("image size:", size)
	if err != nil {
		zlog.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize*1024
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
