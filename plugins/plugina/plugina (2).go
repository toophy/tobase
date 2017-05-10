// Package plugina 插件A
package plugina

import "fmt"

// Plugina 简单模块对象
type Plugina struct {
	Age int
}

// Name Plugina 获取名称接口实现
func (o *Plugina) Name() string {
	return "Plugina"
}

// Init Plugina 模块初始化接口实现
func (o *Plugina) Init() error {
	fmt.Printf("%v", o)
	return nil
}
