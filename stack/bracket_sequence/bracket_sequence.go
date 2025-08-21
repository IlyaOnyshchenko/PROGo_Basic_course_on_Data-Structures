package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scnr := bufio.NewScanner(os.Stdin)
	_ = scnr.Scan()
	str := scnr.Text()
	balance := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			fmt.Println("NO")
			break
		}
	}
	if balance == 0 {
		fmt.Println("YES")
	} else if balance > 0 {
		fmt.Println("NO")
	}
}
