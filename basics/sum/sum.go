package main

import "fmt"

func sum(x, y int) (ans int) {
	ans = x + y
	return
}

func main() {
	var x int = 9
	var y int = 12
	var ans int = sum(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, ans)
}
