package skills

import "encoding/json"

// SpaceTime 时空
type SpaceTime struct {
	Name    string             // 名称
	Parent  string             // 前辈
	Time    uint64             // 时间
	Strings map[string]string  // 字符串参数表
	Ints    map[string]int64   // 整数参数表
	Floats  map[string]float64 // 浮点参数表
}

// World 世界
type World struct {
	SpaceTime
	Universes map[string]*Universe // 宇宙们
}

// Universe 宇宙
type Universe struct {
	SpaceTime
	Scenes map[string]*Scene // 场景们
}

// Scene 场景
type Scene struct {
	SpaceTime
}

// RealScene 真实场景
type RealScene struct {
	Scene
	NewName   string           // 新名称
	SceneRect Rect             // 限制区域
	Roles     map[string]*Role // 角色们
	Npcs      map[string]*Npc  // 电脑角色们
}

// Actor 演员,RealScene的标准对象
type Actor struct {
	NewName string     // 演员名称
	CurrPos Vector3D   // 当前位置
	MyScene *RealScene // 场景
	Move    IActorMove // 行走
}

// Actor 行为, 这些行为只能串行么?
// 走
// 跑
// 跳
// 停
// 使用技能攻击
//
//
// 这些行为, 执行起来大多不同, 参数也不同
//
// 行为, 状态机,
// 玩家控制的Actor状态机, 半自动
// 电脑控制的Actor状态机, 自动

// Vector3D 3D坐标点
type Vector3D struct {
	X, Y, Z float32
}

// Rect 3D区域
type Rect struct {
	X, Y, Z float32
}

// IActorMove 行走接口
type IActorMove interface {
	Init(a *Actor, c Vector3D, s float32, d float32)
	Walk(t float32) bool
	Run(t float32) bool
}

// ActorMove 演员行走能力
type ActorMove struct {
	MyActor    *Actor
	TargetPath []Vector3D
	Speed      float32
	Direct     float32
}

// Init 演员行走能力初始化
func (a *ActorMove) Init(myActor *Actor, c Vector3D, s float32, d float32) {
	a.MyActor = myActor
	a.Speed = s
	a.Direct = d
}

// Walk 行走tms时间
func (a *ActorMove) Walk(t float32) bool {
	// 当前朝向, 继续行走
	return false
}

// Run 跑tms时间
func (a *ActorMove) Run(t float32) bool {
	// 当前朝向, 继续跑
	return false
}

// NewRole 新角色,包容所有角色,配备插件
type NewRole struct {
}

// // DoWalk 行走
// func (n *NewRole) DoWalk(speed int32, x float32, y float32) error {
// 	if plugin, ok := p.plugins[i].(INewRolePluginWalk); ok {
// 		err := plugin.Walk(speed, x, y)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // 新角色综合插件
// type (
// 	// INewRolePluginWalk 动作接口:行走
// 	INewRolePluginWalk interface {
// 		Walk(speed int32, x float32, y float32) error
// 	}

// 	// INewRolePluginJump 动作接口:跳
// 	INewRolePluginJump interface {
// 		Jump(speed int32, hight float32) error
// 	}
// )

var (
	gWorld *World
)

func init() {
	gWorld = NewWorld("世界")
}

// OutputJSON 导出世界信息到Json
func (w *World) OutputJSON() string {
	data, _ := json.MarshalIndent(w, "", "  ")
	return string(data)
}

// InputJSON 从Json导入世界信息
func (w *World) InputJSON(data string) {
	err := json.Unmarshal([]byte(data), w)
	if err != nil {
		println("%s", err.Error())
	}
}

// NewWorld 创建世界
func NewWorld(name string) *World {
	if len(name) > 0 {
		p := new(World)
		p.Name = name
		p.Time = 0
		p.Strings = make(map[string]string, 0)
		p.Ints = make(map[string]int64, 0)
		p.Floats = make(map[string]float64, 0)
		p.Universes = make(map[string]*Universe, 0)
		return p
	}
	return nil
}

// NewUniverse 创建宇宙
func NewUniverse(name string, p *World) *Universe {
	if len(name) > 0 && p != nil {
		if _, ok := p.Universes[name]; !ok {
			s := new(Universe)

			s.SpaceTime = p.SpaceTime
			s.Name = name
			s.Scenes = make(map[string]*Scene, 0)

			p.Universes[name] = s
			return s
		}
	}
	return nil
}

// NewScene 创建场景
func NewScene(name string, p *Universe) *Scene {
	if len(name) > 0 && p != nil {
		if _, ok := p.Scenes[name]; !ok {
			s := new(Scene)

			s.SpaceTime = p.SpaceTime
			s.Name = name

			p.Scenes[name] = s
			return s
		}
	}
	return nil
}

// NewRealScene 创建场景
// 只有场景会分布, 其他都集中在中心服
// 所有分布者都拥有完整的世界镜像
func NewRealScene(name string, p *Scene) *RealScene {
	if len(name) > 0 && p != nil {
		s := new(RealScene)

		s.NewName = name
		s.Scene = *p

		return s
	}
	return nil
}
