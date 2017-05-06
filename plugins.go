// Package tobase 基本库练习
package tobase

import (
	"errors"
)

// IPlugin 插件接口
type IPlugin interface {
	// Name 获取插件名称
	Name() string

	// Init 插件初始化
	Init() error
}

// Plugins 插件注册表
type Plugins struct {
	ps map[string]IPlugin
}

// InitPlugins 初始化插件注册表
func (p *Plugins) InitPlugins() error {
	p.ps = make(map[string]IPlugin, 0)
	return nil
}

// GetName 获取插件名称
func (p *Plugins) GetName(plugin IPlugin) string {
	return plugin.Name()
}

// RegistePlugin 注册插件
func (p *Plugins) RegistePlugin(plugin IPlugin) error {
	if plugin != nil && plugin.Name() != "" {
		if _, ok := p.ps[plugin.Name()]; !ok {
			p.ps[plugin.Name()] = plugin
			return nil
		}
	}

	return errors.New("Invalid plugin name")
}

// GetPlugin 通过名称获取插件
func (p *Plugins) GetPlugin(name string) IPlugin {
	println("get")
	if _, ok := p.ps[name]; ok {
		println("geta")
		return p.ps[name]
	}
	println("getb")
	return nil
}
