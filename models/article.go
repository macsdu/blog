package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type articles byte

var _gArticles articles = ARTICLES_MODEL

func ArticlesModel() *articles {
	return &_gArticles
}

func (this *articles) ArticleFromId(id int) ([]orm.Params, error) {
	o := orm.NewOrm()
	sqlStr := "SELECT title,content,a.ctime a_ctime,comment,ip,c.ctime FROM articles a LEFT JOIN comments c ON a.id = c.article_id WHERE a.id = " + strconv.Itoa(id)
	var maps []orm.Params
	_, err := o.Raw(sqlStr).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	beego.Debug(maps)
	if len(maps) == 0 {
		return nil, E_NOT_FOUND
	}
	return maps, nil
}

func (this *articles) Articles() ([]orm.Params, error) {
	o := orm.NewOrm()
	sqlStr := "SELECT * FROM articles"
	var maps []orm.Params
	_, err := o.Raw(sqlStr).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	beego.Debug(maps)
	if len(maps) == 0 {
		return nil, E_NOT_FOUND
	}
	return maps, nil
}

func (this *articles) Add(title, content string) error {
	o := orm.NewOrm()
	_, err := o.Raw("insert articles(title,content, ctime) values(?,?,?)", title, content, time.Now()).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}