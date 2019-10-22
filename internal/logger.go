package internal

import (
	"github.com/op/go-logging"
	"os"
	"strings"
)

var log = logging.MustGetLogger("")

func LogInit() {
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	format := logging.MustStringFormatter(
		`%{color}%{time} [%{level}] ▶ %{color:reset}%{message}`,
	)
	logging.SetFormatter(format)

	levelBackend := logging.AddModuleLevel(stdout)
	switch strings.ToUpper(*LOG_LEVEL) {
	case "CRITICAL":
		levelBackend.SetLevel(logging.CRITICAL, "")
	case "ERROR":
		levelBackend.SetLevel(logging.ERROR, "")
	case "WARN":
		levelBackend.SetLevel(logging.WARNING, "")
	case "NOTICE":
		levelBackend.SetLevel(logging.NOTICE, "")
	case "DEBUG":
		levelBackend.SetLevel(logging.DEBUG, "")
	default:
		levelBackend.SetLevel(logging.INFO, "")
	}

	logging.SetBackend(levelBackend)
	log.Info(logging.GetLevel(log.Module))
}



