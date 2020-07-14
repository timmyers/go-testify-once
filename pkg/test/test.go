//nolint:gochecknoglobals
package test

import (
	"fmt"
	"sync"
)

var once sync.Once
var Count = 0

func Init() {
	once.Do(func() {
		Count = Count + 1
		fmt.Println("Only ONCE")
	})
}
