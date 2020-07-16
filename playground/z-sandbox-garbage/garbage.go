// B''H

/*
go mod init sandbox/garbage
go run garbage.go
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
}
