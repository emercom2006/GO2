package test

import (
	"blog/models"
	"fmt"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func TestOne(t *testing.T) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:8889)/blog_bd")
	beego.Run()

	beeOrm := orm.NewOrm()

	posts := []models.Posts{}

	_, err := beeOrm.QueryTable("tablename").All(&posts)
	if err != nil {
		fmt.Println("Ошибка")
	}

}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:8889)/blog_bd")
		beego.Run()

		beeOrm := orm.NewOrm()

		posts := []models.Posts{}

		_, err := beeOrm.QueryTable("tablename").All(&posts)
		if err != nil {
			fmt.Println("Ошибка")
		}

	}
}
