package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

func TestAddImage(t *testing.T) {
	img := &model.Image{
		ImageID:   "safsafadfasfsfaf24424234234",
		ImageName: "safsafadfasfsfasf4424234234.jpg",
		ImagePath: "/runtime/upload/images/safsafadfafsfasf24424234234.jpg",
	}
	err := AddImage(img)
	if err != nil {
		t.Error("err:", err)
	}
}
func TestGetImage(t *testing.T) {
	imgs, err := GetImage(0, 2)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Printf("imgs:%v\n", imgs)
}
