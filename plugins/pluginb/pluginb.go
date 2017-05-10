// Package pluginb 插件A
package pluginb

import "fmt"

// Pluginb 简单模块对象
type Pluginb struct {
	Weight string
}

// Name Pluginb 获取名称接口实现
func (o *Pluginb) Name() string {
	return "Pluginb"
}

// Init Pluginb 模块初始化接口实现
func (o *Pluginb) Init() error {
	fmt.Printf("%v", o)
	return nil
}

// Eat Pluginb 吃
func (o *Pluginb) Eat(food string) error {
	fmt.Printf("Pluginb : 吃%s", food)
	return nil
}

// Walk Pluginb 行走
func (o *Pluginb) Walk(speed int32, x float32, y float32) error {
	fmt.Printf("Pluginb : 行走")

	return nil
}
