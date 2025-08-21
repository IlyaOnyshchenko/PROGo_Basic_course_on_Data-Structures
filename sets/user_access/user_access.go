package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	permissions := make(map[string]struct{})
	for i := 0; i < n; i++ {
		var path string
		fmt.Scan(&path)
		permissions[path] = struct{}{}
	}

	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var request string
		fmt.Scan(&request)
		if _, ok := permissions[request]; ok {
			fmt.Println("YES")
			continue
		}
		parts := strings.Split(request, "/")
		path := ""
		for i := 1; i < len(parts); i++ {
			if path != "" {
				path += "/"
			}
			path += parts[i]
			if _, ok := permissions["/"+path]; ok {
				fmt.Println("YES")
				goto nextRequest
			}
		}

		fmt.Println("NO")

	nextRequest:
		continue
	}
}
