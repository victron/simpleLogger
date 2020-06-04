package simpleLogger

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var v, vv, vvv *bool

var (
	Debug   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	v = flag.Bool("v", false, "verbose level is WARNING")
	vv = flag.Bool("vv", false, "verbose level is INFO")
	vvv = flag.Bool("vvv", false, "verbose level is DEBUG")
}

func logVerbosity() {
	Debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Info = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Warning = log.New(ioutil.Discard, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

	switch {
	case *v:
		Warning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	case *vv:
		Warning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	case *vvv:
		Warning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	}
}

func FlagsParse() {
	flag.Parse()
	logVerbosity()
}
