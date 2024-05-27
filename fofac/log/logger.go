package log

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func Info(format string, arg ...any) {
	fmt.Fprint(os.Stdout, fmt.Sprintf(format, arg...), "\n")
}
func Debug(format string, arg ...any) {
	fmt.Fprint(os.Stdout, fmt.Sprintf(format, arg...), "\n")
}

func Warn(format string, arg ...any) {
	fmt.Fprint(os.Stdout, fmt.Sprintf(format, arg...), "\n")
}

func Error(format string, arg ...any) {
	var fileClient *os.File
	defer func() {
		fileClient.Close()
	}()
	dir, err2 := os.Getwd()
	if err2 != nil {
		fmt.Println("获取当前目录失败")
	}
	filePath := fmt.Sprintf("%s%s%s.log", dir, string(os.PathSeparator), time.Now().Format(time.DateOnly))
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		create, err := os.Create(filePath)
		if err != nil {
			debug.PrintStack()
		}
		fileClient = create
	}
	fileClient, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		debug.PrintStack()
	}
	fmt.Fprint(os.Stderr, fmt.Sprintf(format, arg...), "\n")
	fmt.Fprint(fileClient, fmt.Sprintf(format, arg...), "\n")
	debug.PrintStack()
}
