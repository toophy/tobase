package behaviortree

func init() {
	AppendNodeCreator("Sequence", func() INode { return new(Sequence) })
}

// Sequence 顺序组合
// -------------------------------------------------------------
type Sequence struct {
	Node
}

// OnEnter 响应登录
func (s *Sequence) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *Sequence) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *Sequence) OnRun(a *Agent) bool {
	for k := range s.Childs {
		if s.Childs[k].OnRun(a) == false {
			return false
		}
	}
	return true
}
