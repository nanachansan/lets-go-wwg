package main

import (
	"fmt"
	"time"
)

func main() {
	var anoNasc int
	currentYear := time.Now().Year()

	fmt.Scan(&anoNasc)
	currentAge := currentYear - anoNasc

	fmt.Println(currentAge)
}
