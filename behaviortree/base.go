package behaviortree

import "encoding/json"

// 控制节点
//   1. 串行节点
//   2. 并行节点
//   3. 串行找到一个可行节点
//   4. 选择节点
//   5. 描述节点
// 前提条件
//   1. 进入节点的条件
// 动作节点
//   1. 入住节点后执行的动作
//
// Agent
//    节点和外界的代理(数据Input和Output的代理), 减少耦合
//

//
//  1. composite       组合
//     1.1 Sequence    顺序     有多个子节点
//  2. decorator       修饰     它有且只有一个子节点, 将子节点的结果传递给父节点, 停止, 重复子节点, 比如 Inverter(反相器,NOT)
//  3. Leaf            叶子     没有子节点
//
// 用游戏代码去类比，可以将合成和修饰节点当作函数、、分支结构和循环结构，
// 还有其它编程语言的结构，用它们来定义你代码的逻辑。
// 而叶子节点就像游戏中具体的代码逻辑，
// 会让你的AI角色做一些实际上的事情或者检测它们的状态或场景

// Sequence
//    顺序执行每个子节点, 有一个失败, 就向父节点返回失败, 全节点成功, 就向父节点返回成功, 结束
//    相当于 And, 一般前面都是条件, 最后一个是动作

// Selector
//    顺序执行每个子节点, 有一个成功, 就向父节点返回成功, 全节点失败才是返回失败, 结束
//    相当于 Or, 一般前面是最顺利的操作, 后面是一个个更困难的操作
//

// 各种修饰, 是, 否, 循环n次

// 平行
//    顺序平行执行所有子节点

// 叶子
//    条件节点 , 动作节点, 直接执行, 动作一般返回true, 条件则返回各种实际判断结果
//

//
// 一张树说明表
//      主要是给编辑器看, 包括, 树名称, 树struct名, 树用途(AI,任务,技能...) 树的Agent是谁(可以多种,比如玩家,怪物,帮会,国家等)
//
// 一张节点说明表
//      主要是节点名称, 节点介绍, 节点的Agent是谁, 节点struct类型, 节点是否是热点调用(脚本不能用在热点)
//
// 树中明确的属性
//      树名称, 树对应的struct类型, 节点对应的struct类型, 节点说明
//

// INode 行为树节点接口
type INode interface {
	IsRunning() bool
	OnEnter(a *Agent)
	OnLeave(a *Agent)
	OnRun(a *Agent) bool
}

// Node 节点基础
type Node struct {
	Running bool
	Childs  []INode
}

// IsRunning 正在运行中
func (n *Node) IsRunning() bool {
	return n.Running
}

// PushNode 增加节点
func (n *Node) PushNode(c INode) bool {
	if n.Childs == nil {
		n.Childs = make([]INode, 0)
	}
	n.Childs = append(n.Childs, c)
	return true
}

// PopNode 删除节点
func (n *Node) PopNode(c INode) bool {
	if n.Childs == nil {
		return false
	}

	ret := false
	cs := make([]INode, 0)
	for k := range n.Childs {
		if n.Childs[k] != c {
			cs = append(cs, n.Childs[k])
			ret = true
		}
	}
	if ret {
		n.Childs = cs
	}

	return ret
}

// UnmarshalJSON 反序列化JSON为对象
func (n *Node) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawMessagesForColoredThings []*json.RawMessage
	err = json.Unmarshal(*objMap["Childs"], &rawMessagesForColoredThings)
	if err != nil {
		return err
	}

	n.Childs = make([]INode, len(rawMessagesForColoredThings))

	for index, rawMessage := range rawMessagesForColoredThings {
		var m map[string]interface{}
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		p := gNodeCreators[m["type"].(string)]()
		err := json.Unmarshal(*rawMessage, &p)
		if err != nil {
			return err
		}

		n.Childs[index] = p
	}

	return nil
}

type funcCreateNode func() INode // 声明了一个函数类型
var gNodeCreators map[string]funcCreateNode

func init() {
	if gNodeCreators == nil {
		gNodeCreators = make(map[string]funcCreateNode, 0)
	}
}

// AppendNodeCreator 附加一个节点创建函数
func AppendNodeCreator(name string, f func() INode) {
	if gNodeCreators == nil {
		gNodeCreators = make(map[string]funcCreateNode, 0)
	}
	gNodeCreators[name] = f
}
