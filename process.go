package exit

import (
	"container/list"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	data = Group{
		NeedExit: true,
		WG:       &sync.WaitGroup{},
		Data:     list.New(),
	}
)

func notify() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			go func() {
				if data.Exit() {
					os.Exit(0)
				}
			}()
		}
	}
}
