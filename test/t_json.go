package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//slice切片转json
	sliceData := make([]interface{}, 0)
	sliceData = append(sliceData, "沉默小管")
	sliceData = append(sliceData, 18)
	fmt.Println(sliceData)
	bytes1, _ := json.Marshal(sliceData)
	stringData1 := string(bytes1)
	fmt.Println(stringData1)

	var result1 []interface{}
	//json转为slice数据结构
	json.Unmarshal([]byte(stringData1), &result1)
	fmt.Println(result1, "--slice")

}
