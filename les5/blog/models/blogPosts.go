package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	Id    uint64
	Title string
	Short string
	Long  string
}

func (t *Posts) TableName() string {
	return "tablename"
}

func NewPost(title, short, long string) (*Posts, error) {
	if title == "" || short == "" || long == "" {
		return nil, fmt.Errorf("Empty")
	}
	return &Posts{Title: title, Short: short, Long: long}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
