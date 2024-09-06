package main

import (
	"encoding/json"
	"fmt"
)

type IT struct { //可二次编码，使得json输出的是小写
	Company  string   `json:"-"`        //不输出该字段内容
	Subjects []string `json:"subjects"` //二次编码
	Isok     bool     `json:",string"`  //value格式转化成string再输出
	Price    float64  `json:"price,string"`
}

func main() {
	s := IT{"zlican", []string{"go", "js", "python", "mysql"}, true, 666}
	jsonS, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", string(jsonS))
	//{"Company":"zlican","Subjects":["go","js","python","mysql"],"Isok":true,"Price":666}
}
