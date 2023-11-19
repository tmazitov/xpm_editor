package tools

type ToolHandler interface {
	Execute(filePaths []string) error
}

type Tool struct {
	Id      int
	Title   string
	Handler ToolHandler
}

func NewTool(id int, title string, handler ToolHandler) *Tool {
	return &Tool{
		Id:      id,
		Title:   title,
		Handler: handler,
	}
}
