package main

import (
	"github.com/tmazitov/xpm_editor.git/app"
	"github.com/tmazitov/xpm_editor.git/tools/scale"
	"github.com/tmazitov/xpm_editor.git/tools/trim"
)

func main() {
	var core *app.App = app.NewApp(app.AppCreateOptions{
		Name:         "xpm_editor",
		MajorVersion: 0,
		MinorVersion: 1,
	})
	core.AddTool("scale", &scale.ScaleTool{})
	core.AddTool("trim", &trim.TrimTool{})
	core.Run()
}
