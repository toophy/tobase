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

// TreeStack 行为树内存栈
// 1. 行为树名称
// 2. 执行的层数
// 3. 第一层, 运行状态, 数组下标, 参数
// 4. 第二层, 运行状态, 数组下标, 参数
// 5. ...
type TreeStack struct {
	Running  bool
	MyAgent  *Agent
	TreeName string
	Layers   []TreeStackLayer
}

// TreeStackLayer 行为树栈层
type TreeStackLayer struct {
	Running bool
	Idx     int
	Params  []interface{}
}
