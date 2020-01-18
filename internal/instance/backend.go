package instance

import (
	"fmt"
	"os"
)

type Backend string

const (
	Node   Backend = "node"
	Npm    Backend = "npm"
	Python Backend = "python"
	Web    Backend = "web"
	Flask  Backend = "flask"
)

func IsBackendValid(bkend Backend) bool {
	switch bkend {
	case Npm, Node, Python, Web, Flask:
		return true
	}
	_, _ = fmt.Fprintf(os.Stderr, "invalid enum '%s'\n", bkend)
	return false
}
