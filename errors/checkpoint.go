package errors

import (
	"fmt"
	"time"
)

func Check() {
	currentTime := time.Now()
	fmt.Println("Request : ", currentTime.Format("2006-01-02 15:04:05"))
}
