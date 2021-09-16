// Package service
/**
dataService用于取出未完成的workflow
 */
package service

import (
	"encoding/json"
	"log"
	"taskSystem/dao"
	"taskSystem/entity"
)

func FinishUndoWorkflow(name string, ch chan string) {
	workflow := dao.GetWorkflow(name)
	if workflow.WorkflowStatus != -1 {
		log.Println("该工作流已经完成")
		return
	}
	var arr [][]*entity.TaskBo
	json.Unmarshal([]byte(workflow.Tasks), &arr)
	// 取出该workflow对应的任务，选择其中未执行完成的执行
	ExecuteWorkflow(arr, ch)
}


