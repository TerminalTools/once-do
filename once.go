package once

import (
	"errors"
	"sync/atomic"
)

type Once struct {
	ongoing  uint32
	function func()
}

func NewOnce(function func()) (object *Once, newErr error) {
	if function == nil {
		return nil, errors.New("function cannot be empty")
	}

	object = &Once{
		ongoing:  0,
		function: function,
	}
	return object, nil
}

// Do goroutine中只有一个任务在执行, 多次调用时跳过执行
func (self *Once) Do() {
	if atomic.LoadUint32(&self.ongoing) == 1 {
		return
	}

	atomic.StoreUint32(&self.ongoing, 1)
	defer atomic.StoreUint32(&self.ongoing, 0)
	self.function()
}
