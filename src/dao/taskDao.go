package dao

import (
	"gorm.io/gorm"
	"taskSystem/entity"
)

var db *gorm.DB

func InitDao(database *gorm.DB)  {
	db = database
}

func CreateTask(task *entity.TaskDo) (uint, error) {
	res := db.Create(task)
	return task.ID, res.Error
}

func FindUnDoTasks() (res []entity.TaskDo) {
	db.Where("task_status = ?", "-1").Find(&res)
	return res
}

func GetTask(name string) (res entity.TaskDo) {
	// 获取第一条匹配的记录
	db.Where("name = ?", name).First(&res)
	return res
}

func DeleteTask(name string) {
	// 获取第一条匹配的记录
	db.Where("name = ?", name).Delete(&entity.TaskDo{})
}

func UpdateTask(task *entity.TaskDo) {
	db.Save(&task)
}

func FinishTask(name string)  {
	db.Model(&entity.TaskDo{}).Where("name = ?", name).Update("task_status", "1")
}
