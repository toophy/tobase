package behaviortree

// Selector 选择组合
// -------------------------------------------------------------
type Selector struct {
	Node
}

// OnEnter 响应登录
func (s *Selector) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *Selector) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *Selector) OnRun(a *Agent) bool {
	for k := range s.Childs {
		if s.Childs[k].OnRun(a) == true {
			return true
		}
	}
	return false
}
