package service

import (
	"fmt"
	"sync"
	"taskSystem/dao"
	"taskSystem/entity"
	"taskSystem/util"
	"time"
)

// BFSNew Dag图的广度优先遍历
func BFSNew(workflow *entity.WorkflowBo) [][]*entity.TaskBo {
	inDegree := make(map[string]int)
	relation := make(map[string]*entity.TaskBo)
	// 计算出每个任务的前驱任务---入度
	for _, task := range workflow.Tasks {
		relation[task.Name] = task
		inDegree[task.Name] = len(task.Parents)
	}
	q := util.NewQueue()
	// 先将入度为0的任务添加到queue中
	for s, i := range inDegree {
		if i == 0 {
			q.Offer(relation[s])
		}
	}
	visited := make(map[string]bool)
	all := make([][]*entity.TaskBo, 0)
	for q.Size() > 0 {
		qSize := q.Size()
		tmp := make([]*entity.TaskBo, 0)
		for i := 0; i < qSize; i++ {
			currVert := q.Poll().(*entity.TaskBo)
			// visited map防止死循环
			visited[currVert.Name] = true
			tmp = append(tmp, currVert)
			for _, val := range currVert.Children {
				// 将后继任务到入度减1，变为0了就加入队列
				inDegree[val.Name]--
				if inDegree[val.Name] == 0 && !visited[val.Name] {
					q.Offer(relation[val.Name])
				}
			}
		}
		all = append(all, tmp)
	}
	// 最后的结果是一个二维数组，同一层的任务没有先后制约，可以并发执行
	return all
}

func ExecuteWorkflow(all [][] *entity.TaskBo, ch chan string)  {
	for _, layer := range all {
		fmt.Println("------------------")
		doTasksNew(layer, ch)
	}
}


//并发执行
func doTasksNew(tasks []*entity.TaskBo, ch chan string) {
	var wg sync.WaitGroup
	for _, v := range tasks {
		if v.TaskStatus != -1 {
			continue
		}
		wg.Add(1)
		// 调用goroutine同时传入当前参数，避免循环引用
		go func(t *entity.TaskBo) {
			defer wg.Done()
			time.Sleep(5 * time.Second)
			fmt.Printf("任务%s已经执行", t.Name)
			ch <- t.Name
		}(v) //notice
	}
	wg.Wait()
}

// TaskListener 监听任务完成的事件，将数据库中任务状态设置为已完成
func TaskListener(ch chan string)  {
	for true {
		// IO多路复用，阻塞式轮询获取channel中的内容
		select {
		case name, isOpen := <- ch:
			if !isOpen {
				return
			}
			dao.FinishTask(name)
		}
	}

}