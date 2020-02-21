package log

import (
	"log"
	"os"
)

const (
	Lerror int = 0
	Linfo  int = 1
	Ltrace int = 2
)

var (
	level int = 1

	info  *log.Logger = log.New(os.Stdout, "[info ] ", log.Llongfile)
	trace *log.Logger = log.New(os.Stdout, "[trace] ", log.Llongfile)
	err   *log.Logger = log.New(os.Stdout, "[error] ", log.Llongfile)
	fatal *log.Logger = log.New(os.Stdout, "[fatal] ", log.Llongfile)

	Fatal func(v ...interface{}) = fatal.Fatal
)

func At(new_level int) {
	level = new_level
}

func Printf(base string, v ...interface{}) {
	if level >= Linfo {
		info.Printf(base, v...)
	}
}

func Println(v ...interface{}) {
	if level >= Linfo {
		info.Println(v...)
	}
}

func Errorf(base string, v ...interface{}) {
	if level >= Lerror {
		err.Printf(base, v...)
	}
}
