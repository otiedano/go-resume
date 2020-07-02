package apipublic

import (
	"log"
	"testing"

	"github.com/astaxie/beego/validation"
)

func TestValid(t *testing.T) {
	var u struct {
		Name string
		Age  int
	}

	u.Age = 20
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 18, "age")
	valid.Email(u.Name, "email")
	valid.Required(u.Name, "afas")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Print(err.Key, err.Message, "\n")
		}
	}

}
