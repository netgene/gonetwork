package business

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  logger "github.com/donnie4w/go-logger/logger"
)

var db *gorm.DB
var err error

type MyInfo struct {
  id int
  myname string     
}

func InitDB() {
  //"user:password@(host:port)/database?charset=utf8&parseTime=True&loc=Local"
  db, err = gorm.Open("mysql", "root:0709@(127.0.0.1:3306)/example?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println("failed to connect database")
    return
  }
  db.SingularTable(true)
  //db.LogMode(true)
  logger.Info("InitDB")
}

func CloseDB() {
  db.Close()
  logger.Info("CloseDB")
}

func GetInfo(info *MyInfo) {
  logger.Info("GetInfo")

  err = db.Raw("select id, myname from myinfo limit 1").Scan(info).Error
  if err != nil {
    fmt.Println("select error:")
    return
  }
  logger.Info(fmt.Sprintf("%+v\n", info))
}
