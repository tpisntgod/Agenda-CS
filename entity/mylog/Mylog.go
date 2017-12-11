package mylog

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var logDivPath = "src/github.com/tpisntgod/Agenda/Log"
var logFilePath = "/" + time.Now().Format("2006-01-02") + ".txt"

func init() {
	logDivPath = filepath.Join(*GetGOPATH(), logDivPath)
}

//GetGOPATH 获得用户环境的gopath
func GetGOPATH() *string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	goPath := strings.Split(os.Getenv("GOPATH"), sp)
	for _, v := range goPath {
		if _, err := os.Stat(filepath.Join(v, "/src/github.com/tpisntgod/Agenda/entity/meeting/meeting.go")); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}

func getFileHandle() *os.File {
	if _, err := os.Open(logDivPath + logFilePath); err != nil {
		os.Create(logDivPath + logFilePath)
	}

	// 以追加模式打开文件,并向文件写入
	fi, _ := os.OpenFile(logDivPath+logFilePath, os.O_RDWR|os.O_APPEND, 0)
	return fi
}

// AddLog : 添加记录
func AddLog(user string, command string, oldStr string, newStr string) {
	file := getFileHandle()
	l := log.New(file, "[INFO]", log.Ltime)
	outStr := ""
	if user != "" {
		outStr += "User:" + user + "  "
	}
	if command != "" {
		outStr += "Command:" + command + "\n"
	}
	if oldStr != "" {
		outStr += "From:" + oldStr + "\n"
	}
	if newStr != "" {
		outStr += "To:" + newStr + "\n"
	}
	l.Print(outStr)
	file.Close()
}
