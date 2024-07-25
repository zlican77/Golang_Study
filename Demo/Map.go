package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println(m)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting value")
	siteName, ok := m["site"]
	fmt.Println(siteName, ok)
	if soteName, ok := m["sote"]; ok {
		fmt.Println(soteName)
	} else {
		fmt.Println("Error ")
	}
}
