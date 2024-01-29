package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	size := 1
	if len(os.Args) > 1 {
		var err error
		size, err = strconv.Atoi(os.Args[1:2][0])
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < size; i++ {
		id := uuid.NewString()
		fmt.Println(id)
	}
}
