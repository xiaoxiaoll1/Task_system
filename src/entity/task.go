package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

// TaskBo 业务对象
type TaskBo struct {
	ID uint
	Name string
	WorkflowId uint
	TaskStatus int
	Parents []*TaskBo
	Children []*TaskBo
}

// TaskDo 数据库对象
type TaskDo struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);uniqueIndex"`
	WorkflowId uint
	TaskStatus int
	Parents string
	Children string
}

func (cur *TaskBo) AddEdge(parent *TaskBo)  {
	parent.Children = append(parent.Children, cur)
	cur.Parents = append(cur.Parents, parent)
}

// ConvertToDo 将taskBo转为taskDo
func (bo *TaskBo) ConvertToDo() (do *TaskDo) {
	do.Name = bo.Name
	do.WorkflowId = bo.WorkflowId
	do.TaskStatus = bo.TaskStatus
	pStr, _ := json.Marshal(bo.Parents)
	cStr, _ := json.Marshal(bo.Children)
	do.Parents = string(pStr)
	do.Children = string(cStr)
	return do
}