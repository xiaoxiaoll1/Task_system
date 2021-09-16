package entity

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T)  {
	var urls = []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://yahoo.com",
		"http://twitter.com",
		"http://live.com",
	}
	fmt.Println(urls)
	urlsJson, _ := json.Marshal(urls)
	jsonStr := string(urlsJson)
	fmt.Println(jsonStr)
	var arr []string
	json.Unmarshal([]byte(jsonStr), &arr)
	fmt.Println(arr)
}

// 测试数组与json的转化
func TestConvert(t *testing.T)  {
	test1 := TaskBo{
		Name: "test1",
	}
	test2 := TaskBo{
		Name: "test2",
	}
	test3 := TaskBo{
		Name: "test3",
	}
	test4 := TaskBo{
		Name: "test4",
	}

	target := TaskBo{
		Name: "target",
	}
	target.Parents = []*TaskBo{&test1, &test2}
	target.Children = []*TaskBo{&test3, &test4}
	// 指针变量的json仍然表示的结构体内容
	byteStr, err := json.Marshal(target.Parents)
	if err != nil {
		fmt.Println("转化json出错")
	}
	fmt.Println(string(byteStr))
	jsonStr := string(byteStr)
	var arr []TaskBo
	json.Unmarshal([]byte(jsonStr), &arr)
	fmt.Println(arr[0].Name)


}

