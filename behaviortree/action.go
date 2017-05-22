package behaviortree

import (
	"fmt"
)

func init() {
	AppendNodeCreator("Action", func() INode { return new(Action) })
}

// Action 动作:叶子
// -------------------------------------------------------------
type Action struct {
	NodeLeaf
	Name string
}

// OnEnter 响应登录
func (s *Action) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *Action) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *Action) OnRun(a *Agent) bool {
	fmt.Printf("[%s] 执行动作 : %s", a.GetName, s.Name)
	return true
}
