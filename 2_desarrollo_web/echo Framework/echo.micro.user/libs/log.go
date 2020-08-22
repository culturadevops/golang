package libs

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {
	os.Mkdir("logs", 0755)
	logFile, err := os.OpenFile("./logs/run_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("open log file failed", err)
	}

	//日志
	LogInfo = log.New(io.MultiWriter(logFile), "【Info】:", log.Ldate|log.Ltime|log.Lshortfile)   //LogInfo.Println(1, 2, 3)
	LogError = log.New(io.MultiWriter(logFile), "【Error】:", log.Ldate|log.Ltime|log.Lshortfile) //LogError.Println(4, 5, 6)
}
