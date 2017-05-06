// Package pluginb 插件b
package pluginb

// Pluginb 插件b
type Pluginb struct {
	Age int32
}

// Name 插件b名称
func (p *Pluginb) Name() string {
	return "pluginb"
}

// Init 初始化
func (p *Pluginb) Init() error {
	return nil
}
