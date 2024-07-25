package main

import (
	"fmt"
	"retriver/mock"
	"retriver/real"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriver
	r = mock.Retriver{Contents: "This is a mockretriver"}
	r = real.Retriver{}
	fmt.Println(download(r))
}
