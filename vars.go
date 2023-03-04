package hypolasjsondecomposer

import (
	"github.com/hypolas/hypolaslogger"
	"os"
	"time"
)

type JSONKey struct {
	Name       string
	KeyIsArray bool
	ArrayIndex int
	IsLast     bool
}

type JSONData struct {
	FlattenPath string
	JSONFile    []byte
}

type HlckHYPOLAS struct {
	IsUp      bool
	WaitUp    bool
	WaitRetry time.Duration
}

var (
	separator = "__"
	logf      = hypolaslogger.NewLogger(os.Getenv("HYPOLAS_LOGS_FOLDER"))
)
