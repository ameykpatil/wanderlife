package utils

import (
	"encoding/json"
	"flag"
	"log"
	"runtime"
	"strings"
	"time"
)

func init() {
	if flag.Lookup("test.v") != nil {
		logLine = func(userID, authUID, message string, err error, logLevel string) {
			return
		}
	}
}

//LogObj is struct with all the log info
type LogObj struct {
	Timestamp string `json:"timestamp"`
	LogLevel  string `json:"logLevel"`
	Class     string `json:"class"`
	AuthUID   string `json:"authUid"`
	UserID    string `json:"userId"`
	Message   string `json:"msg"`
	Error     string `json:"error,omitempty"`
}

//LogInfo just logs with logLevel INFO
func LogInfo(userID, authUID, message string) {
	logLine(userID, authUID, message, nil, "INFO")
}

//LogErr just logs with logLevel ERROR with error
func LogErr(userID, authUID, message string, err error) {
	logLine(userID, authUID, message, err, "ERROR")
}

//LogWarn just logs with logLevel WARN with error
func LogWarn(userID, authUID, message string, err error) {
	logLine(userID, authUID, message, err, "WARN")
}

//LogFatal just logs with logLevel FATAL with error
func LogFatal(userID, authUID, message string, err error) {
	logLine(userID, authUID, message, err, "FATAL")
}

var logLine = func(userID, authUID, message string, err error, logLevel string) {
	_, className, _, _ := runtime.Caller(2)
	parts := strings.Split(className, "/")
	part := parts[len(parts)-1]
	arr := strings.Split(part, ".")
	class := arr[0]
	logObj := LogObj{
		Timestamp: time.Now().Format("2006/01/02 15:04:05"),
		UserID:    userID,
		AuthUID:   authUID,
		Message:   message,
		Class:     class,
		LogLevel:  logLevel,
	}
	if err != nil {
		logObj.Error = err.Error()
	}
	json, _ := json.Marshal(logObj)
	log.Println(string(json))
}
