package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `{"company":"zlican","subjects":["go","js","python","mysql"],"isok":true,"price":666}` //json

	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(tmp)
}
