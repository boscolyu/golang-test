package data

import (
	"fmt"
)

const (
        SUCCESS = iota
        FAIL 
        FAIL_INVALID_PARAMETER
        FAIL_DB_CONNECTION
)

func main() {
	fmt.Println(SUCCESS)
	fmt.Println(FAIL)
	fmt.Println(FAIL_DB_CONNECTION)
}
