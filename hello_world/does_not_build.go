package main

import "fmt"

func not_main() {
	fmt.Println("This will not buld because it;s in a `package main` but there is no func main")
}
