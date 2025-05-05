package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルの読み込み
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 読み書き、ファイルの作成、追記ができるように。

	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // ログの書き込み先を標準出力とログファイルに指定している
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // setFlagsでログフォーマットを指定している
	log.SetOutput(multiLogFile)                          // ログの出力先を指定

}
