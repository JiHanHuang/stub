package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/JiHanHuang/stub/pkg/file"
	"github.com/JiHanHuang/stub/pkg/setting"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	std        *log.Logger
	logPrefix  = ""
	LevelFlags = []string{"DEBU", "INFO", "WARN", "ERRO", "FATA"}
	level      = INFO
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("[%s] Logging.Setup err: %v", LevelFlags[FATAL], err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	std = log.New(os.Stdout, DefaultPrefix, log.LstdFlags)

	switch setting.AppSetting.LogLevel {
	case "debug":
		level = DEBUG
	case "info":
		level = INFO
	case "warn":
		level = WARNING
	case "error":
		level = ERROR
	case "fatal":
		level = FATAL
	default:
		level = INFO
		log.Printf("[%s] Not support log level %s. reset to info level",
			LevelFlags[ERROR], setting.AppSetting.LogLevel)
	}
	log.Printf("[%s] Log level:%s	Log std: %v", LevelFlags[INFO], LevelFlags[level],
		setting.AppSetting.LogStdOut)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	if level <= DEBUG {
		if setting.AppSetting.LogStdOut {
			std.Println(v...)
		}
		logger.Println(v...)
	}
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	if level <= INFO {
		if setting.AppSetting.LogStdOut {
			std.Println(v...)
		}
		logger.Println(v...)
	}
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	if level <= WARNING {
		if setting.AppSetting.LogStdOut {
			std.Println(v...)
		}
		logger.Println(v...)
	}
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	if level <= ERROR {
		if setting.AppSetting.LogStdOut {
			std.Println(v...)
		}
		logger.Println(v...)
	}
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	if level <= FATAL {
		if setting.AppSetting.LogStdOut {
			std.Fatalln(v...)
		}
		logger.Fatalln(v...)
	}
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
	}

	if setting.AppSetting.LogStdOut {
		std.SetPrefix(logPrefix)
	}

	logger.SetPrefix(logPrefix)
}
