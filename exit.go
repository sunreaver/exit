package exit

import (
	"sync"
)

var (
	once sync.Once
)

// RegisteExiter will Registe退出信号
// 当进程接到退出信号
// 会写入exit信号
// 并等待业务端delay信号的写入
// 当接收到业务端delay信号，会执行os.Exit(0)
func RegisteExiter(delay <-chan *Channel) (exit <-chan *Channel) {
	once.Do(func() {
		go notify()
	})

	e := make(chan *Channel, 1)
	v := &GroupValue{
		C:     e,
		Delay: delay,
	}
	data.Add(v)
	return e
}

// UnRegistExiter will UnRegist退出信号
func UnRegistExiter(c <-chan *Channel) (unregisted bool) {
	return data.Remove(c)
}
