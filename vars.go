package hypolasjsondecomposer

import (
	"github.com/hypolas/hypolaslogger"
	"os"
)

// JSONKey store JSONKey configuration
type JSONKey struct {
	Name       string
	KeyIsArray bool
	ArrayIndex int
	IsLast     bool
}

var (
	separator = "__"
	logf      = hypolaslogger.NewLogger(os.Getenv("HYPOLAS_LOGS_FOLDER"))
)
