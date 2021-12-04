package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string){
// リードライト、なければ作成、あったら追記
// |は論理和(or)を示しており、10進数を2進数に変えて、論理話の数値になる
	logfile,err:=os.OpenFile(logFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln(err)
	}
	// 標準出力とログファイルの両方に書き込む,io.Writerを宣言
	multiLogFile:=io.MultiWriter(os.Stdout,logfile)

	//ログのフォーマットを指定。今回は1,2,4なので、
	// 0001,0010,0100,=>0111=>7になる
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	
	//ログの出力先を指定=>標準出力とログファイル
	log.SetOutput(multiLogFile)
}