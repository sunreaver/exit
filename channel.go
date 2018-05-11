package exit

// Channel 退出信号
type Channel interface{}

// GroupValue define 退出信号和等待时间
type GroupValue struct {
	c     chan *Channel
	delay <-chan *Channel
}

func (g *GroupValue) do() (done func()) {
	go func() {
		defer func() {
			recover()
		}()
		close(g.c)
	}()

	delay := g.delay

	return func() {
		select {
		case <-delay:
			return
		}
	}
}
