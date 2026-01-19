package tools


type PingTools struct {
	ToolsItem
}

func(t PingTools) Ping() (string, error) {
	return "In progress", nil
}