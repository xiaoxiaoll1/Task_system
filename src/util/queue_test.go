// queue 单元测试
package util

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}
	fmt.Println("size:",queue.Size())
	fmt.Println("移除最前面的元素：",queue.Poll())
	fmt.Println("size:",queue.Size())
	fmt.Println("清空：",queue.Clear())
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Size())

}




