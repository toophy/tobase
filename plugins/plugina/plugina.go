// Package plugina 插件b
package plugina

// Plugina 插件a
type Plugina struct {
	Age int32
}

// Name 插件a名称
func (p *Plugina) Name() string {
	return "plugina"
}

// Init 初始化
func (p *Plugina) Init() error {
	return nil
}
