package behaviortree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Tree 行为树
type Tree struct {
	Node
}

// Load 加载
func (t *Tree) Load(filename string) bool {

	b, _ := ioutil.ReadFile(filename)
	err := json.Unmarshal(b, t)
	if err != nil {
		fmt.Printf("behaviortree Load failed : %s", err.Error())
		return false
	}

	return true
}

// Save 保存
func (t *Tree) Save(filename string) bool {

	treeAJSON, _ := json.MarshalIndent(t, "", "  ")
	err := ioutil.WriteFile(filename, treeAJSON, 0)
	if err != nil {
		fmt.Printf("behaviortree Save failed : %s", err.Error())
		return false
	}

	return true
}

// OnEnter 响应登录
func (t *Tree) OnEnter(a *Agent) {
}

// OnLeave 响应离开
func (t *Tree) OnLeave(a *Agent) {
}

// OnRun 响应刷新
func (t *Tree) OnRun(a *Agent) bool {
	return true
}
