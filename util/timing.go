package util

import (
	"fmt"
	"time"
)

func MeasureExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
