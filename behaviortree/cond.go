package behaviortree

const (
	// CondAttrNq !=
	CondAttrNq = iota
	// CondAttrLt <
	CondAttrLt
	// CondAttrLq <=
	CondAttrLq
	// CondAttrEq ==
	CondAttrEq
	// CondAttrGq >=
	CondAttrGq
	// CondAttrGt >
	CondAttrGt
)

// CondAttr 条件:属性条件
// --------------------------------------------------------
type CondAttr struct {
	Name     string
	AttrIdx  int
	Operator int
	Value    int64
}

// IsRunning 正在运行中
func (s *CondAttr) IsRunning() bool {
	return false
}

// OnEnter 响应登录
func (s *CondAttr) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *CondAttr) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *CondAttr) OnRun(a *Agent) bool {
	switch s.Operator {
	case CondAttrNq:
		return a.GetValue(s.AttrIdx) != s.Value
	case CondAttrLt:
		return a.GetValue(s.AttrIdx) < s.Value
	case CondAttrLq:
		return a.GetValue(s.AttrIdx) <= s.Value
	case CondAttrEq:
		return a.GetValue(s.AttrIdx) == s.Value
	case CondAttrGq:
		return a.GetValue(s.AttrIdx) >= s.Value
	case CondAttrGt:
		return a.GetValue(s.AttrIdx) > s.Value
	}
	return false
}

// CondAttrTarget 条件:目标属性条件
// --------------------------------------------------------
type CondAttrTarget struct {
	Name     string
	AttrIdx  int
	Operator int
	Value    int64
}

// IsRunning 正在运行中
func (s *CondAttrTarget) IsRunning() bool {
	return false
}

// OnEnter 响应登录
func (s *CondAttrTarget) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *CondAttrTarget) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *CondAttrTarget) OnRun(a *Agent) bool {
	t := a.GetTarget()
	if t == nil {
		return false
	}

	switch s.Operator {
	case CondAttrNq:
		return t.GetValue(s.AttrIdx) != s.Value
	case CondAttrLt:
		return t.GetValue(s.AttrIdx) < s.Value
	case CondAttrLq:
		return t.GetValue(s.AttrIdx) <= s.Value
	case CondAttrEq:
		return t.GetValue(s.AttrIdx) == s.Value
	case CondAttrGq:
		return t.GetValue(s.AttrIdx) >= s.Value
	case CondAttrGt:
		return t.GetValue(s.AttrIdx) > s.Value
	}
	return false
}

// CondAttrCompare 条件:自己和目标属性条件(互相比较)
// --------------------------------------------------------
type CondAttrCompare struct {
	Name     string
	AttrIdx  int
	Operator int
	Value    int64
}

// IsRunning 正在运行中
func (s *CondAttrCompare) IsRunning() bool {
	return false
}

// OnEnter 响应登录
func (s *CondAttrCompare) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *CondAttrCompare) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *CondAttrCompare) OnRun(a *Agent) bool {
	t := a.GetTarget()
	if t == nil {
		return false
	}

	switch s.Operator {
	case CondAttrNq:
		return a.GetValue(s.AttrIdx)+s.Value != t.GetValue(s.AttrIdx)
	case CondAttrLt:
		return a.GetValue(s.AttrIdx)+s.Value < t.GetValue(s.AttrIdx)
	case CondAttrLq:
		return a.GetValue(s.AttrIdx)+s.Value <= t.GetValue(s.AttrIdx)
	case CondAttrEq:
		return a.GetValue(s.AttrIdx)+s.Value == t.GetValue(s.AttrIdx)
	case CondAttrGq:
		return a.GetValue(s.AttrIdx)+s.Value >= t.GetValue(s.AttrIdx)
	case CondAttrGt:
		return a.GetValue(s.AttrIdx)+s.Value > t.GetValue(s.AttrIdx)
	}
	return false
}

// CondAttrPosCompare 条件:自己和目标距离条件(互相比较)
// --------------------------------------------------------
type CondAttrPosCompare struct {
	Name     string
	Operator int
	Value    int64
}

// IsRunning 正在运行中
func (s *CondAttrPosCompare) IsRunning() bool {
	return false
}

// OnEnter 响应登录
func (s *CondAttrPosCompare) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (s *CondAttrPosCompare) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (s *CondAttrPosCompare) OnRun(a *Agent) bool {
	t := a.GetTarget()
	if t == nil {
		return false
	}

	dis := int64(a.CurrPos.Dist(&t.CurrPos))

	switch s.Operator {
	case CondAttrNq:
		return dis != s.Value
	case CondAttrLt:
		return dis < s.Value
	case CondAttrLq:
		return dis <= s.Value
	case CondAttrEq:
		return dis == s.Value
	case CondAttrGq:
		return dis >= s.Value
	case CondAttrGt:
		return dis > s.Value
	}
	return false
}
