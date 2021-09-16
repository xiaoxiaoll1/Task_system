package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"taskSystem/dao"
	"taskSystem/entity"
)

// init 完成数据库链接的初始化
func init()  {
	log.Println("数据表初始化中")
	dsn := "root:123456@tcp(121.4.177.109:3307)/task_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})

	dao.InitDao(db)
	if err != nil {
		log.Println(err)
	}
	// 自动建表
	db.AutoMigrate(&entity.TaskDo{}, &entity.WorkflowDo{})
}
