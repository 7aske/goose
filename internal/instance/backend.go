package instance

type Backend string

const (
	Node   Backend = "node"
	Npm    Backend = "npm"
	Python Backend = "python"
	Web    Backend = "web"
	Flask  Backend = "flask"
)
