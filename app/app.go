package app

import (
	"log"

	"github.com/tmazitov/xpm_editor.git/tools"
)

type AppCreateOptions struct {
	Name         string
	MajorVersion int8
	MinorVersion int8
}

type App struct {
	Name         string
	MajorVersion int8
	MinorVersion int8
	toolStorage  *tools.ToolStorage
}

func NewApp(options AppCreateOptions) *App {
	return &App{
		Name:         options.Name,
		MajorVersion: options.MajorVersion,
		MinorVersion: options.MinorVersion,
		toolStorage:  tools.NewToolStorage(),
	}
}

func (a *App) AddTool(title string, handler tools.ToolHandler) {
	a.toolStorage.AddTool(title, handler)
}

func (a *App) Run() {
	var (
		toolId      int
		filePaths   []string
		toolHandler tools.ToolHandler
	)

	toolId = a.selectTool()
	toolHandler = nil
	for _, tool := range a.toolStorage.Tools {
		if tool.Id == toolId {
			toolHandler = tool.Handler
		}
	}
	if toolHandler == nil {
		return
	}
	filePaths = a.enterFilePaths()
	if len(filePaths) == 0 {
		return
	}
	if err := toolHandler.Execute(filePaths); err != nil {
		log.Println(err)
	}
}
