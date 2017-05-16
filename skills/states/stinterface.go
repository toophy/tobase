package states

// IActorFsmState Actor有限状态接口
type IActorFsmState interface {
	// Check 检查能否进入状态
	Check(data interface{}) bool
	// OnEnter 进入状态
	OnEnter(data interface{})
	// OnLeave 离开状态
	OnLeave(data interface{})
	// OnRun 正在状态中
	OnRun(data interface{})
}

// IActorCondition Actor简易条件接口
type IActorCondition interface {
	// CheckCond 检查条件
	CheckCond(data interface{}) bool
}
