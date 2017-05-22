package behaviortree

import (
	"os"
	"testing"
)

//
const (
	AttrHP = iota + 1
	AttrPhyAta
	AttrPhyDef
)

func CreateAgent(name string, hp int64, phyata int64, phydef int64) *Agent {
	a := &Agent{Name: name}
	a.InitValue(AttrHP, hp)
	a.InitValue(AttrPhyAta, phyata)
	a.InitValue(AttrPhyDef, phydef)
	return a
}

func CreateBehaviorTree(a *Agent) *Tree {
	t := &Tree{}
	return t
}

// 判断文件/文件存在
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func TestMonster(t *testing.T) {
	agentA := CreateAgent("A", 100, 20, 15)

	treeA := CreateBehaviorTree(agentA)

	treeA.Load("./doc/tree.json")
	treeA.Save("./tmp/tree.json")
}
