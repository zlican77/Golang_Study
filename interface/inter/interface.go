package inter

import "fmt"

type PostRetriever struct {
	Contents string
}

func (r PostRetriever) Get(Contents string) string {
	fmt.Println(r.Contents)
	return "ok"
}

func (r PostRetriever) Post(url string, Contents map[string]string) string {
	fmt.Println(url, Contents)
	return "ok"
}
