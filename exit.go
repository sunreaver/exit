package exit

import (
	"sync"
)

var (
	once sync.Once
)

// RegistExiter will Registe退出信号
// 当进程接到退出信号, 会写入exit信号
// 并等待业务端delay信号的写入
// 当接收到业务端delay信号，会执行os.Exit(0)
func RegistExiter() (exiter <-chan *Channel, delay chan<- *Channel) {
	once.Do(func() {
		go notify()
	})

	e := make(chan *Channel, 1)
	d := make(chan *Channel, 1)
	v := &GroupValue{
		C:     e,
		Delay: d,
	}
	data.Add(v)
	return e, d
}

// UnRegistExiter will 取消监听退出信号
// 参数exiter必须为RegistExiter方法返回的exiter参数
func UnRegistExiter(exiter <-chan *Channel) (unregisted bool) {
	return data.Remove(exiter)
}
