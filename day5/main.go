package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	dayFive(true)
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
