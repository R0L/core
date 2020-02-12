package logger

type Config struct {
	Type   string // logrus, default
	Out    string // path
	Level  string // info，debug，error
	Format string // json，text
}

var (
	config = &Config{
		Type:   "logrus",
		Out:    "/Users/thooh/Workspace/go/src/github.com/R0L/core/logs",
		Level:  "debug",
		Format: "json",
	}
)
