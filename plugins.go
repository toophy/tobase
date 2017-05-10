// Package tobase 基础插件类练习
package tobase

// Plugins 插件注册表
type Plugins struct {
	ps []IPlugin
}

// GetName 获取插件名称
func (p *Plugins) GetName(plugin IPlugin) string {
	return plugin.Name()
}

// GetPlugin 通过名称获取插件对象

// DoInit 初始化插件
func (p *Plugins) DoInit(plugin IPlugin) error {
	return plugin.Init()
}

// DoEat 做动作 : 吃
func (p *Plugins) DoEat() error {
	return nil
}

// 对象基本动作接口
type (
	// IPluginEat 动作接口:吃
	IPluginEat interface {
		Eat(food string) error
	}

	// IPluginWalk 动作接口:行走
	IPluginWalk interface {
		Walk(speed int32, x float32, y float32) error
	}

	// IPluginJump 动作接口:跳
	IPluginJump interface {
		Jump(speed int32, hight float32) error
	}
)
