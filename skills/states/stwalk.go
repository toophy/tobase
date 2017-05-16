package states

// ActorWalk
type ActorWalk struct {
}

// Check 检查能否进入状态
func (a *ActorWalk) Check(data interface{}) bool {
	return false
}

// OnEnter 进入状态
func (a *ActorWalk) OnEnter(data interface{}) {
}

// OnLeave 离开状态
func (a *ActorWalk) OnLeave(data interface{}) {
}

// OnRun 正在状态中
func (a *ActorWalk) OnRun(data interface{}) {
}
