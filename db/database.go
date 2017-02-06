package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//gormで依存がある為

	"database/sql"

	"github.com/funayoseyoshito/yakiniku-image-id/lib"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DatabaseSet struct {
	User, Password, Host, Port, Name string
}

func NewDatabaseSet(user string, password string, host string, port string, name string) *DatabaseSet {
	return &DatabaseSet{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name}
}

//GetConnection データベースコネクションを返却します
func (d *DatabaseSet) Connection() *gorm.DB {
	if db == nil {
		databaseInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			d.User,
			d.Password,
			d.Host,
			d.Port,
			d.Name)

		var err error
		db, err = gorm.Open("mysql", databaseInfo)
		if err != nil {
			fmt.Println(err)
			panic("failed to connect database")
		}
	}
	return db
}

//SelectProcessingRows 処理対象のレコードを定数ずつ取得する
func (d *DatabaseSet) SelectProcessingRows(offset int) *sql.Rows {

	rows, err := d.Connection().
		Model(&Images{}).
		Where("kind in (?)", []int{lib.Config.Cooking.OriginID, lib.Config.Other.OriginID}).
		Offset(offset).Limit(SelectLimit).Rows()

	if err != nil {
		fmt.Println(err)
		panic("オリジナル画像の取得に失敗")
	}
	return rows
}
