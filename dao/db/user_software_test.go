package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

func TestAddSoftware(t *testing.T) {
	sw := &model.Software{
		SoftwareNO:   3,
		SoftwareName: "vscode",
		Img:          "/runtime/upload/images/avatar.jpg",
		UserID:       1,
	}
	err := AddSoftware(sw)
	if err != nil {
		t.Error("err:", err)
	}

}
func TestGetSoftwares(t *testing.T) {
	sw, err := GetSoftwares(1)
	if err != nil {
		t.Error("err:", err)
	}
	for _, v := range sw {
		fmt.Printf("v:%+v\n", v)
	}
}

func TestEditSoftware(t *testing.T) {
	sw := &model.Software{
		SoftwareName: "haha",
		SoftwareID:   20,
		Img:          "/runtion/upload/images/avatar.jpg",
		UserID:       1,
	}
	err := EditSoftware(sw)
	if err != nil {
		t.Error("err:", err)
	}
}
func TestEditSoftwares(t *testing.T) {
	sws := make(model.Softwares, 0)
	sws = append(sws, &model.Software{
		SoftwareID:   19,
		SoftwareName: "PS",
		Img:          "/runtime/upload/images/1.jpg",
		UserID:       1,
	}, &model.Software{
		SoftwareID:   21,
		SoftwareName: "DW",
		SoftwareNO:   20,
		Img:          "runtime/haha",
		UserID:       1,
	})
	EditSoftwares(&sws)
}
func TestDelSoftwares(t *testing.T) {
	ids := []int{10, 11, 12}
	err := DelSoftwares(1, ids)
	if err != nil {
		t.Error("err:", err)
	}
}

func TestExistSoftware(t *testing.T) {
	sw := &model.Software{
		SoftwareID: 13,
		UserID:     1,
	}
	b, err := ExistSoftware(sw)
	if err != nil {
		t.Error("err:", err)
	}
	if !b {
		t.Error("result wrong")
	}
}
