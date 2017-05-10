// Package tobase 基础插件类练习
package tobase

// IPlugin Plugin接口
type IPlugin interface {
	Name() string
	Init() error
}
