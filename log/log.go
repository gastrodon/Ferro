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

	info  *log.Logger = log.New(os.Stdout, "[info ] ", log.LstdFlags)
	trace *log.Logger = log.New(os.Stdout, "[trace] ", log.LstdFlags)
	err   *log.Logger = log.New(os.Stdout, "[error] ", log.LstdFlags)
	fatal *log.Logger = log.New(os.Stdout, "[fatal] ", log.LstdFlags)

	Fatal  func(v ...interface{})              = fatal.Fatal
	Fatalf func(base string, v ...interface{}) = fatal.Fatalf
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

func Traceln(v ...interface{}) {
	if level >= Ltrace {
		trace.Println(v...)
	}
}

func Tracef(base string, v ...interface{}) {
	if level >= Ltrace {
		trace.Printf(base, v...)
	}
}
