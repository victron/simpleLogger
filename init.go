package simpleLogger

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var path *string
var v, vv, vvv *bool

var (
	DebugLogger *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	path = flag.String("uri", "", "\"uri\" to use")
	v = flag.Bool("v", false, "verbose level is WARNING")
	vv = flag.Bool("vv", false, "verbose level is INFO")
	vvv = flag.Bool("vvv", false, "verbose level is DEBUG")
	flag.Parse()

	logVerbosity(v, vv, vvv)

}

func logVerbosity(v *bool, vv *bool, vvv *bool) {
	DebugLogger = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	InfoLogger = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	WarningLogger = log.New(ioutil.Discard, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

	switch {
	case *v:
		WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	case *vv:
		WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	case *vvv:
		WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
		DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	}
}



