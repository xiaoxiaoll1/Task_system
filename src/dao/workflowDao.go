package dao

import "taskSystem/entity"

func CreateWorkflow(workflow *entity.WorkflowDo) (uint, error) {
	res := db.Create(workflow)
	return workflow.ID, res.Error
}

func FinishWorkflow(name string)  {
	db.Model(&entity.WorkflowDo{}).Where("name = ?", name).Update("workflow_status", "1")
}

func GetWorkflow(name string) (res entity.WorkflowDo) {
	// 获取第一条匹配的记录
	db.Where("name = ?", name).First(&res)
	return res
}

