package logger

import (
	"currency-api/internal/tracing"
	formatters "github.com/fabienm/go-logrus-formatters"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
)

/*
var Log *zap.SugaredLogger

func NewLogger(graylog string) error {
	zap.NewDevelopment()
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	core, err := gelf.NewCore(
		gelf.Addr(graylog),
		gelf.Host(host),
		gelf.Level(zap.InfoLevel),
	)
	if err != nil {
		return err
	}
	l := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return core.Enabled(l)
		})),
	)
	sugaredLogger := l.Sugar()
	Log = sugaredLogger
	return nil
}
*/

func InitializeLogger() {
	fmter := formatters.NewGelf(tracing.Name)
	hooks := []logrus.Hook{graylog.NewGraylogHook(os.Getenv("GRAYLOG_HOST"), map[string]interface{}{})}
	log.SetFormatter(fmter)
	for _, h := range hooks {
		log.AddHook(h)
	}
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	lvl, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		lvl = log.DebugLevel
	}
	log.SetLevel(lvl)

}
