package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

type Logger struct {
	*logrus.Entry
}

var instance *Logger
var once sync.Once

func GetLogger(level string) *Logger {
	once.Do(func() {
		logrusLevel, err := logrus.ParseLevel(level)
		if err != nil {
			log.Fatalln(err)
		}

		l := logrus.New()

		l.SetReportCaller(true)
		l.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
			},
			DisableColors: false,
			FullTimestamp: true,
		}
		l.SetOutput(os.Stdout)
		l.SetLevel(logrusLevel)

		instance = &Logger{logrus.NewEntry(l)}
	})

	return instance
}
