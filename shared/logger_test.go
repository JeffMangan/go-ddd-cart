package shared


import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// PanicLevel level, highest level of severity. Logs and then calls panic with the
// message passed to Debug, Info, ...
// PanicLevel Level = iota
// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
// logging level is set to Panic.
// FatalLevel
// ErrorLevel level. Logs. Used for errors that should definitely be noted.
// Commonly used for hooks to send errors to an error tracking service.
//	 ErrorLevel
// WarnLevel level. Non-critical entries that deserve eyes.
//	 WarnLevel
// InfoLevel level. General operational entries about what's going on inside the
// application.
//	 InfoLevel
// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
//	 DebugLevel

const (
	checkloglevel = "Debug"
)

func TestInterface(t *testing.T) {
	var li ILogger
	var err interface{}
	li, err = NewLogger(checkloglevel, "json")
	//fmt.Println(li)
	//fmt.Println(err)
	assert.NotNil(t, li, "Expected logger to not be empty")
	assert.Nil(t, err, "Expecting an empty error")
}


func TestLoggerConstruction(t *testing.T) {
	_, err := NewLogger(checkloglevel, "json")
	assert.Nil(t, err, "Expecting an empty error")
}
func TestLoggerFieldsAdded(t *testing.T) {
	l, _ := NewLogger(checkloglevel, "json")
	l.AddFieldToLog("Test", "testing")
	assert.Equal(t, len(l.fields), 1, "Expected logger to have 1 field")
}