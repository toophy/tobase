package states

// ActorRun
type ActorRun struct {
}

// Check 检查能否进入状态
func (a *ActorRun) Check(data interface{}) bool {
	return false
}

// OnEnter 进入状态
func (a *ActorRun) OnEnter(data interface{}) {
}

// OnLeave 离开状态
func (a *ActorRun) OnLeave(data interface{}) {
}

// OnRun 正在状态中
func (a *ActorRun) OnRun(data interface{}) {
}

// ActorCond
type CondEq struct {
}

//
