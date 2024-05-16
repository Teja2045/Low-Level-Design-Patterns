package main

import (
	"fmt"
	functionalOptions "go-LLD/functional-options"
)

func main() {
	server := functionalOptions.NewServer(functionalOptions.WithMaxConns(100), functionalOptions.WithTLS(true))

	fmt.Println(server)
}
