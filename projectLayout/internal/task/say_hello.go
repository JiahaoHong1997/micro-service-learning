package task

import (
	"fmt"
	"github.com/gotomicro/ego/task/ejob"
)

func Hello(ctx ejob.Context) error {
	fmt.Printf("hello, world")
	return nil
}
