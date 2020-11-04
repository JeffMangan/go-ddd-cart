package shared

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type FieldType string
type LogType string

func StandardErrorLog(er error) {
	log.Print(er)
}

const (
	FieldTypeError          FieldType = "Error"
	FieldTypeUserInput      FieldType = "UserInput"
	FieldTypeOutput         FieldType = "Output"
	FieldTypeStructName     FieldType = "StructName"
	FieldTypeFunctionName   FieldType = "FunctionName"
	FieldTypeUserError      FieldType = "UserError"
	FieldTypeSystemError    FieldType = "SystemError"
	FieldTypeFunctionCalled FieldType = "MethodCall"
	FieldTypeErrorInfo		FieldType = "ErrorInfo"
	//FieldTypeErrorContent FieldType = "ErrorContent"
	//FieldTypeConfigInfo   FieldType = "ConfigInfo"

	LogTypeInfo  LogType = "Info Log"
	LogTypeDebug LogType = "Debug Log"
	LogTypePanic LogType = "Panic Log"
	LogTypeError LogType = "Error Log"
)


type Logger struct {
	logger *logrus.Logger
	fields map[string]interface{}
}

func (l *Logger) AddFieldToLog(fieldType FieldType, val interface{}) {

	v := fmt.Sprintf("%v", val)

	if v == "" {
		panic("missing log field value")
	}

	l.fields[string(fieldType)] = v
}

func (l *Logger) Level() string {
	return l.logger.Level.String()
}

func (l *Logger) LogInfo(msg interface{}) {
	if !l.hasFields() {
		return
	}
	l.logger.WithFields(l.fields).Info(msg)
	l.clear()
	//l.log()
}
func (l *Logger) LogWarn(msg interface{}) {
	if !l.hasFields() {
		return
	}
	l.logger.WithFields(l.fields).Warn(msg)
	l.clear()
	//l.log()
}
func (l *Logger) LogDebug(msg interface{}) {
	if !l.hasFields() {
		return
	}
	l.logger.WithFields(l.fields).Debug(msg)
	l.clear()
	//l.log()
}
func (l *Logger) LogError(msg interface{}) {
	//fmt.Println("in logger")
	if !l.hasFields() {
		return
	}
	//fmt.Println("in logger2")
	//fmt.Println(len(l.fields))
	//fmt.Print(msg)
	//fmt.Print(l.logger)
	l.logger.WithFields(l.fields).Error(msg)
	//fmt.Println("in logger3")
	l.clear()
	//l.log()
}
func (l *Logger) LogPanic(msg interface{}) {
	if !l.hasFields() {
		return
	}
	l.logger.WithFields(l.fields).Panic(msg)
	l.clear()
	//l.log()
}

func (l *Logger) hasFields() bool {
	return len(l.fields) > 0
}

func (l *Logger) log() {
	//	fields, _ := json.Marshal(l.fields)
	l.clear()
	//		//fmt.Printf("this is a logger thing..")
	//	//fmt.Printf("%s\n\n", string(fields))
}

func (l *Logger) clear() {
	for k := range l.fields {
		delete(l.fields, k)
	}
}

// NewLogger returns a logger that logs to standard Out
func NewLogger(level string, formatter string) (*Logger, *CustomError) {
	var log = logrus.New()
	if formatter == "json" {
		log.Formatter = new(logrus.JSONFormatter)
	} else {
		log.Formatter = new(logrus.TextFormatter)
	}
	//fmt.Println("here...", level, formatter)
	var err error
	log.Level, err = logrus.ParseLevel(level)
	//fmt.Println("here2...")
	if err != nil {
		//fmt.Println(err.Error())
		return nil, NewCustomError(err.Error(), ErrorTypeSystem)
	}
	log.Out = os.Stderr
	l := &Logger{
		logger: log,
		fields: make(map[string]interface{}),
	}
	//fmt.Println("returning logger")
	return l, nil
}
