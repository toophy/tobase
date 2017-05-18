package behaviortree

import "testing"
import "encoding/json"
import "encoding/xml"

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
	t := &Tree{MyAgent: a}

	// 逃离 : and
	taoLiSeq := &Sequence{}
	t.PushNode(taoLiSeq)
	{
		// 逃离条件组合 : or
		taoLiCondSel := &Selector{}
		taoLiSeq.PushNode(taoLiCondSel)
		{
			// 逃离 : 条件1 : 血量<10
			taoLiCond := &CondAttr{"血量低", AttrHP, CondAttrLq, 10}
			taoLiSeq.PushNode(taoLiCond)

			// 逃离 : 条件2 : 自己血量+50<对方血量
			taoLiCond2 := &CondAttrCompare{"血量压制", AttrHP, CondAttrLt, 50}
			taoLiSeq.PushNode(taoLiCond2)
		}

		// 逃离 : 动作节点
		taoLiAction := &Action{false, "逃离"}
		taoLiSeq.PushNode(taoLiAction)
	}

	// 追击 : and
	zhuiSeq := &Sequence{}
	t.PushNode(zhuiSeq)
	{
		// 追击条件组合 : or
		zhuiCondSel := &Selector{}
		zhuiSeq.PushNode(zhuiCondSel)
		{
			// 追击 : 条件1 : 距离<10
			taoLiCond := &CondAttrPosCompare{"目标超出攻击范围", CondAttrLq, 10}
			zhuiSeq.PushNode(taoLiCond)

			// 追击 : 条件2 : 自己距离+50<对方血量
			taoLiCond2 := &CondAttrPosCompare{"目标在追击范围", CondAttrLt, 30}
			zhuiSeq.PushNode(taoLiCond2)
		}

		// 追击 : 动作节点
		zhuiAction := &Action{false, "追击"}
		zhuiSeq.PushNode(zhuiAction)
	}

	return t
}

func TestMonster(t *testing.T) {
	agentA := CreateAgent("A", 100, 20, 15)
	agentB := CreateAgent("B", 100, 25, 10)

	treeA := CreateBehaviorTree(agentA)
	treeB := CreateBehaviorTree(agentB)

	treeAjson, _ := json.MarshalIndent(treeA, "", "  ")
	println(string(treeAjson))

	nTreeA := new(Tree)
	json.Unmarshal(treeAjson, nTreeA)

	treeAJson2, _ := json.MarshalIndent(nTreeA, "", "  ")
	println(string(treeAJson2))

	treeBjson, _ := json.MarshalIndent(treeB, "", "  ")
	println(string(treeBjson))

	treeAXml, _ := xml.MarshalIndent(treeA, "", "  ")
	println(string(treeAXml))
}
