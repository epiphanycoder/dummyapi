package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const baseLogLevel = log.DebugLevel
const baseLogFilePath = "./dummyApi.log"

type logFormatter struct {
	log.TextFormatter
}

//func init() {
//	log.SetFormatter(&logFormatter{
//		log.TextFormatter{
//			FullTimestamp:          true,
//			ForceColors:            false,
//			TimestampFormat:        "2006-01-02 15:04:05",
//			DisableLevelTruncation: true,
//		},
//	})
//	log.SetReportCaller(true)
//}

func Configure(filename, level string) *os.File {
	logFile, err := os.OpenFile(getFileName(filename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed to create file %s: %v", filename, err)
	}
	logger := io.MultiWriter(os.Stderr, logFile)
	log.SetLevel(getLogLevel(level))
	log.SetOutput(logger)
	return logFile
}

func getFileName(filename string) string {
	if strings.EqualFold(filename, "") {
		return baseLogFilePath
	}
	return filename
}

func (lf *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] [%s] [%s:%s:%d] => %s\n", strings.ToUpper(entry.Level.String()),
		entry.Time.Format(lf.TimestampFormat),
		entry.Caller.File,
		entry.Caller.Function,
		entry.Caller.Line, entry.Message),
	), nil
}

func getLogLevel(level string) log.Level {
	lvl, err := log.ParseLevel(strings.ToLower(level))
	if err != nil {
		return baseLogLevel
	}
	return lvl
}
