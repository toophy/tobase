package behaviortree

// Tree 行为树
type Tree struct {
	Node
	MyAgent *Agent
}

// Load 加载
func (b *Tree) Load() bool {
	return true
}

// Save 保存
func (b *Tree) Save() bool {
	return true
}

// OnEnter 响应登录
func (b *Tree) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (b *Tree) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (b *Tree) OnRun(a *Agent) bool {
	return true
}
