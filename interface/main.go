package main

import (
	"fmt"
	"interface/inter"
)

type retriever interface {
	Get(contents string) string
}

type poster interface {
	Post(url string, contents map[string]string) string
}

type postRetriever interface {
	retriever
	poster
}

func main() {
	var s postRetriever
	s = inter.PostRetriever{Contents: "interface"}
	fmt.Println(s.Get("go"))

	Map := make(map[string]string)
	Map["A"] = "100"
	s.Post("imooc.com", Map)
}
