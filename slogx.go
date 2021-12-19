package slogx

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level uint

const (
	NONE Level = iota
	FATAL
	ERROR
	WARNING
	INFO
	DEBUG
)

var levelToString = map[Level]string{
	NONE:    "NONE",
	FATAL:   "FATAL",
	ERROR:   "ERROR",
	WARNING: "WARNING",
	INFO:    "INFO",
	DEBUG:   "DEBUG",
}

func (l Level) String() string {
	return levelToString[l]
}

var stringToLevel = map[string]Level{
	"NONE":    NONE,
	"FATAL":   FATAL,
	"ERROR":   ERROR,
	"WARNING": WARNING,
	"INFO":    INFO,
	"DEBUG":   DEBUG,
}

var loggers = make(map[string]*Logger)

type Logger struct {
	Name       string
	Level      Level
	Format     string
	TimeFormat string
	Output     io.Writer
	Mutex      sync.Mutex
}

// NewLogger returns a new Logger.
func NewLogger(name string) *Logger {
	logger := &Logger{
		Name:       name,
		Level:      INFO,
		Format:     "%[1]s %[2]s %[3]s:%[4]d %[5]s: %[6]s",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     os.Stdout,
	}
	loggers[logger.Name] = logger
	return logger
}

// GetLogger returns a Logger by its name.
func GetLogger(name string) *Logger {
	return loggers[name]
}

// ParseLevel returns a logging Level based on its string name.
func ParseLevel(level string) Level {
	return stringToLevel[strings.ToUpper(level)]
}

// SetLevel sets the logging Level for the Logger.
func (l *Logger) SetLevel(level Level) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	l.Level = level
}

// GetLevel returns the current logging Level for the Logger.
func (l *Logger) GetLevel() Level {
	return l.Level
}

// SetFormat sets the Format for the Logger.
func (l *Logger) SetFormat(format string) error {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	parsed, err := parseFormat(format)
	if err != nil {
		return err
	}
	l.Format = parsed
	return nil
}

// SetTimeFormat sets the TimeFormat for the Logger.
func (l *Logger) SetTimeFormat(layout string) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	l.TimeFormat = layout
}

// SetOutput sets the Output for the Logger.
func (l *Logger) SetOutput(writer io.Writer) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	l.Output = writer
}

var formatPlaceholders = map[string]string{
	"${time}":    "%[1]s",
	"${level}":   "%[2]s",
	"${file}":    "%[3]s",
	"${line}":    "%[4]d",
	"${name}":    "%[5]s",
	"${message}": "%[6]s",
}

func parseFormat(format string) (string, error) {
	format = strings.Replace(format, "%", "%%", -1)
	re := regexp.MustCompile("\\${([a-zA-Z]+)}")
	m := re.FindAllStringSubmatch(format, -1)
	if m != nil {
		for _, v := range m {
			placeholder := formatPlaceholders[v[0]]
			if placeholder == "" {
				return "", fmt.Errorf("slogx: invalid verb '%s'", v[0])
			}
			format = strings.Replace(format, v[0], placeholder, -1)
		}
	} else {
		return "", fmt.Errorf("slogx: invalid format '%s'", format)
	}
	return format, nil
}

func (l *Logger) write(log string) {
	_, err := fmt.Fprintln(l.Output, log)
	if err != nil {
		fmt.Println(fmt.Errorf("slogx: %v", err))
	}
}

// Log logs a message at the specified Level.
func (l *Logger) Log(level Level, args ...interface{}) {
	if l.Level < level || level == NONE {
		return
	}
	msg := fmt.Sprint(args...)
	ts := time.Now().Format(l.TimeFormat)
	_, fl, ln, _ := runtime.Caller(2)
	log := fmt.Sprintf(l.Format, ts, level.String(), filepath.Base(fl), ln, l.Name, msg)
	l.write(log)
}

// Logf logs a message at the specified Level with formatting.
func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	l.Log(level, fmt.Sprintf(format, args...))
}

// Fatal logs a message at FATAL Level and exits.
func (l *Logger) Fatal(args ...interface{}) {
	l.Log(FATAL, fmt.Sprint(args...))
	os.Exit(1)
}

// Fatalf logs a message at FATAL Level with formatting and exits.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Log(FATAL, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Error logs a message at ERROR Level.
func (l *Logger) Error(args ...interface{}) {
	l.Log(ERROR, fmt.Sprint(args...))
}

// Errorf logs a message at ERROR Level with formatting.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Log(ERROR, fmt.Sprintf(format, args...))
}

// Warning logs a message at WARNING Level.
func (l *Logger) Warning(args ...interface{}) {
	l.Log(WARNING, fmt.Sprint(args...))
}

// Warningf logs a message at WARNING Level with formatting.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Log(WARNING, fmt.Sprintf(format, args...))
}

// Info logs a message at INFO Level.
func (l *Logger) Info(args ...interface{}) {
	l.Log(INFO, fmt.Sprint(args...))
}

// Infof logs a message at INFO Level with formatting.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Log(INFO, fmt.Sprintf(format, args...))
}

// Debug logs a message at DEBUG Level.
func (l *Logger) Debug(args ...interface{}) {
	l.Log(DEBUG, fmt.Sprint(args...))
}

// Debugf logs a message at DEBUG Level with formatting.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Log(DEBUG, fmt.Sprintf(format, args...))
}
