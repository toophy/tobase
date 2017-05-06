package tobase

import (
	"testing"
	"tobase/plugins/plugina"
)

func TestPlugins(t *testing.T) {
	var pls Plugins
	pls.InitPlugins()

	plsa := &plugina.Plugina{}
	pls.RegistePlugin(plsa)
	ret := pls.GetName(plsa)
	println(ret)

	retp := pls.GetPlugin("plugina")
	println(retp.Name())
}
