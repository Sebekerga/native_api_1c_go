package main

/*
#include <logging.h>
*/
import "C"
import (
	"fmt"
	"io"
	"log"
	"os"
	"unsafe"

	"github.com/imroc/req/v3"
)

func openLogFile() io.Writer {
	f, err := os.OpenFile("C:\\Users\\sebek\\Documents\\native_api_1c_go\\component.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

type HttpLogger struct {
	url string
}

func (l HttpLogger) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	req.MustGet(l.url + "/log/" + message)
}

func (l HttpLogger) Println(v ...interface{}) {
	message := fmt.Sprint(v...)
	req.MustGet(l.url + "/log/" + message)
}

const DUMMYHTTP_URL = "http://localhost:8080"

var Logger = log.New(openLogFile(), "GO: ", log.LstdFlags|log.Lshortfile)

// var Logger = &HttpLogger{url: DUMMYHTTP_URL}

func CreateLogger() HttpLogger {
	return HttpLogger{url: DUMMYHTTP_URL}
}

//export logToConsole
func logToConsole(message *C.char) {
	Logger.Println(C.GoString(message))
}

//export logPointerToConsole
func logPointerToConsole(pointer *C.void) {
	Logger.Printf("Pointer: %v\n", unsafe.Pointer(pointer))
}
