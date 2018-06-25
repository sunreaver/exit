package exit

// Channel 退出信号
type Channel interface{}

// NewChannel NewChannel
func NewChannel() chan *Channel {
	c := make(chan *Channel)
	return c
}

// GroupValue define 退出信号和等待时间
type GroupValue struct {
	C     chan *Channel
	Delay chan *Channel
}

// NotifyExit 通知注册方退出信号到达
func (g *GroupValue) NotifyExit() (done func()) {
	go func() {
		defer func() {
			recover()
		}()
		close(g.C)
	}()

	delay := g.Delay

	return func() {
		select {
		case <-delay:
		}
		return
	}
}
