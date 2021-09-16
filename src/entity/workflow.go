package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

type WorkflowBo struct {
	ID uint
	Name string
	Tasks []*TaskBo
	WorkflowStatus int
}


type WorkflowDo struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);uniqueIndex"`
	// 存储的是task id数组对应的json
	Tasks string
	WorkflowStatus int
}

func (workflow *WorkflowBo) addTasks(task *TaskBo)  {
	workflow.Tasks = append(workflow.Tasks, task)
}

// ConvertToWorkflowDo 将workflowBo转为workflowDo
func (bo *WorkflowBo) ConvertToWorkflowDo() (do *WorkflowDo) {
	do.Name = bo.Name
	do.WorkflowStatus = bo.WorkflowStatus
	pStr, _ := json.Marshal(bo.Tasks)
	do.Tasks = string(pStr)
	return do
}


