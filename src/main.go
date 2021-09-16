package main

import (
	"encoding/json"
	"log"
	_ "taskSystem/config"
	"taskSystem/dao"
	"taskSystem/entity"
	"taskSystem/service"
)
func main()  {

	//task := entity.TaskDo{
	//	Name: "task7",
	//	WorkflowId: 3,
	//	TaskStatus: -1,
	//	Parents: "{1, 2, 3}",
	//	Children: "{4, 5}",
	//}
	//dao.CreateTask(&task)
	//dao.DeleteTask("task3")
	//res := dao.GetTask("task4")
	//res.Name = "wo"
	//dao.UpdateTask(&res)
	//fmt.Println(res)

	taskbo := entity.TaskBo{
		TaskStatus: -1,
		Name: "task1",
		WorkflowId: 1,
	}
	workflow := entity.WorkflowBo{
		WorkflowStatus: -1,
		Tasks: []*entity.TaskBo{&taskbo},
	}

	ch := make(chan string, 5)
	// 先获取一条可行的执行链，存入数据库中
	allTasks := service.BFSNew(&workflow)
	bytes, err := json.Marshal(allTasks)
	if err != nil {
		log.Println("序列化数组时发生错误")
	}
	workflowDo := entity.WorkflowDo{
		WorkflowStatus: -1,
		Tasks: string(bytes),
		Name: "workflow1",
	}
	dao.CreateWorkflow(&workflowDo)

	service.ExecuteWorkflow(allTasks, ch)
	service.TaskListener(ch)


}
