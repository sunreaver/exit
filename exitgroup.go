package exit

import (
	"container/list"
	"sync"
)

// Group define 退出信号的存储列表
type Group struct {
	sync.RWMutex
	NeedExit bool
	WG       *sync.WaitGroup
	Data     *list.List
}

// Add will add
func (g *Group) Add(c *GroupValue) {
	g.Lock()
	defer g.Unlock()
	g.Data.PushBack(c)
}

// Remove will remove c in group
func (g *Group) Remove(c <-chan *Channel) (removed bool) {
	g.Lock()
	defer g.Unlock()
Remove:
	for item := g.Data.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(*GroupValue); ok && v.C == c {
			g.Data.Remove(item)
			removed = true
			break Remove
		}
	}
	return
}

// Exit 执行退出流程
func (g *Group) Exit() bool {
	g.Lock()
	tmpList := g.Data
	wg := g.WG
	need := g.NeedExit

	g.Data = list.New()
	g.WG = &sync.WaitGroup{}
	g.NeedExit = false
	g.Unlock()

	for item := tmpList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(*GroupValue); ok {
			wg.Add(1)
			done := v.NotifyExit()
			go func() {
				done()
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return need
}
