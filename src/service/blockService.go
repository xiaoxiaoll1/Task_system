// Package service 阻塞式执行，该方法已废弃
package service

//// BFS 获取阻塞式的任务执行顺序
//func BFS(root *entity.TaskBo) []*entity.TaskBo {
//	q := util.NewQueue()
//	q.Offer(root)
//	visited := make(map[string]*entity.TaskBo)
//	all := make([]*entity.TaskBo, 0)
//	for q.Size() > 0 {
//		qSize := q.Size()
//		for i := 0; i < qSize; i++ {
//			//pop vertex
//			currVert := q.Poll().(*entity.TaskBo)
//			if _, ok := visited[currVert.Name]; ok {
//				continue
//			}
//			visited[currVert.Name] = currVert
//			all = append(all, currVert)
//			for _, val := range currVert.Children {
//				if _, ok := visited[val.Name]; !ok {
//					q.Offer(val) //add child
//				}
//			}
//		}
//	}
//	return all
//}
//
//
//func doTasks(tasks []*entity.TaskBo) {
//	for _, t := range tasks {
//		time.Sleep(5 * time.Second)
//		fmt.Printf("当前执行任务%s\n", t.Name)
//	}
//}
