package shared

type ILogger interface {
	//AddFieldToLog will add a log field to the collection to be logged
	AddFieldToLog(fieldType FieldType, val interface{})

	//Level returns the logger level that was set
	Level() string

	//LogInfo logs everything
	LogInfo(msg interface{})

	//LogWarn logs everything from warn and above and then clears out all fields
	LogWarn(msg interface{})

	//LogDebug logs everything from debug and above and then clears out all fields
	LogDebug(msg interface{})

	//LogError logs everything from error and above and then clears out all fields
	LogError(msg interface{})

	//LogPanic logs everything from panic and above and then clears out all fields
	LogPanic(msg interface{})
}
