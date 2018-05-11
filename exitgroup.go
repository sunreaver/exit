package exit

import (
	"container/list"
	"sync"
)

// Group define 退出信号的存储列表
type Group struct {
	sync.RWMutex
	needExit bool
	wg       *sync.WaitGroup
	group    *list.List
}

// Add will add
func (g *Group) Add(c *GroupValue) {
	g.Lock()
	defer g.Unlock()
	g.group.PushBack(c)
}

// Remove will remove c in group
func (g *Group) Remove(c <-chan *Channel) (removed bool) {
	g.Lock()
	defer g.Unlock()
Remove:
	for item := g.group.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(*GroupValue); ok && v.c == c {
			g.group.Remove(item)
			removed = true
			break Remove
		}
	}
	return
}

func (g *Group) exit() bool {
	g.Lock()
	tmpList := g.group
	wg := g.wg
	need := g.needExit

	g.group = list.New()
	g.wg = &sync.WaitGroup{}
	g.needExit = false
	g.Unlock()

	for item := tmpList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(*GroupValue); ok {
			wg.Add(1)
			done := v.do()
			go func() {
				done()
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return need
}
