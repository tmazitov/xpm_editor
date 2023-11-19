package tools

type ToolStorage struct {
	Tools []*Tool
}

func NewToolStorage() *ToolStorage {
	return &ToolStorage{
		Tools: []*Tool{},
	}
}

func (ts *ToolStorage) AddTool(title string, handler ToolHandler) {
	ts.Tools = append(ts.Tools, NewTool(len(ts.Tools), title, handler))
}
