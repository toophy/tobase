package behaviortree

var (
	// MaxRoleAttr 属性数量
	MaxRoleAttr = 10
)

// Agent 宿主代理
type Agent struct {
	Name    string
	Attrs   []int64
	CurrPos Vec3
	Target  *Agent
	Stacks  TreeStack
}

// GetName 名称
func (a *Agent) GetName() string {
	return a.Name
}

// GetTarget 目标Agent
func (a *Agent) GetTarget() *Agent {
	return a.Target
}

// CheckAttr 检查属性ID
func (a *Agent) CheckAttr(i int) bool {
	return i > 0 && i < MaxRoleAttr
}

// InitValue 初始化值
func (a *Agent) InitValue(i int, v int64) {
	if a.CheckAttr(i) {
		if a.Attrs == nil {
			a.Attrs = make([]int64, MaxRoleAttr)
		}
		a.Attrs[i-1] = v
	}
}

// GetValue 获取属性值
func (a *Agent) GetValue(i int) int64 {
	if a.CheckAttr(i) {
		return a.Attrs[i-1]
	}
	return 0
}

// SetValue 设置属性
func (a *Agent) SetValue(i int, v int64) {
	if a.CheckAttr(i) {
		a.Attrs[i-1] = v
	}
}

// GetStack 获取当前行为树堆栈
func (a *Agent) GetStack() *TreeStack {
	return &a.Stacks
}

// TreeStack 行为树内存栈
// 1. 行为树名称
// 2. 执行的层数
// 3. 第一层, 运行状态, 数组下标, 本层保存的数据(共调用的节点使用), 每种节点用了不同的数据, 只需要节点知道
// 4. 第二层, 运行状态, 数组下标, 本层保存的数据(共调用的节点使用)
// 5. ...
// 6. 当并行的时候, 有可能一层有多个节点对象, 怎么处理这个并行呢?
// 7. 栈也是做成树么?
// 8. 每个栈都有一个Prev和Next, 用来区分层次, 而且可以有多个, 也就是有读个 Next, Prev只有一个
type TreeStack struct {
	Running   bool
	MyAgent   *Agent
	TreeName  string
	CurrLayer *TreeStackLayer
	Layers    []*TreeStackLayer
}

// Push 压入一层栈
func (t *TreeStack) Push(running bool, param interface{}) {
	tsl := &TreeStackLayer{Running: running, Idx: len(t.Layers), Prev: t.CurrLayer, Params: param}
	t.Layers = append(t.Layers, tsl)
	// 压入之后要回调么?
}

// Pop 弹出一层
func (t *TreeStack) Pop() {
	if len(t.Layers) > 0 {
		// 离开之前要回调么?
		t.Layers = t.Layers[:len(t.Layers)-1]
	}
}

// GetLayers 获取指定层堆栈
func (t *TreeStack) GetLayers(idx int) *TreeStackLayer {
	if idx >= 0 && idx < len(t.Layers) {
		return t.Layers[idx]
	}
	return nil
}

// GetLayerSize 获取堆栈层数
func (t *TreeStack) GetLayerSize() int {
	return len(t.Layers)
}

// TreeStackLayer 行为树栈层
type TreeStackLayer struct {
	Running bool
	Idx     int
	Prev    *TreeStackLayer   // 单独上层
	Nexts   []*TreeStackLayer // 多个下层
	Params  interface{}
}
